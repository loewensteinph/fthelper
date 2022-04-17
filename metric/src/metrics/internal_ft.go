package metrics

import (
	"github.com/loewensteinph/fthelper/metric/v4/src/clients"
	"github.com/loewensteinph/fthelper/metric/v4/src/collectors"
	"github.com/loewensteinph/fthelper/metric/v4/src/connection"
	"github.com/loewensteinph/fthelper/shared/caches"
	"github.com/loewensteinph/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

func callerClusterBuilder(desc *prometheus.Desc, cacheKey, cluster string) []prometheus.Metric {
	var cache = caches.Global
	var data = cache.Get(caches.Join(cacheKey, cluster))
	var metric = 0
	if data.IsExist() {
		metric = data.Data.(int)
	}

	return []prometheus.Metric{prometheus.MustNewConstMetric(
		desc,
		prometheus.CounterValue,
		float64(metric),
		cluster,
	)}
}

var InternalFT = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "http_call"),
			"How many time do we call freqtrade apis",
			[]string{"cluster"},
			prometheus.Labels{
				"type": "total",
			},
		),
		func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerClusterBuilder(desc, clients.HTTP_CALL, connector.Cluster())
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "http_call"),
			"How many time do we call freqtrade apis",
			[]string{"cluster"},
			prometheus.Labels{
				"type": "success",
			},
		),
		func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerClusterBuilder(desc, clients.HTTP_SUCCESS, connector.Cluster())
		},
	),
)
