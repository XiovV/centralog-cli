package collector

type Log struct {
	Timestamp int64 `json:"timestamp"`
	LogMessage string `json:"log_message"`
}
