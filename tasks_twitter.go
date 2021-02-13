package thenovadiary

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"

	"github.com/kris-nova/logger"
)

const (
	WellKnownAlbumID = "aqofyebwd9187e3e"
)

func (d *Diary) SendDailyTweet() error {
	// ------- [ Get Photo From Photo Search Algorithm ] ------
	pClient, err := d.NewPhotoprismClient()
	if err != nil {
		return fmt.Errorf("unable to build new photoprism client: %v", err)
	}
	pphoto, err := FindNextPhoto(pClient)
	if err != nil {
		return err
	}
	photo := *pphoto
	logger.Debug("Processing photo: %s", photo.UUID)

	// ------- [ Get Photo Content ] ------
	pBytes, err := pClient.V1().GetPhotoDownload(photo.PhotoUID)
	if err != nil {
		return fmt.Errorf("unable to download photo: %v", err)
	}
	logger.Debug("Downloaded photo %d kb", len(pBytes)/1024)

	// ------- [ Auth Twitter Client ] ------
	tURL := "https://twitter.com/thenovadiary"
	if !BypassTwitter {
		tClient := anaconda.NewTwitterApiWithCredentials("", "", "", "")
		tURL, err = SendPhotoTweet(tClient, photo, pBytes)
		if err != nil {
			return fmt.Errorf("unable to send tweet: %v", err)
		}
	}

	// ------- [ Photo Found, Tweet Sent (lol), Update Cache ] ------
	logger.Debug("Syncing photo: %s", photo.PhotoUID)
	aPhoto, err := pClient.V1().GetPhoto(photo.PhotoUID)
	if err != nil {
		return fmt.Errorf("unable to sync photo: %v", err)
	}
	aPhoto.PhotoDescription = TimeTimeToString(TimeToday())
	logger.Debug("Updated timestamp: %s", aPhoto.PhotoDescription)
	_, err = pClient.V1().UpdatePhoto(aPhoto)
	if err != nil {
		return fmt.Errorf("unable to update photo: %s", err)
	}
	logger.Always("Successful tweet: %s", tURL)
	return nil
}
