package thenovadiary

import (
	"fmt"
)

type DiaryConfig struct {
	Name               string // Unique identifier for logs
	TwitterUser        string
	TwitterPass        string
	GoogleClientID     string
	GoogleClientSecret string
	//GoogleUser         string
	//GooglePass         string
	ActionString string
	Actions      []Action
	validated    bool
}

type Diary struct {
	config *DiaryConfig
}

func New(cfg *DiaryConfig) *Diary {
	return &Diary{
		config: cfg,
	}
}

func (d *Diary) ExecuteActions() error {
	if d.config.validated != true {
		return fmt.Errorf("Please call ValidateConfig() before attempting to run actions")
	}
	for _, action := range d.config.Actions {
		err := action(d)
		if err != nil {
			// Break
			return fmt.Errorf("Critical error, halting action processing: %v", err)
		}
	}
	return nil
}
