package configuration

import (
	"time"
	"crypto/sha1"
	"bytes"
	"encoding/binary"
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils"
	"errors"
	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
)

const SALT = "5ubSp4ce@EcoworkInc"

type ConfigV1 struct {
	ConfigSchemaVersion uint
	CreatedTime         time.Time
	InstanceId          string
	Uuid                string
	Host                string
	Ip                  string
	PreSharedKey        string
	CheckSum            string

	SmtpHost     string
	SmtpPort     uint
	SmtpUsername string
	SmtpPassword string
	SmtpValid    bool

	SubspaceVersion     string
	SubspaceBuildNumber uint

	VpnServerVersion     string
	VpnServerBuildNumber uint

	UserSchemaVersion uint
	Users             []UserV1

	ProfileSchemaVersion uint
	Profiles             []ProfileV1
}

func (c *ConfigV1) GetConfigSchemaVersion() uint {
	return c.ConfigSchemaVersion
}

func (c ConfigV1) CalculateCheckSum() string {
	c.CheckSum = ""
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, []byte(SALT))
	binary.Write(&buffer, binary.BigEndian, c)
	sum := fmt.Sprintf("% x", sha1.Sum(buffer.Bytes()))
	return sum
}

func (c ConfigV1) IsCheckSumMatch() bool {
	original := c.CheckSum
	chkSum := c.CalculateCheckSum()
	return chkSum == original
}

func (c ConfigV1) IsValid() bool {
	return nil == c.Validate()
}

func (c ConfigV1) GetSystem() model.System {
	sys := ParseSystemFromConfigV1(c)
	return sys
}

func (c ConfigV1) GetUsers() []model.User {
	users := make([]model.User, 0)
	for _, u := range c.Users {
		users = append(users, ParseUserV1(u))
	}
	return users
}

func (c ConfigV1) GetProfiles() []model.Profile {
	profiles := make([]model.Profile, 0)
	for _, p := range c.Profiles {
		profiles = append(profiles, ParseProfileV1(p))
	}
	return profiles
}

func (c ConfigV1) Validate() error {
	if utils.IsValidPreSharedKey(c.PreSharedKey) {
		return errors.New("Pre-shared key is invalid.")
	}

	if "" == c.Host || utils.IsValidHost(c.Host) {
		return errors.New("Host is invalid.")
	}

	if "" == c.Ip || utils.IsValidIp(c.Ip) {
		return errors.New("IP is invalid.")
	}

	if "" == c.Uuid || utils.IsValidUuidV4(c.Uuid) {
		return errors.New("UUID is not valid.")
	}

	// Validate users. Fetch accept enum values from database.
	db, err := gorm.Open("mysql", config.GetDefaultMysqlUri())
	defer db.Close()
	if nil != err {
		return errors.New("Cannot open database connection to get user role.")
	}

	roles, err := utils.GetAcceptableRoles(db, "users", "role")
	if nil != err {
		return errors.New("Cannot read users table role data")
	}

	for _, user := range c.Users {
		if err := user.Validate(roles); nil != err {
			return err
		}
	}

	// Validate profiles
	for _, profile := range c.Profiles {
		if err := profile.Validate(); nil != err {
			return err
		}
	}

	return nil
}
