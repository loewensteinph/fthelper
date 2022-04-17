module github.com/loewensteinph/fthelper/metric/v4

go 1.16

replace github.com/loewensteinph/fthelper/shared v0.0.0 => ../shared

require (
	github.com/loewensteinph/fthelper/shared v0.0.0
		github.com/jackc/pgx/v4 v4.14.1
	github.com/prometheus/client_golang v1.11.0
)
