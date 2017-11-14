package configuration

import (
	"errors"
	"time"
)

type ProfileV1 struct {
	Id             uint   `yaml:"id"`
	Hub            string `yaml:"hub"`
	UserName       string `yaml:"username"`
	PasswordHash   string `yaml:"password_hash"`
	NtLmSecureHash string `yaml:"ntlm_secure_hash"`
	UserId         uint   `yaml:"user_id"`
	FullName       string `yaml:"full_name"`
	Description    string `yaml:"description"`
	Enabled        bool   `yaml:"enabled"`
	LoginCount     uint   `yaml:"login_count"`
	VpnHost        string `yaml:"vpn_host"`
	PreSharedKey   string `yaml:"pre_shared_key"`
	//IncomingBytes  uint       `yaml:"incoming_bytes"`
	//OutgoingBytes  uint       `yaml:"outgoing_bytes"`
	RevokedDate   *time.Time `yaml:"revoked_date,omitempty"`
	LastLoginDate *time.Time `yaml:"lastLogin_date,omitempty"`
	UpdatedDate   time.Time  `yaml:"updated_date"`
	CreatedDate   time.Time  `yaml:"created_date"`
}

func (profile *ProfileV1) Validate() error {
	if 0 >= profile.Id {
		return errors.New("Profile id must > 0.")
	}

	if 0 >= profile.UserId {
		return errors.New("Profile belongs user id must > 0.")
	}

	if "" == profile.Hub {
		return errors.New("Hub name cannot empty.")
	}

	if "" == profile.UserName {
		return errors.New("Profile user name cannot empty.")
	}

	if "" == profile.PasswordHash {
		return errors.New("Profile password hash cannot empty.")
	}

	if 0 > profile.LoginCount {
		return errors.New("Profile login count must >= 0.")
	}

	if profile.CreatedDate.IsZero() {
		return errors.New("CreateDate cannot be empty.")
	}

	if profile.UpdatedDate.IsZero() {
		return errors.New("UpdateDate cannot be empty.")
	}

	return nil
}
