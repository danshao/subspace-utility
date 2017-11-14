package configuration

import (
	"errors"
	"time"
)

type PolicyV1 struct {
	ID          uint      `yaml:"id"`
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	UpdatedAt   time.Time `yaml:"updated_date"`
	CreatedAt   time.Time `yaml:"created_date"`
}

func (policy *PolicyV1) Validate() error {
	if 0 >= policy.ID {
		return errors.New("User id must > 0.")
	}

	if "" == policy.Name {
		return errors.New("Name cannot be empty.")
	}

	if policy.CreatedAt.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if policy.UpdatedAt.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
