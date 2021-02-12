package thenovadiary

import (
	"sync"
	"time"

	"github.com/kris-nova/logger"
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
		cache:  NewCache(cfg.Name),
		config: cfg,
	}
}

func (d *Diary) Service() error {
	logger.Always("Starting service...")
	run := true
	delta, err := d.cache.Recover()
	logger.Info("Delta found: %d", delta)
	if err != nil {
		logger.Info("Unable to recover cache %s, starting with empty cache: %v", d.cache.path.Name(), err)
	} else {
		logger.Info("Successful cache recovery from %s", d.cache.path.Name())
	}
	for run {
		d.lock.Lock()
		// ----------------------------------
		{
			logger.Info("Running...")
			// TODO @kris-nova please remove this
			time.Sleep(2 * time.Second)

		}
		// ----------------------------------
		d.cache.Persist()
		d.lock.Unlock()
	}
	return nil

}
