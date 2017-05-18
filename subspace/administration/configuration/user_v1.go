package configuration

import (
	"time"
)

type UserV1 struct {
	Id               uint
	Email            string
	Alias            string
	Role             string
	EmailVerified    bool
	Enabled          bool
	PasswordHash     string
	SetPasswordToken string

	RevokedDate   *time.Time  `yaml:",omitempty"`
	LastLoginDate *time.Time  `yaml:",omitempty"`
	UpdatedDate   *time.Time
	CreatedDate   *time.Time
}
