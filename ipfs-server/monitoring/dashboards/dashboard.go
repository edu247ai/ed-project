package dashboards

import (
	"encoding/json"
	"net/http"

	"ipfs-server/monitoring/metrics"
)

type Dashboard struct {
	IngestCount   *metrics.Counter
	DownloadCount *metrics.Counter
	AvgIngestTime *metrics.Timer
}

func NewDashboard(ingest *metrics.Counter, download *metrics.Counter, ingestTimer *metrics.Timer) *Dashboard {
	return &Dashboard{
		IngestCount:   ingest,
		DownloadCount: download,
		AvgIngestTime: ingestTimer,
	}
}

func (d *Dashboard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	snap := map[string]interface{}{
		"ingest_count":    d.IngestCount.Value(),
		"download_count":  d.DownloadCount.Value(),
		"avg_ingest_time": d.AvgIngestTime.Avg().String(),
	}
	json.NewEncoder(w).Encode(snap)
}
