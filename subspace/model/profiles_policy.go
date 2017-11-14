package model

import "time"

type ProfilesPolicy struct {
	ID        uint      `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_date"`
	UpdatedAt time.Time `gorm:"column:updated_date"`

	Profile   Profile
	ProfileID uint `gorm:"column:profile_id"`

	Policy   Policy
	PolicyID uint `gorm:"column:policy_id"`
}

func (ProfilesPolicy) TableName() string {
	return "profiles_policies"
}
