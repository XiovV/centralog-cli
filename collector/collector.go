package collector

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"strings"
	"time"
	"unicode"
)

func New(ID string, name string) *Instance {
	return &Instance{
		ID:   ID,
		Name: name,
	}
}

func (i *Instance) CollectLogs() []Log {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Timestamps: true, Since: fmt.Sprintf("%d", i.LastTimestamp+1)}
	out, err := cli.ContainerLogs(ctx, i.ID, options)
	if err != nil {
		panic(err)
	}

	str := apiResponseToString(out)

	var logs []Log
	parseLogs(str, i.Name, &logs)

	i.LastTimestamp = logs[len(logs) - 1].Timestamp
	return logs
}

func apiResponseToString(out io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(out)

	return buf.String()
}

func parseLogs(logStr, containerName string, dest *[]Log) {
	var logs []string

	logs = strings.Split(logStr, "\n")

	var lg Log
	for _, log := range logs {
		split := strings.SplitN(log, " ", 2)

		if len(split) != 1 && validateLogMessage(split[0]) {
			lg = Log{
				Timestamp:  timestampToUnix(split[0]),
				LogMessage: split[1],
				Instance: containerName,
			}

			*dest = append(*dest, lg)
		}
	}
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

