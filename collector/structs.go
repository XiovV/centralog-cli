package collector

type Log struct {
	Timestamp int64 `json:"timestamp"`
	LogMessage string `json:"log_message"`
	Instance string `json:"instance"`
}

type Instance struct {
	ID string
	Name string
	LastTimestamp int64
}