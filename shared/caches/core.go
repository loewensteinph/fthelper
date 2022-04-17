package caches

import (
	"sync"

	"github.com/loewensteinph/fthelper/shared/loggers"
)

func New() *Service {
	return &Service{
		caches: make(map[string]*Data),
		mutex:  sync.RWMutex{},
		logger: loggers.Get("cache", "service"),
	}
}

var Global = New()
