package compute

import "time"

func TimeStamp() string {
	return time.Now().In(time.UTC).Format(time.RFC3339)
}
