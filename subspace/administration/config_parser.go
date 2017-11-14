package administration

import (
	"errors"

	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration/configuration"
	"gopkg.in/yaml.v2"
)

func ParseConfig(yamlData []byte) (configuration.SubspaceConfig, error) {
	// Detect config version
	c := configuration.ConfigBase{}
	if err := yaml.Unmarshal(yamlData, &c); nil != err {
		return nil, err
	}

	//TODO maybe use reflect?

	version := c.GetConfigSchemaVersion()
	switch version {
	case 0: // Using in version do not want to support
		return nil, errors.New("Unsupport version 0.")
	default:
		cfg := configuration.ConfigV1{}
		error := yaml.Unmarshal(yamlData, &cfg)
		return cfg, error
	}
}
