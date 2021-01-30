package thenovadiary

import (
	"github.com/kris-nova/logger"
)

type Photo struct {
	Data  []byte
	Title string
}

func GetPhoto(cfg *DiaryConfig) (*Photo, error) {
	logger.Info("Requesting Photo from CDN")

	photo := &Photo{}

	return photo, nil
}
