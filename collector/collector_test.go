package collector

import "testing"

const BINARY_BUFFER = "00000000"

func TestParseLogString(t *testing.T) {
	tests := []struct{
		input string
		expected Log
	}{
		{BINARY_BUFFER + "2021-05-10T12:22:03.466390106Z 2021/05/10 12:22:03 http error: Access denied (err=Access denied to resource) (code=403)", Log{
			Timestamp:  1620649323,
			LogMessage: "2021/05/10 12:22:03 http error: Access denied (err=Access denied to resource) (code=403)",
		}},
		{BINARY_BUFFER + "2021-05-05T09:54:51.680968604Z 2021/05/05 09:54:51 server: Reverse tunnelling enabled", Log{
			Timestamp:  1620208491,
			LogMessage: "2021/05/05 09:54:51 server: Reverse tunnelling enabled",
		}},
		{BINARY_BUFFER + "2021-05-05T09:54:51.681008294Z 2021/05/05 09:54:51 server: Fingerprint 8c:15:0e:53:5f:ff:db:5c:22:7f:07:03:55:f7:2e:11", Log{
			Timestamp:  1620208491,
			LogMessage: "2021/05/05 09:54:51 server: Fingerprint 8c:15:0e:53:5f:ff:db:5c:22:7f:07:03:55:f7:2e:11",
		}},
		{BINARY_BUFFER + "2021-05-10T08:50:35.738716896Z [cont-init.d] 30-keygen: executing...", Log{
			Timestamp:  1620636635,
			LogMessage: "[cont-init.d] 30-keygen: executing...",
		}},
		{BINARY_BUFFER + "2021-05-10T08:50:36.216897110Z Setting permissions", Log{
			Timestamp:  1620636636,
			LogMessage: "Setting permissions",
		}},
		{BINARY_BUFFER + "2021-05-05T09:03:26.545186727Z writing new private key to '/config/keys/cert.key'", Log{
			Timestamp:  1620205406,
			LogMessage: "writing new private key to '/config/keys/cert.key'",
		}},
	}

	for _, tt := range tests {
		parsed := parseLogString(tt.input)

		if parsed.Timestamp != tt.expected.Timestamp {
			t.Errorf("timestamp is not correct. got=%d expected %d", parsed.Timestamp, tt.expected.Timestamp)
		}

		if parsed.LogMessage != tt.expected.LogMessage {
			t.Errorf("log message is not correct. got=%s expected %s", parsed.LogMessage, tt.expected.LogMessage)
		}
	}
}
