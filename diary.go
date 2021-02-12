package thenovadiary

import "github.com/kubicorn/kubicorn/pkg/logger"

type DiaryConfig struct {
	Name           string // Unique identifier for logs
	TwitterUser    string
	TwitterPass    string
	PhotoprismPass string
	PhotoprismUser string
	PhotoprismConn string
	validated      bool
}

type Diary struct {
	config *DiaryConfig
}

func New(cfg *DiaryConfig) *Diary {
	return &Diary{
		config: cfg,
	}
}

func (d *Diary) Service() error {
	logger.Always("Starting service...")

	return nil
}
