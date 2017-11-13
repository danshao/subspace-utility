package model

import (
	"time"
)

type Log struct {
	Id              uint       `gorm:"column:id"`
	Source          int        `gorm:"column:source"`
	Type            int        `gorm:"column:type"`
	LogTime         *time.Time `gorm:"column:log_time"`
	RawLog          string     `gorm:"column:raw_log"`
	UserId          uint       `gorm:"column:user_id"`
	Hub             string     `gorm:"column:hub"`
	ProfileUsername string     `gorm:"column:profile_username"`
	SessionName     string     `gorm:"column:session_name"`
	OperatorId      uint       `gorm:"column:operator_id"`
	ClientIp        string     `gorm:"column:client_ip"`
	CreatedDate     *time.Time `gorm:"column:created_date"`
}

func (Log) TableName() string {
	return "logs"
}
