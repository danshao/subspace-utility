package config

import (
	"fmt"
)

const (
	REDIS_DEFAULT_HOST        = "localhost"
	REDIS_VPN_SESSIONS_DB     = 1
	REDIS_WEB_SESSION_DB      = 2
	REDIS_ACCOUNTS_DB         = 3
	REDIS_BACKUP_SNAPSHOTS_DB = 4
)

func GetDefaultRedisUri() string {
	return FormatRedisUri(REDIS_DEFAULT_HOST)
}

func FormatRedisUri(host string) string {
	return fmt.Sprintf("%s:6379", host)
}
