package configuration

import (
	"time"
	"crypto/sha1"
	"bytes"
	"encoding/binary"
	"fmt"
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

	SubspaceVersion     string
	SubspaceBuildNumber uint

	VpnServerVersion     string
	VpnServerBuildNumber uint

	UserSchemaVersion uint
	Users             []UserV1

	ProfileSchemaVersion uint
	Profiles             []ProfileV1
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
