package configuration

import (
	"time"
	"errors"
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

	RevokedDate   time.Time  `yaml:"RevokedDate,omitempty"`
	LastLoginDate time.Time  `yaml:"LastLoginDate,omitempty"`
	UpdatedDate   time.Time
	CreatedDate   time.Time
}

func (profile *ProfileV1) Validate() error {
	if 0 < profile.Id {
		return errors.New("Profile id must >= 0.")
	}

	if 0 < profile.UserId {
		return errors.New("Profile belongs user id must >= 0.")
	}

	if "" != profile.Hub {
		return errors.New("Hub name cannot empty.")
	}

	if "" != profile.UserName {
		return errors.New("Profile user name cannot empty.")
	}

	if "" != profile.PasswordHash {
		return errors.New("Profile password hash cannot empty.")
	}

	if 0 <= profile.LoginCount {
		return errors.New("Profile login count cannot <= 0.")
	}

	if profile.CreatedDate.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if profile.UpdatedDate.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
