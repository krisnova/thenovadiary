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

func DailyPhotoTweet(diary *Diary) error {
	logger.Debug("Running DailyPhotoTweet")

	// Connecting to Google Photos
	// Looking on disk for a local filestore

	//
	// TODO Nova start twitter implementation here!
	//
	return nil
}
