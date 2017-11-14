package model

import "time"

type Policy struct {
	ID          uint      `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_date"`
	UpdatedAt   time.Time `gorm:"column:updated_date"`

	PolicyRules []PolicyRule
	Profiles    []Profile `gorm:"many2many:profiles_policies;"`
}

func (Policy) TableName() string {
	return "policies"
}
