package configuration

import (
	"errors"
	"time"
)

type PolicyRuleV1 struct {
	Id                uint      `yaml:"id"`
	PolicyId          uint      `yaml:"policy_id"`
	Priority          uint      `yaml:"priority"`
	TargetDestination string    `yaml:"target_destination"`
	Action            string    `yaml:"action"`
	UpdatedDate       time.Time `yaml:"updated_date"`
	CreatedDate       time.Time `yaml:"created_date"`
}

func (policyRule *PolicyRuleV1) Validate(acceptRoles []string) error {
	if 0 >= policyRule.Id {
		return errors.New("User id must > 0.")
	}

	if 0 >= policyRule.PolicyId {
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

	if policyRule.CreatedDate.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if policyRule.UpdatedDate.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
