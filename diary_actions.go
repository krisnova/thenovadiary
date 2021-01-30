package thenovadiary

import (
	"fmt"

	"github.com/kris-nova/logger"
)

var ActionMap map[string]Action = map[string]Action{
	"daily": DailyPhotoTweet,
}

func ActionsString() string {
	a := ""
	for key, _ := range ActionMap {
		if a == "" {
			a = fmt.Sprintf("%s", key)
		} else {
			a = fmt.Sprintf("%s, %s", a, key)
		}
	}
	return a
}

type Action func(diary *Diary) error

// DailyPhotoTweet is used to pull a photo from
// the CDN photo storage and send a tweet with the
// content and title of the Photo
func DailyPhotoTweet(diary *Diary) error {
	logger.Debug("Running DailyPhotoTweet")
	photo, err := GetPhoto(diary.config)
	if err != nil {
		return fmt.Errorf("Unable to reach CDN for photo: %v", err)
	}
	tweet := &DailyTweet{
		Body:  photo.Title,
		Photo: photo,
	}
	return SendTweet(diary.config, tweet)
}
