package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

type IpString struct {
	Ip string
}

func GetReadableIp(db *gorm.DB, tableName string, columnName string) (string, error) {
	ipString := IpString{}
	db.Raw(fmt.Sprintf("SELECT INET6_NTOA(%s) AS ip FROM %s", columnName, tableName)).Scan(&ipString)
	if nil != db.Error {
		return "", db.Error
	}

	if "" == ipString.Ip {
		db.Raw(fmt.Sprintf("SELECT INET_NTOA(%s) AS ip FROM %s", columnName, tableName)).Scan(&ipString)
		if nil != db.Error {
			return "", db.Error
		} else {
			return ipString.Ip, nil
		}
	} else {
		return ipString.Ip, nil
	}
}

func IsIpV6(db *gorm.DB, tableName string, columnName string) (bool, error) {
	ipString := IpString{}
	db.Raw(fmt.Sprintf("SELECT INET6_NTOA(%s) AS ip FROM %s", columnName, tableName)).Scan(&ipString)
	if nil != db.Error {
		return false, db.Error
	}

	if "" == ipString.Ip {
		return false, nil
	} else {
		return true, nil
	}
}