package model

import "time"

type PolicyRule struct {
	Id                uint      `gorm:"column:id"`
	PolicyId          uint      `gorm:"column:policy_id"`
	Priority          uint      `gorm:"column:priority"`
	TargetDestination string    `gorm:"column:target_destination"`
	Action            string    `gorm:"column:action"`
	CreatedDate       time.Time `gorm:"column:created_date"`
	UpdatedDate       time.Time `gorm:"column:updated_date"`
}

func (PolicyRule) TableName() string {
	return "policy_rules"
}
