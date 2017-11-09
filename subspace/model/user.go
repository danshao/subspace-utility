package model

import (
	"time"
)

type User struct {
	Id               uint       `gorm:"column:id"`
	Email            string     `gorm:"column:email"`
	Alias            string     `gorm:"column:alias"`
	Role             string     `gorm:"column:role"`
	EmailVerified    bool       `gorm:"column:email_verified"`
	Enabled          bool       `gorm:"column:enabled"`
	PasswordHash     string     `gorm:"column:password_hash"`
	SetPasswordToken string     `gorm:"column:set_password_token"`
	RevokedDate      *time.Time `gorm:"column:revoked_date"`
	LastLoginDate    *time.Time `gorm:"column:last_login_date"`
	CreatedDate      time.Time  `gorm:"column:created_date"`
	UpdatedDate      time.Time  `gorm:"column:updated_date"`
}

func (User) TableName() string {
	return "users"
}
