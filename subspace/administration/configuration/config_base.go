package configuration

import (
	"time"

	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
)

type VersionedConfig interface {
	GetConfigSchemaVersion() uint
}

type SubspaceConfig interface {
	CalculateCheckSum() string
	IsCheckSumMatch() bool
	IsValid(db *gorm.DB) bool
	Validate(db *gorm.DB) error
	GetSystem() model.System
	GetUsers() []model.User
	GetProfiles() []model.Profile
	GetPolicies() []model.Policy
	GetPolicyRules() []model.PolicyRule
	GetProfilesPolicies() []model.ProfilesPolicy
	GetConfigCreateTime() time.Time
}

type ConfigBase struct {
	ConfigSchemaVersion uint      `yaml:"config_schema_version"`
	CreatedTime         time.Time `yaml:"created_time"`
}

func (c ConfigBase) GetConfigSchemaVersion() uint {
	return c.ConfigSchemaVersion
}

func (c ConfigBase) GetConfigCreateTime() time.Time {
	return c.CreatedTime
}
