package plugins

import (
	"github.com/loewensteinph/fthelper/shared/commandline/commands"
	"github.com/loewensteinph/fthelper/shared/commandline/flags"
	"github.com/loewensteinph/fthelper/shared/commandline/hooks"
	"github.com/loewensteinph/fthelper/shared/commandline/models"
	"github.com/loewensteinph/fthelper/shared/loggers"
	"github.com/loewensteinph/fthelper/shared/maps"
)

type PluginParameter struct {
	Metadata models.Metadata

	NewCommand commands.Creator
	NewFlags   flags.Creator
	NewHook    hooks.Creator

	Config maps.Mapper
	Logger *loggers.Logger
}

type Plugin func(p *PluginParameter) error
