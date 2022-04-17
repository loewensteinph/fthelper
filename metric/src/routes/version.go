package routes

import (
	"fmt"

	"github.com/loewensteinph/fthelper/metric/v4/src/connection"
	"github.com/loewensteinph/fthelper/shared/commandline/commands"
)

var Version = &Route{
	Path: "/version",
	Handler: func(p *commands.ExecutorParameter, connectors []connection.Connector) (int, interface{}) {
		return 200, fmt.Sprintf("%s: %s (%s)", p.Meta.Name, p.Meta.Version, p.Meta.Commit)
	},
}
