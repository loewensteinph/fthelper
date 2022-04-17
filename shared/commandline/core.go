package commandline

import (
	"github.com/loewensteinph/fthelper/shared/caches"
	"github.com/loewensteinph/fthelper/shared/commandline/commands"
	"github.com/loewensteinph/fthelper/shared/commandline/flags"
	"github.com/loewensteinph/fthelper/shared/commandline/hooks"
	"github.com/loewensteinph/fthelper/shared/commandline/models"
	"github.com/loewensteinph/fthelper/shared/commandline/plugins"
	"github.com/loewensteinph/fthelper/shared/loggers"
)

func New(cache *caches.Service, metadata *models.Metadata) *cli {
	return &cli{
		Metadata: metadata,
		flags:    flags.New(),
		commands: commands.New(),
		hooks:    hooks.New(),
		plugins:  plugins.New(),

		cache:  cache,
		logger: loggers.Get("commandline"),
	}
}
