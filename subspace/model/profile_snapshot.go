package model

import (
	"time"
)

//Generated using JSON 2 Go struct tool: https://mholt.github.io/json-to-go/

type ProfileSnapshot struct {
	// Additional data
	Id           uint `json:"id" gorm:"column:id; primary_key:yes; type: int(11) unsigned NOT NULL AUTO_INCREMENT"`
	Hub          string `json:"hub" grom:"column:hub; unique_index:profile_snapshot_unique"`
	SnapshotDate *time.Time `json:"snapshot_date" gorm:"column:snapshot_date; unique_index:profile_snapshot_unique" sql:"default: current_timestamp"`
	UserId       uint `json:"user_id" gorm:"column:user_id; unique_index:profile_snapshot_unique" sql:"type: int(11) unsigned NOT NULL"`

	// UserList and UserGet shared column
	UserName       string `json:"user_name" gorm:"column:user_name; unique_index:profile_snapshot_unique"`
	Description    string `json:"description" gorm:"column:description"`
	FullName       string `json:"full_name" gorm:"column:full_name"`
	AuthType       string `json:"auth_type" gorm:"column:auth_type"`
	ExpirationDate string `json:"expiration_date" gorm:"column:expiration_date"`
	NumberOfLogins uint `json:"number_of_logins" sql:"column:number_of_logins; type: int(11) unsigned NOT NULL"`

	//TODO maybe I don't need UserList
	// UserList
	//GroupName string `json:"Group Name"`
	//LastLogin string `json:"Last Login"`
	//TransferBytes big.Int `json:"Transfer Bytes"`
	//TransferPackets big.Int `json:"Transfer Packets"`

	// UserGet
	IncomingBroadcastPackets   uint64 `json:"incoming_broadcast_packets" sql:"column:incoming_broadcast_packets; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	IncomingBroadcastTotalSize uint64 `json:"incoming_broadcast_total_size" sql:"column:incoming_broadcast_total_size; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	IncomingUnicastPackets     uint64 `json:"incoming_unicast_packets" sql:"column:incoming_unicast_packets; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	IncomingUnicastTotalSize   uint64 `json:"incoming_unicast_total_size" sql:"column:incoming_unicast_total_size; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	OutgoingBroadcastPackets   uint64 `json:"outgoing_broadcast_packets" sql:"column:outgoing_broadcast_packets; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	OutgoingBroadcastTotalSize uint64 `json:"outgoing_broadcast_total_size" sql:"column:outgoing_broadcast_total_size; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	OutgoingUnicastPackets     uint64 `json:"outgoing_unicast_packets" sql:"column:outgoing_unicast_packets; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	OutgoingUnicastTotalSize   uint64 `json:"outgoing_unicast_total_size" sql:"column:outgoing_unicast_total_size; type: decimal(65,0) unsigned NOT NULL; default: '0'"`
	CreatedOn                  *time.Time `json:"created_on" sql:"column:created_on; type: datetime NOT NULL"`
	UpdatedOn                  *time.Time `json:"updated_on" sql:"column:updated_on; type: datetime NOT NULL"`
}

func (ProfileSnapshot) TableName() string {
	return "profile_snapshots"
}