package model

import (
	"fmt"
	"net"
	"time"
)

type PolicyRule struct {
	ID                uint      `gorm:"column:id"`
	PolicyID          uint      `gorm:"column:policy_id"`
	Priority          uint      `gorm:"column:priority"`
	TargetDestination string    `gorm:"column:target_destination"`
	Action            string    `gorm:"column:action"`
	CreatedAt         time.Time `gorm:"column:created_date"`
	UpdatedAt         time.Time `gorm:"column:updated_date"`

	Policy Policy
}

func (PolicyRule) TableName() string {
	return "policy_rules"
}

func (p PolicyRule) ParseTargetDestination() (string, string) {
	ip, ipnet, _ := net.ParseCIDR(p.TargetDestination)
	netmask := fmt.Sprintf("%d.%d.%d.%d", ipnet.Mask[0], ipnet.Mask[1], ipnet.Mask[2], ipnet.Mask[3])
	return ip.String(), netmask
}
