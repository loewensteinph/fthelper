package plugins

import (
	"github.com/loewensteinph/fthelper/generator/v4/src/clusters"
	"github.com/loewensteinph/fthelper/shared/maps"
	"github.com/loewensteinph/fthelper/shared/runners"
)

func Empty(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		return nil
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
