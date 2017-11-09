package model

import "time"

type Policy struct {
	Id          uint      `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedDate time.Time `gorm:"column:created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
}

func (Policy) TableName() string {
	return "policies"
}
