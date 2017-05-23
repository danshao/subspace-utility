package model

import (
	"time"
)

//Generated using JSON 2 Go struct tool: https://mholt.github.io/json-to-go/

type Profile struct {
	Id           uint   `json:"id" gorm:"column:id; primary_key:yes; type: int(11) unsigned NOT NULL AUTO_INCREMENT"`
	Hub          string `json:"hub" grom:"column:hub; unique_index:hub_username_unique"`
	UserName     string `json:"username" gorm:"column:username; unique_index:hub_username_unique"`
	UserId       uint   `json:"user_id" gorm:"column:user_id" sql:"type: int(11) unsigned NOT NULL"`
	Description  string `json:"description" gorm:"column:description"`
	Enabled      bool   `json:"enabled" gorm:"column:enabled"`
	LoginCount   uint   `json:"login_count" sql:"column:login_count; type: int(11) unsigned NOT NULL"`
	VpnHost      string `json:"hub" gorm:"column:vpn_host"`
	PreSharedKey string `json:"pre_shared_key" gorm:"column:pre_shared_key"`
	PasswordHash string `json:"password_hash" gorm:"column:password_hash"`

	IncomingBytes uint64 `json:"incoming_bytes" sql:"column:incoming_bytes; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	OutgoingBytes uint64 `json:"outgoing_bytes" sql:"column:outgoing_bytes; type: decimal(65,0) unsigned NOT NULL; default: '0'"`

	RevokedDate   time.Time `json:"revoked_date" sql:"column:revoked_date; type: datetime NOT NULL"`
	LastLoginDate time.Time `json:"last_login_date" sql:"column:last_login_date; type: datetime NOT NULL"`
	CreatedDate   time.Time `json:"created_date" sql:"column:created_date; type: datetime NOT NULL"`
	UpdatedDate   time.Time `json:"updated_date" sql:"column:updated_date; type: datetime NOT NULL"`
}

func (Profile) TableName() string {
	return "profiles"
}
