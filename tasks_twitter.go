package thenovadiary

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kris-nova/photoprism-client-go/api/v1"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/photoprism-client-go"
)

const (
	DailyTweetCacheKey = "dailytweetkey"
	PhotoCountKey      = "dailytweetcount"
	WellKnownAlbumID   = "aqofyebwd9187e3e"
)

func Today() time.Time {
	now := time.Now()
	date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	//logger.Debug("Date: %s", date.String())
	return date
}

func (d *Diary) SendDailyTweet() error {
	// System that will determine eligibility
	// to run this method
	// ---
	v := d.cache.Get(DailyTweetCacheKey)
	today := Today()
	sendTweet := false
	if !v.Found {
		logger.Info("Daily Tweet cache not found...")
		sendTweet = true
	} else {
		// TODO Nova check the error...
		// TODO Nova write a fucking unit test
		// TODO make the cron schedulable dynamically
		lastTweetRecord := v.Value
		if lastTweetRecord != nil {
			// lastTweetRecord string
			lastTweet, err := time.Parse("2006-01-02T15:04:05Z07:00", fmt.Sprintf("%v", lastTweetRecord))
			if err != nil {
				return fmt.Errorf("unable to parse timestamp: %v", err)
			}
			days := int(today.Sub(lastTweet).Hours() / 24)
			//logger.Debug("LastTweet:  %s", lastTweet)
			//logger.Debug("Today:      %s", today)
			logger.Debug("[DailyTweet] Delta Days: %d", days)
			if days > 0 {
				sendTweet = true
			}
		} else {
			logger.Warning("nil time value")
		}
		// Edge case to move on that will hit almost always
		//logger.Debug("...")
	}

	// ---

	// Check to see if we need to send a tweet
	if !sendTweet {
		return nil
	}

	// If we get here we should tweet!

	// Init Photoprism Client
	logger.Debug("Connecting to photoprism (%s)", d.config.PhotoprismConn)
	logger.Debug("Username: %s", d.config.PhotoprismUser)
	mask := ""
	for i := 0; i < len(d.config.PhotoprismPass); i++ {
		mask = fmt.Sprintf("%s%s", mask, "*")
	}
	logger.Debug("Password: %s", mask)
	client := photoprism.New(d.config.PhotoprismConn)
	err := client.Auth(photoprism.NewClientAuthLogin(d.config.PhotoprismUser, d.config.PhotoprismPass))
	if err != nil {
		return fmt.Errorf("unable to auth with photoprism client: %v", err)
	}
	logger.Always("Successfully authenticated with Photoprism!")
	// Get Photo

	// TODO Pageinate if needed later
	photosInAlbum, err := client.V1().GetPhotos(&api.PhotoOptions{
		AlbumUID: WellKnownAlbumID,
		Count:    500,
	})

	logger.Debug("Length of photos: %d", len(photosInAlbum))

	if len(photosInAlbum) < 2 {
		return fmt.Errorf("unable to find photos in album")
	}

	logger.Info("Found [%d] photos to process..", len(photosInAlbum))
	count := 1
	countRecord := d.cache.Get(PhotoCountKey)
	if countRecord.Found {
		count = countRecord.Value.(int)
	}
	// --- [ Searching Algorithm ] ---
	//
	found := false
	for _, photo := range photosInAlbum {
		// Check if we have posted this
		pCount := 0
		for _, label := range photo.Labels {
			if label.Label.LabelName == "pcount" {
				p, err := strconv.Atoi(label.Label.LabelDescription)
				if err != nil {
					//No error here
					continue
				}
				pCount = p
			}
		}
		// Now pcount is updated
		// count 1, pcount 0
		// count 1, pcount 1
		// count 2, pcount 1
		if pCount < count {
			found = true
			err := tweet(photo)
			if err != nil {
				return fmt.Errorf("unable to tweet photo: %v", err)
			}
			logger.Always("Successful tweet of photo: %s", photo.PhotoTitle)
			labelUpdate := false
			for i, label := range photo.Labels {
				if label.Label.LabelName == "pcount" {
					//pCount++
					label.Label.LabelDescription = fmt.Sprintf("%d", pCount+1)
					photo.Labels[i] = label
					uPhoto, err := client.V1().UpdatePhoto(photo)
					if err != nil {
						return fmt.Errorf("unable to update photo: %v", err)
					}
					logger.Info("Successfully updated photo: %s", uPhoto.PhotoTitle)
					labelUpdate = true
				}
			}
			if !labelUpdate {
				// Create Label
				newLabel := api.PhotoLabel{
					Label: &api.Label{
						LabelName:        "pcount",
						LabelDescription: fmt.Sprintf("1"),
					},
				}
				photo.Labels = append(photo.Labels, newLabel)
				uPhoto, err := client.V1().UpdatePhoto(photo)
				if err != nil {
					return fmt.Errorf("unable to update photo: %v", err)
				}
				logger.Info("Successfully created photo pcount label: %s", uPhoto.PhotoTitle)
				labelUpdate = true
			}
		}
	}

	if !found {
		// Note we are safe to hardcode this 0 because we check above
		photo := photosInAlbum[0]
		err := tweet(photo)
		if err != nil {
			return fmt.Errorf("unable to tweet photo: %s", photo.PhotoTitle)
		}
		countRecord.Value = count + 1
		d.cache.Set(PhotoCountKey, countRecord)
	}

	// Init Twitter Client
	// Tweet Photo + Description
	logger.Always("Sending Tweet")

	save := &Record{
		Key:   DailyTweetCacheKey,
		Value: today,
	}
	logger.Always("Saving cache with today's date")
	d.cache.Set(DailyTweetCacheKey, save)
	d.cache.Persist()
	return nil
}

func tweet(photo api.Photo) error {

	logger.Always("Win! Will Tweet: %s", photo.PhotoTitle)
	return nil
}