package configuration

import (
	"time"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils"
	"errors"
	"fmt"
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

	RevokedDate   time.Time  `yaml:",omitempty"`
	LastLoginDate time.Time  `yaml:",omitempty"`
	UpdatedDate   time.Time
	CreatedDate   time.Time
}

func (user *UserV1) Validate(acceptRoles []string) error {
	if 0 < user.Id {
		return errors.New("User id must >= 0.")
	}

	if utils.IsValidEmailFormat(user.Email) {
		return errors.New("Email is invalid.")
	}

	if utils.IsStringInArray(user.Role, acceptRoles) {
		return errors.New(fmt.Sprintf("Role is not in %v.", acceptRoles))
	}

	if "" != user.PasswordHash {
		return errors.New("Password hash cannot be empty.")
	}

	if user.CreatedDate.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if user.UpdatedDate.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}