package configuration

import (
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
)

type VersionedConfig interface {
	GetConfigSchemaVersion() uint
}

type SubspaceConfig interface {
	CalculateCheckSum() string
	IsCheckSumMatch() bool
	IsValid() bool
	Validate() error
	GetSystem() model.System
	GetUsers() []model.User
	GetProfiles() []model.Profile
}

type ConfigBase struct {
	ConfigSchemaVersion uint
}

func (c *ConfigBase) GetConfigSchemaVersion() uint {
	return c.ConfigSchemaVersion
}
