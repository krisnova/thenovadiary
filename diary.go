package thenovadiary

import (
	"sync"
	"time"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

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
	lock   sync.Mutex
	cache  *Cache
}

func New(cfg *DiaryConfig) *Diary {
	return &Diary{
		config: cfg,
	}
}

func (d *Diary) Service() error {
	logger.Always("Starting service...")
	cache := NewCache("Nova")
	run := true
	delta, err := cache.Recover()
	logger.Info("Delta found: %d", delta)
	if err != nil {
		logger.Info("Unable to recover cache %s, starting with empty cache: %v", cache.path.Name(), err)
	} else {
		logger.Info("Successful cache recovery from %s", cache.path.Name())
	}
	for run {
		d.lock.Lock()
		// ----------------------------------
		// 1)
		{
			logger.Debug("Sleeping...")
			time.Sleep(2 * time.Second)
		}
		// 2)
		// 3)
		// ----------------------------------
		cache.Persist()
		d.lock.Unlock()
	}

	return nil
}
