package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"regexp"
	"bytes"
	"strings"
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

type EnumResult struct {
	Type string
}


func GetAcceptableRoles(db *gorm.DB, tableName string, columnName string) ([]string, error) {
	results := make([]string, 0)
	enumResult := EnumResult{}
	db.Raw(fmt.Sprintf("DESC %s %s", tableName, columnName)).Scan(&enumResult)
	if nil != db.Error {
		return results, db.Error
	}

	re, _ := regexp.Compile(`'([^']+)'`)
	res := re.FindAllStringSubmatch(enumResult.Type, -1)
	for _, group := range res {
		results = append(results, group[1])
	}
	return results, nil
}


func UnlockTable(db *gorm.DB) error {
	db.Raw("UNLOCK TABLES")
	return db.Error
}


func LockTableWrite(db *gorm.DB, tableNames ...string) error {
	lockTableSql := formatWriteLockTables(tableNames)
	db.Raw(lockTableSql)
	return db.Error
}

func formatWriteLockTables(tableNames []string) string {
	var buffer bytes.Buffer

	buffer.WriteString("LOCK TABLES")
	for index, name := range tableNames {
		buffer.WriteString(fmt.Sprintf(" %s WRITE", name))
		if index < len(tableNames) - 1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(";")
	return buffer.String()
}


func TruncateTable(db *gorm.DB, tableNames ...string) error {
	sqlStatements := make([]string, 0)
	for _, tableName := range tableNames {
		sqlStatements = append(sqlStatements, fmt.Sprintf("TRUNCATE TABLE %s;", tableName))
	}

	allSql := strings.Join(sqlStatements, " ")
	db.Raw(allSql)
	return db.Error
}

