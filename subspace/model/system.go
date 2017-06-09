package model

import (
	"time"
)

type System struct {
	Restriction                     string      `gorm:"column:restriction"`
	InstanceId                      string      `gorm:"column:instance_id"`
	SubspaceVersion                 string      `gorm:"column:subspace_version"`
	SubspaceBuildNumber             uint        `gorm:"column:subspace_build_number"`
	VpnServerVersion                string      `gorm:"column:vpn_server_version"`
	VpnServerBuildNumber            uint        `gorm:"column:vpn_server_build_number"`
	VpnServerAdministrationPassword string      `gorm:"column:vpn_server_administration_password"`
	VpnServerAdministrationPort     uint        `gorm:"column:vpn_server_administration_port"`
	VpnHubName                      string      `gorm:"column:vpn_hub_name"`
	Ip                              string      `gorm:"column:ip"`
	IpUpdatedDate                   *time.Time  `gorm:"column:ip_updated_date"`
	Host                            string      `gorm:"column:host"`
	HostUpdatedDate                 *time.Time  `gorm:"column:host_updated_date"`
	PreSharedKey                    string      `gorm:"column:pre_shared_key"`
	PreSharedKeyUpdatedDate         *time.Time  `gorm:"column:pre_shared_key_updated_date"`
	Uuid                            string      `gorm:"column:uuid"`
	UuidUpdatedDate                 *time.Time  `gorm:"column:uuid_updated_date"`
	SmtpHost                        string      `gorm:"column:smtp_host"`
	SmtpPort                        uint        `gorm:"column:smtp_port"`
	SmtpAuthentication              bool        `gorm:"column:smtp_authentication"`
	SmtpUsername                    string      `gorm:"column:smtp_username"`
	SmtpPassword                    string      `gorm:"column:smtp_password"`
	SmtpValid                       bool        `gorm:"column:smtp_valid"`
	SmtpSenderName                  string      `gorm:"column:smtp_sender_name"`
	SmtpSenderEmail                 string      `gorm:"column:smtp_sender_email"`
	UserSchemaVersion               uint        `gorm:"column:user_schema_version"`
	ProfileSchemaVersion            uint        `gorm:"column:profile_schema_version"`
	ConfigSchemaVersion             uint        `gorm:"column:config_schema_version"`
	UpdatedDate                     *time.Time  `gorm:"column:updated_date"`
	CreatedAt                       *time.Time  `gorm:"column:created_at"`
}

func (System) TableName() string {
	return "system"
}

/**
	Do NOT update following columns, it's depends on current subspace, database and vpn server:
		restriction,
    subspace_version,
    subspace_build_number,
    vpn_server_version,
    vpn_server_build_number,
    vpn_server_administration_port,
    vpn_server_administration_password,
    vpn_hub_name,
    ip,
    ip_updated_date,
    user_schema_version,
    profile_schema_version,
    config_schema_version,
*/
func (sys System) DataToRestore() map[string]interface{} {
	now := time.Now()
	return map[string]interface{}{
		"host":                        sys.Host,
		"host_updated_date":           now, // Set to true to trigger DNS check.
		"pre_shared_key":              sys.PreSharedKey,
		"pre_shared_key_updated_date": now,
		"uuid":                        sys.Uuid,
		"uuid_updated_date":           now,
		"smtp_host":                   sys.SmtpHost,
		"smtp_port":                   sys.SmtpPort,
		"smtp_authentication":         sys.SmtpAuthentication,
		"smtp_username":               sys.SmtpUsername,
		"smtp_password":               sys.SmtpPassword,
		"smtp_valid":                  sys.SmtpValid,
		"smtp_sender_name":            sys.SmtpSenderName,
		"smtp_sender_email":           sys.SmtpSenderEmail,
	}
}
