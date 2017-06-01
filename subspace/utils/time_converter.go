package utils

import "time"

func ToUnixTimestampInMillisecond(t *time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
