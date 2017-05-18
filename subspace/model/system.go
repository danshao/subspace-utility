package model

import "time"

type System struct {
	Restriction             string    `gorm:"column:id"`
	SubspaceVersion         string    `gorm:"column:subspace_version"`
	SubspaceBuildNumber     uint      `gorm:"column:subspace_build_number"`
	VpnServerVersion        string    `gorm:"column:vpn_server_version"`
	VpnServerBuildNumber    uint      `gorm:"column:vpn_server_build_number"`
	Ip                      string    `gorm:"column:ip"`
	IpUpdatedDate           time.Time `gorm:"column:ip_updated_date"`
	Host                    string    `gorm:"column:host"`
	HostUpdatedDate         time.Time `gorm:"column:host_updated_date"`
	PreSharedKey            string    `gorm:"column:pre_shared_key"`
	PreSharedKeyUpdatedDate time.Time `gorm:"column:pre_shared_key_updated_date"`
	Uuid                    string    `gorm:"column:uuid"`
	UuidUpdatedDate         time.Time `gorm:"column:uuid_updated_date"`
	UserSchemaVersion       uint      `gorm:"column:user_schema_version"`
	ProfileSchemaVersion    uint      `gorm:"column:profile_schema_version"`
	ConfigSchemaVersion     uint      `gorm:"column:config_schema_version"`
	UpdatedDate             time.Time `gorm:"column:updated_date"`
	CreatedAt               time.Time `gorm:"column:created_at"`
}

func (System) TableName() string {
	return "system"
}