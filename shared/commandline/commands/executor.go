package commands

import (
	"github.com/loewensteinph/fthelper/shared/caches"
	"github.com/loewensteinph/fthelper/shared/commandline/models"
	"github.com/loewensteinph/fthelper/shared/loggers"
	"github.com/loewensteinph/fthelper/shared/maps"
)

type ExecutorParameter struct {
	Name   string
	Meta   *models.Metadata
	Config maps.Mapper
	Cache  *caches.Service
	Logger *loggers.Logger
	Args   []string
}

type Executor func(p *ExecutorParameter) error
