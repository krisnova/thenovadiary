package main

import (
	"fmt"
	"os"

	"github.com/kris-nova/thenovadiary"

	"github.com/kris-nova/logger"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name: "DiaryApplication",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "puser",
				Value:       "admin",
				Usage:       "Used to pass a Photoprism sername at runtime.",
				Destination: &diaryConfig.PhotoprismUser,
			},
			&cli.StringFlag{
				Name:        "ppass",
				Value:       "",
				Usage:       "Used to pass a Photoprism password secret at runtime.",
				Destination: &diaryConfig.PhotoprismPass,
			},
			&cli.StringFlag{
				Name:        "twuser",
				Value:       "",
				Usage:       "Used to pass a Twitter username at runtime.",
				Destination: &diaryConfig.TwitterUser,
			},
			&cli.StringFlag{
				Name:        "twpass",
				Value:       "",
				Usage:       "Used to pass a Twitter username at runtime.",
				Destination: &diaryConfig.TwitterPass,
			},
			&cli.StringFlag{
				Name:        "actions",
				Value:       "daily",
				Usage:       fmt.Sprintf("Comma dilimited set of action strings: %s", thenovadiary.ActionsString()),
				Destination: &diaryConfig.ActionString,
			},
			&cli.StringFlag{
				Name:        "name",
				Value:       "daily",
				Usage:       "A unique name for this particular set of actions.",
				Destination: &diaryConfig.Name,
			},
			&cli.IntFlag{
				Name:        "verbose",
				Value:       4,
				Usage:       "Log level (nova logger)",
				Destination: &logger.Level,
			},
		},
		Action: func(c *cli.Context) error {
			return RunDiary(diaryConfig)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Critical("Error running application: %v", err.Error())
	}
}

var diaryConfig = &thenovadiary.DiaryConfig{
	//
}

func RunDiary(cfg *thenovadiary.DiaryConfig) error {
	err := thenovadiary.ValidateConfig(cfg)
	if err != nil {
		return fmt.Errorf("Unable to init config: %v", err)
	}
	logger.Always("Running Diary Program [%s]", cfg.Name)
	diary := thenovadiary.New(cfg)
	return diary.ExecuteActions()
}
