package configuration

import (
	"errors"
	"time"
)

type ProfilesPolicyV1 struct {
	Id          uint      `yaml:"id"`
	ProfileId   uint      `yaml:"profile_id"`
	PolicyId    uint      `yaml:"policy_id"`
	UpdatedDate time.Time `yaml:"updated_date"`
	CreatedDate time.Time `yaml:"created_date"`
}

func (profiles_policy *ProfilesPolicyV1) Validate(acceptRoles []string) error {
	if 0 >= profiles_policy.Id {
		return errors.New("User id must > 0.")
	}

	if 0 >= profiles_policy.ProfileId {
		return errors.New("Profile id must > 0.")
	}

	if 0 >= profiles_policy.PolicyId {
		return errors.New("Policy id must > 0.")
	}

	if profiles_policy.CreatedDate.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if profiles_policy.UpdatedDate.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
