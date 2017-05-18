package configuration

import (
	"time"
)

type ProfileV1 struct {
	Id           uint
	Hub          string
	UserName     string
	UserId       uint
	Description  string
	Enabled      bool
	LoginCount   uint
	VpnHost      string
	PreSharedKey string
	PasswordHash string

	//IncomingBytes uint
	//OutgoingBytes uint

	RevokedDate   *time.Time  `yaml:"RevokedDate,omitempty"`
	LastLoginDate *time.Time  `yaml:"LastLoginDate,omitempty"`
	UpdatedDate   *time.Time
	CreatedDate   *time.Time
}
