package workers

import (
	"net"
	"sort"
)

type MXRecord struct {
	Host     string
	Priority uint16
}

func ResolveMX(domain string) ([]MXRecord, error) {
	mxs, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}

	records := make([]MXRecord, 0, len(mxs))
	for _, m := range mxs {
		records = append(records, MXRecord{
			Host:     m.Host,
			Priority: m.Pref,
		})
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Priority < records[j].Priority
	})

	return records, nil
}
