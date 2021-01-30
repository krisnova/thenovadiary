package thenovadiary

import "github.com/kris-nova/logger"

type DailyTweet struct {
	Body  string
	Photo *Photo
}

func SendTweet(cfg *DiaryConfig, dt *DailyTweet) error {
	logger.Info("Sending tweet")
	return nil
}
