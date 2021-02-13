package thenovadiary

import (
	"fmt"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/photoprism-client-go"
	"github.com/kris-nova/photoprism-client-go/api/v1"
)

// List all photos
// Check each photo for last tweet date
// If last tweet date < today | tweet | flip
// If !flip | tweet[0] | cache ++
// Save cache

// FindNextPhoto will list all photos and
// linear search for a photo that was last
// updated with the greatest delta in days
// ago based on the timestamp found in
//
//      photo.PhotoDescription (timestamp)
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

	// ------ [ System to Find Photo ] -------
	today := TimeToday()
	var photoGreatestDelta *api.Photo
	gDelta := 0 // This will let us know if we have found a photo
	for _, photo := range photos {
		// ----------------------------------
		// We need to handle 3 cases
		//   - No  *Details{} (No  timestamp)
		//   - Yes *Details{} (No  timestamp)
		//   - Yes *Details{} (Yes timestamp)
		if photo.Details == nil {
			photo.Details = &api.Details{}
			// This photo has never been processed
			//   - No  *Details{} (No  timestamp)
			return &photo, nil
		} else {
			lastTweetStr := photo.PhotoDescription
			if lastTweetStr == "" {
				// This photo has never been processed
				//   - Yes  *Details{} (No  timestamp)
				return &photo, nil
			}
			//   - Yes *Details{} (Yes timestamp)
			timeFromDB, err := TimeStringToTime(lastTweetStr)
			if err != nil {
				return nil, fmt.Errorf("unable to parse time from DB have {%s}", lastTweetStr)
			}
			pDelta := TimeDeltaDays(today, *timeFromDB)
			if pDelta > gDelta {
				// This system will grab the greatest delta
				// or oldest photo first
				gDelta = pDelta
				photoGreatestDelta = &photo
			}
		}
	}

	if gDelta > 0 {
		// We have found the oldest photo
		return photoGreatestDelta, nil
	}
	// Interesting case
	// Here we have photos with no timestamp < today
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
