package configuration

import (
	"errors"
	"time"
)

type PolicyRuleV1 struct {
	ID                uint      `yaml:"id"`
	PolicyID          uint      `yaml:"policy_id"`
	Priority          uint      `yaml:"priority"`
	TargetDestination string    `yaml:"target_destination"`
	Action            string    `yaml:"action"`
	UpdatedAt         time.Time `yaml:"updated_date"`
	CreatedAt         time.Time `yaml:"created_date"`
}

func (policyRule *PolicyRuleV1) Validate() error {
	if 0 >= policyRule.ID {
		return errors.New("User id must > 0.")
	}

	if 0 >= policyRule.PolicyID {
		return errors.New("Profile id must > 0.")
	}

	if 0 >= policyRule.Priority {
		return errors.New("Priority must > 0.")
	}

	if "" == policyRule.TargetDestination {
		return errors.New("Target Destination cannot be empty.")
	}

	if "" == policyRule.Action {
		return errors.New("Action cannot be empty.")
	}

	if policyRule.CreatedAt.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if policyRule.UpdatedAt.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
