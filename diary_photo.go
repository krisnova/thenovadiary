package thenovadiary

import (
	"github.com/kris-nova/logger"
)

type Photo struct {
	Data  []byte
	Title string
}

//		  Client
//
//      (INTERNET)
//         ^ v
//    { API Surface }
//		    lib
//         db/fs
//    mysql     /photos

// GetPhoto will pull a photo from the configured
// filestore
//
// TODO We need implement some sort of hashing strategy
// TODO here such that we are maximizing the photos based on
// TODO hierarchy
func GetPhoto(cfg *DiaryConfig) (*Photo, error) {
	logger.Info("Requesting Photo from CDN")

	photo := &Photo{}

	return photo, nil
}
