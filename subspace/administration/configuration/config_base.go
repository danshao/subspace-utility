package configuration

import (
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"github.com/jinzhu/gorm"
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
}

type ConfigBase struct {
	ConfigSchemaVersion uint `yaml:"config_schema_version"`
}

func (c *ConfigBase) GetConfigSchemaVersion() uint {
	return c.ConfigSchemaVersion
}
