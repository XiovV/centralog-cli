package collector

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
	"time"
	"unicode"
)

func CollectLogs(containerId string) []Log {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: true}
	out, err := cli.ContainerLogs(ctx, containerId, options)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(out)
	str := buf.String()

	var splitLogs []string

	splitLogs = strings.Split(str, "\n")

	var logs []Log
	parseLogs(splitLogs, &logs)

	return logs
}

func parseLogs(logs []string, dest *[]Log) {
	for _, log := range logs {
		lg := parseLogString(log)
		*dest = append(*dest, lg)
	}
}

func parseLogString(log string) Log {
	split := strings.SplitN(log, " ", 2)

	if len(split) != 1 && validateLogMessage(split[0]) {
		return Log{
			Timestamp:  timestampToUnix(split[0]),
			LogMessage: split[1],
		}
	}

	return Log{}
}

func validateLogMessage(log string) bool {
	for _, r := range log {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func timestampToUnix(timestamp string) int64 {
	byteTimestamp := []byte(timestamp[8:])
	str := string(byteTimestamp)
	t, err := time.Parse("2006-01-02T15:04:05.000000000Z", str)
	if err != nil {
		panic(err)
	}
	return t.Unix()
}

