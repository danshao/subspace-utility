package configuration

import (
	"time"
	"errors"
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils/validator"
)

type UserV1 struct {
	Id               uint   `yaml:"id"`
	Email            string `yaml:"email"`
	Alias            string `yaml:"alias"`
	Role             string `yaml:"role"`
	EmailVerified    bool   `yaml:"email_verified"`
	Enabled          bool   `yaml:"enabled"`
	PasswordHash     string `yaml:"password_hash"`
	SetPasswordToken string `yaml:"set_password_token"`

	RevokedDate   *time.Time `yaml:"revoked_date,omitempty"`
	LastLoginDate *time.Time `yaml:"last_login_date,omitempty"`
	UpdatedDate   time.Time `yaml:"updated_date"`
	CreatedDate   time.Time `yaml:"created_date"`
}

func (user *UserV1) Validate(acceptRoles []string) error {
	if 0 >= user.Id {
		return errors.New("User id must > 0.")
	}

	if !validator.IsValidEmail(user.Email) {
		return errors.New("Email is invalid.")
	}

	if !validator.IsStringInArray(user.Role, acceptRoles) {
		return errors.New(fmt.Sprintf("Role is not in %v.", acceptRoles))
	}

	if "" == user.PasswordHash {
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
