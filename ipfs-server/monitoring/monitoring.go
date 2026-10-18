package monitoring

import (
	"ipfs-server/monitoring/dashboards"
	"ipfs-server/monitoring/events"
	"ipfs-server/monitoring/metrics"
)

type Monitoring struct {
	Metrics struct {
		IngestCount   *metrics.Counter
		DownloadCount *metrics.Counter
		IngestTimer   *metrics.Timer
	}
	Events    *events.Bus
	Dashboard *dashboards.Dashboard
}

func New() *Monitoring {
	m := &Monitoring{}

	m.Metrics.IngestCount = metrics.NewCounter("ingest_count")
	m.Metrics.DownloadCount = metrics.NewCounter("download_count")
	m.Metrics.IngestTimer = metrics.NewTimer("ingest_timer")

	m.Events = events.NewBus()

	m.Dashboard = dashboards.NewDashboard(
		m.Metrics.IngestCount,
		m.Metrics.DownloadCount,
		m.Metrics.IngestTimer,
	)

	return m
}
