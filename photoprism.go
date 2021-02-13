package thenovadiary

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/photoprism-client-go"
	"github.com/kris-nova/photoprism-client-go/api/v1"
)

type Notes struct {
	LastTweet   *time.Time
	NoteStrings []string
	KeyValue    map[string]string
}

// AddNotes will add a *Notes{} pointer as JSON to:
//     photo.PhotoDescription
func AddNotes(notes *Notes, photo api.Photo) (api.Photo, error) {
	jBytes, err := json.Marshal(&notes)
	if err != nil {
		return photo, fmt.Errorf("unable to add notes: %v", err)
	}
	photo.PhotoDescription = string(jBytes)
	return photo, nil
}

// GetNotes will return a *Notes{} pointer from:
//     photo.PhotoDescription
func GetNotes(photo api.Photo) *Notes {
	notes := &Notes{}
	if photo.PhotoDescription == "" {
		return notes
	}
	noteStr := photo.PhotoDescription
	err := json.Unmarshal([]byte(noteStr), &notes)
	if err != nil {
		logger.Warning("INVALID JSON in Notes: %v", err)
		return notes
	}
	return notes
}

// FindNextPhoto will list all photos and
// linear search for a photo that was last
// updated with the greatest delta in days
// ago based on the timestamp found in
//
//
// The function will return unprocessed
// photos first by design.
func FindNextPhoto(client *photoprism.Client) (*api.Photo, error) {
	// TODO Nóva add pagination
	photos, err := client.V1().GetPhotos(&api.PhotoOptions{
		AlbumUID: WellKnownAlbumID,
		Count:    500, // TODO This is an  enormous number holy shit
	})
	if err != nil {
		return nil, fmt.Errorf("unable to list photos: %v", err)
	}

	today := TimeToday()
	var photoGreatestDelta *api.Photo
	gDelta := 0 // This will let us know if we have found a photo
	for _, photo := range photos {
		notes := GetNotes(photo)
		if notes.LastTweet == nil {
			return &photo, nil
		} else {
			timeFromDB := notes.LastTweet
			pDelta := TimeDeltaDays(today, *timeFromDB)
			if pDelta > gDelta {
				gDelta = pDelta
				photoGreatestDelta = &photo
			}
		}
	}
	if gDelta > 0 {
		return photoGreatestDelta, nil
	}
	// Error
	return nil, fmt.Errorf("unable to find photo with old timestamp")
}

// NewPhotoprismClient will always return a new client with a fresh token
func (d *Diary) NewPhotoprismClient() (*photoprism.Client, error) {
	mask := ""
	for i := 0; i < len(d.config.PhotoprismPass); i++ {
		mask = fmt.Sprintf("%s%s", mask, "*")
	}
	logger.Debug("Photoprism client (%s) %s: %s", d.config.PhotoprismConn, d.config.PhotoprismUser, mask)
	client := photoprism.New(d.config.PhotoprismConn)
	err := client.Auth(photoprism.NewClientAuthLogin(d.config.PhotoprismUser, d.config.PhotoprismPass))
	if err != nil {
		return nil, fmt.Errorf("unable to auth with photoprism client: %v", err)
	}
	return client, nil
}
