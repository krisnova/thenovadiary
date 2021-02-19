package main

import (
	"fmt"
	"os"

	"github.com/kris-nova/thenovadiary"
	"github.com/urfave/cli"

	"github.com/kris-nova/logger"
)

func GetApp() *cli.App {
	app := &cli.App{
		Name: "DiaryApplication",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "puser",
				Value:       "admin",
				Usage:       "Used to pass a Photoprism sername at runtime.",
				Destination: &DiaryConfig.PhotoprismUser,
			},
			&cli.StringFlag{
				Name:        "ppass",
				Value:       "",
				Usage:       "Used to pass a Photoprism password secret at runtime.",
				Destination: &DiaryConfig.PhotoprismPass,
			},
			&cli.StringFlag{
				Name:        "pconn",
				Value:       "",
				Usage:       "Used to pass a Photoprism connection string at runtime.",
				Destination: &DiaryConfig.PhotoprismConn,
			},
			&cli.StringFlag{
				Name:        "palbum",
				Value:       "",
				Usage:       "Used to pass a Photoprism album id at runtime.",
				Destination: &DiaryConfig.PhotoprismAlbum,
			},
			&cli.StringFlag{
				Name:        "twtokensecret",
				Value:       "",
				Usage:       "Used to pass a Twitter token secret at runtime.",
				Destination: &DiaryConfig.TwitterTokenSecret,
			},
			&cli.StringFlag{
				Name:        "twtoken",
				Value:       "",
				Usage:       "Used to pass a Twitter token at runtime.",
				Destination: &DiaryConfig.TwitterToken,
			},
			&cli.StringFlag{
				Name:        "twconsumersecret",
				Value:       "",
				Usage:       "Used to pass a Twitter consumer key secret at runtime.",
				Destination: &DiaryConfig.TwitterConsumerKeySecret,
			},
			&cli.StringFlag{
				Name:        "twconsumer",
				Value:       "",
				Usage:       "Used to pass a Twitter consumer key at runtime.",
				Destination: &DiaryConfig.TwitterConsumerKey,
			},
			&cli.StringFlag{
				Name:        "name",
				Value:       "Nova",
				Usage:       "A unique name for this particular set of actions.",
				Destination: &DiaryConfig.Name,
			},
			&cli.IntFlag{
				Name:        "verbose",
				Value:       4,
				Usage:       "Log level (nova logger)",
				Destination: &logger.Level,
			},
		},
		Action: func(c *cli.Context) error {
			return RunDiary(DiaryConfig)
		},
	}
	return app
}

func main() {
	app := GetApp()
	err := app.Run(os.Args)
	if err != nil {
		logger.Critical("Error running application: %v", err.Error())
	}
}

var DiaryConfig = &thenovadiary.DiaryConfig{
	//
}

func RunDiary(cfg *thenovadiary.DiaryConfig) error {
	err := thenovadiary.ValidateConfig(cfg)
	if err != nil {
		return fmt.Errorf("Unable to init config: %v", err)
	}
	logger.Always("Running Diary Program [%s]", cfg.Name)

	// Start program here!
	app := thenovadiary.New(cfg)
	return app.Service()
}
