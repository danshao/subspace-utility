package config

import (
	"fmt"
)

const (
	REDIS_DEFAULT_HOST      = "localhost"
	REDIS_SESSIONS_DB       = 1
)

func GetDefaultRedisUri() string {
	return FormatRedisUri(REDIS_DEFAULT_HOST)
}

func FormatRedisUri(host string) string {
	return fmt.Sprintf("%s:6379", host)
}