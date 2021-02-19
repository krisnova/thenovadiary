package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/photoprism-client-go"
	"github.com/kris-nova/thenovadiary"
)

func main() {
	logger.Level = 4
	cfg := &thenovadiary.DiaryConfig{}
	err := thenovadiary.ValidateConfig(cfg)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}

	client := photoprism.New(cfg.PhotoprismConn)
	err = client.Auth(photoprism.NewClientAuthLogin(cfg.PhotoprismUser, cfg.PhotoprismPass))
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(2)
	}
	photo, err := thenovadiary.FindNextPhotoInAlbum(client, cfg.PhotoprismAlbum)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(3)
	}
	jBytes, err := json.Marshal(photo)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(4)
	}
	fmt.Println(string(jBytes))
}
