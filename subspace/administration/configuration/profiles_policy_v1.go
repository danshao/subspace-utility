package configuration

import (
	"errors"
	"time"
)

type ProfilesPolicyV1 struct {
	ID        uint      `yaml:"id"`
	ProfileID uint      `yaml:"profile_id"`
	PolicyID  uint      `yaml:"policy_id"`
	UpdatedAt time.Time `yaml:"updated_date"`
	CreatedAt time.Time `yaml:"created_date"`
}

func (profiles_policy *ProfilesPolicyV1) Validate() error {
	if 0 >= profiles_policy.ID {
		return errors.New("User id must > 0.")
	}

	if 0 >= profiles_policy.ProfileID {
		return errors.New("Profile id must > 0.")
	}

	if 0 >= profiles_policy.PolicyID {
		return errors.New("Policy id must > 0.")
	}

	if profiles_policy.CreatedAt.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if profiles_policy.UpdatedAt.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
