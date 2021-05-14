package sorter

import (
	"github.com/XiovV/centralog/collector"
	"sort"
)

func GetNewBatch() []collector.Log {
	instances := []collector.Instance{{
		ID:            "494634d1b9eac2dcf341a237c4bb6948efabc4ce30411e234bbc3eb62095ed02",
		Name:          "autolog1",
		LastTimestamp: 1621023246,
	},{
		ID:            "f827226a0040d05c5e6d5fb1743f485a920f25a6cf48a974093b38dceb331ea6",
		Name:          "autolog2",
		LastTimestamp: 1621023247,
	},{
		ID:            "db0b87d1530404ff5a9809c4d15f64cf8b2c3146c04ba69a06d822c8bc9b8eba",
		Name:          "autolog3",
		LastTimestamp: 1621023247,
	},
	}

	var logs []collector.Log
	for _, instance := range instances {
		collectedLogs := instance.CollectLogs()

		for _, log := range collectedLogs {
			logs = append(logs, log)
		}
	}

	sort.SliceStable(logs, func(i, j int) bool {
		return logs[i].Timestamp < logs[j].Timestamp
	})

	return logs
}
