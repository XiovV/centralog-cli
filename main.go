package main

import (
	"fmt"
	"github.com/XiovV/centralog/collector"
)


func main() {
	logs := collector.CollectLogs("1363ca3a512c62af6772c0a82c6b3ef9bf42d67d977d1ffe2c98d89c526925d5")

	for _, log := range logs {
		fmt.Printf("timestamp: %d log: %s\n", log.Timestamp, log.LogMessage)
	}
}
