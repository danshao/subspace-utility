package model

import "time"

type ProfilesPolicy struct {
	Id          uint      `gorm:"column:id"`
	ProfileId   uint      `gorm:"column:profile_id"`
	PolicyId    uint      `gorm:"column:policy_id"`
	CreatedDate time.Time `gorm:"column:created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date"`
}

func (ProfilesPolicy) TableName() string {
	return "profiles_policies"
}
