package configuration

import (
	"errors"
	"time"
)

type PolicyV1 struct {
	Id          uint      `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	UpdatedDate time.Time `yaml:"updated_date"`
	CreatedDate time.Time `yaml:"created_date"`
}

func (policy *PolicyV1) Validate(acceptRoles []string) error {
	if 0 >= policy.Id {
		return errors.New("User id must > 0.")
	}

	if "" == policy.Name {
		return errors.New("Name cannot be empty.")
	}

	if policy.CreatedDate.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if policy.UpdatedDate.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
