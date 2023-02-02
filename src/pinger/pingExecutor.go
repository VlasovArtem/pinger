package pinger

import (
	"github.com/VlasovArtem/pinger/src/config"
	"sync"
	"sync/atomic"
	"time"
)

type PingInfo struct {
	Config   config.Config
	Result   bool
	PingTime time.Time
}

type Executor interface {
	Start(
		configProvider func() config.Config,
		pingFunc func(runConfig config.Config) bool,
		runResultFunc func(info PingInfo),
	)
	Stop()
	Status() bool
}

type ExecutorImpl struct {
	activationFlag atomic.Bool
	lock           sync.Mutex
}

func NewStarter() Executor {
	return &ExecutorImpl{
		activationFlag: atomic.Bool{},
		lock:           sync.Mutex{},
	}
}

func (s *ExecutorImpl) Start(
	configProvider func() config.Config,
	pingFunc func(runConfig config.Config) bool,
	runResultFunc func(info PingInfo),
) {
	if !s.activationFlag.Load() {
		s.activationFlag.Store(true)
		go func() {
			for s.activationFlag.Load() {
				if !s.lock.TryLock() {
					continue
				}
				runConfig := configProvider()
				result := pingFunc(runConfig)
				runResultFunc(PingInfo{
					Config:   runConfig,
					Result:   result,
					PingTime: time.Now(),
				})
				s.lock.Unlock()
				time.Sleep(runConfig.Timeout)
			}
		}()
	}
}

func (s *ExecutorImpl) Stop() {
	s.activationFlag.Store(false)
}

func (s *ExecutorImpl) Status() bool {
	return s.activationFlag.Load()
}
