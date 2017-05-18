package config

import (
	"fmt"
)

const (
	MYSQL_URI_FORMAT       = "%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	MYSQL_DEFAULT_HOST     = "localhost"
	MYSQL_DEFAULT_ACCOUNT  = "subspace"
	MYSQL_DEFAULT_PASSWORD = "subspace"
	MYSQL_DEFAULT_DATABASE = "subspace"
)

func GetDefaultMysqlUri() string {
	return FormatMysqlUri(
		MYSQL_DEFAULT_HOST,
		MYSQL_DEFAULT_ACCOUNT,
		MYSQL_DEFAULT_PASSWORD,
		MYSQL_DEFAULT_DATABASE,
	)
}

func FormatMysqlUri(host string, account string, password string, databaseName string) string {
	return fmt.Sprintf(MYSQL_URI_FORMAT,
		account,
		password,
		host,
		databaseName,
	)
}
