package plugins

import (
	"github.com/loewensteinph/fthelper/generator/v4/src/clusters"
	"github.com/loewensteinph/fthelper/shared/errors"
	"github.com/loewensteinph/fthelper/shared/fs"
	"github.com/loewensteinph/fthelper/shared/maps"
	"github.com/loewensteinph/fthelper/shared/runners"
)

func Create(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		var output, err = fs.Build(fs.ToObject(p.Data.Zi("output"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		var errs = errors.New()
		for _, f := range output.All() {
			errs.And(f.Build())
		}

		return errs.Error()
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
