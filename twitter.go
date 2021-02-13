package thenovadiary

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/kris-nova/photoprism-client-go/api/v1"
)

const (
	BypassTwitter = true
)

// SendPhotoTweet will return the URL of the sent tweet and/or an
// error.
func SendPhotoTweet(twitter *anaconda.TwitterApi, photo api.Photo, pBytes []byte) (string, error) {
	if BypassTwitter == true {
		return "", nil
	}
	return "", nil
}
