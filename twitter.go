package thenovadiary

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kris-nova/photoprism-client-go/api/v1"
)

const (

	// BypassTwitter
	// This Flag is used to DEBUG the program
	// and will simulate a successful API transaction
	// without actually hitting the Twitter public
	// API
	BypassTwitter = false

	// Default Twitter Values
	// These are constants and if
	// we every change our app account
	// we should update here
	DefaultTwitterHandle = "thenovadiary"
	DefaultTwitterURL    = "https://twitter.com/thenovadiary"
)

// SendPhotoTweet will return the URL of the sent tweet and/or an
// error.
func SendPhotoTweet(twitter *anaconda.TwitterApi, photo api.Photo, pBytes []byte) (string, error) {
	if BypassTwitter == true {
		return "", nil
	}
	// TODO We should be memory conscience here
	b64str := string(b64e(pBytes))

	// ------ [ Upload Photo ] ------
	media, err := twitter.UploadMedia(b64str)
	if err != nil {
		return "", fmt.Errorf("unable to upload photo to twitter: %s", err)
	}
	// ------ [ Send Tweet ] ------
	v := url.Values{}
	v.Set("media_ids", media.MediaIDString)
	notes := GetNotes(photo)

	tweet, err := twitter.PostTweet(status, v)
	if err != nil {
		return DefaultTwitterURL, fmt.Errorf("unable to tweet: %v", err)
	}
	url := fmt.Sprintf("https://twitter.com/%s/status/%s", DefaultTwitterHandle, tweet.IdStr)
	return url, nil
}
func b64e(message []byte) []byte {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return b
}

func b64d(message []byte) (b []byte, err error) {
	var l int
	b = make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	l, err = base64.StdEncoding.Decode(b, message)
	if err != nil {
		return
	}
	return b[:l], nil
}
