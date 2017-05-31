package vpn

import (
	"strings"
	"time"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils"
)

type Softether struct {
	PreSharedKey       string
	AdminPassword      string // Need hashed
	AdminPasswordHash  string
	Hub                Hub
	AdministrationPort uint
}

func (softether Softether) GetAdminPasswordHash() string {
	if "" != softether.AdminPasswordHash {
		return softether.AdminPasswordHash
	} else {
		if hash, err := GenerateSoftetherPasswordHash(softether.AdminPassword); nil == err {
			return hash
		} else {
			return ""
		}
	}
}

func (softether Softether) GetDefaultHub() string {
	return softether.Hub.Name
}

func (softether Softether) GetDefaultAdministrationPort() uint {
	return softether.AdministrationPort
}

type Hub struct {
	Name     string
	Accounts []Account
}

type Account struct {
	Username       string
	Password       string // Need hash
	PasswordHash   string
	NtLmSecureHash string
	RawRealName    string //RealName
	RealName       string
	RawNote        string //Note
	Note           string
	LoginCount     uint
	ExpireTime     time.Time // softether config expire time
	RevokedTime    time.Time // Database revoked_date - 1day = ExpireTime
	LastLoginTime  time.Time
	UpdatedTime    time.Time
	CreatedTime    time.Time
}

/**
Softether concat password and uppercase username as password hash.
If username is "Username" and password is "Password" than hash will be:
$ echo -n "PasswordUSERNAME" | openssl sha -binary | openssl base64
 */
func (ac Account) GetPasswordHash() string {
	if "" != ac.PasswordHash {
		return ac.PasswordHash
	} else {
		if hash, err := GenerateSoftetherPasswordHash(ac.Password + strings.ToUpper(ac.Username)); nil == err {
			return hash
		} else {
			//TODO panic?
			return ""
		}
	}
}

func (ac Account) GetNtLmSecureHash() string {
	if "" != ac.NtLmSecureHash {
		return ac.NtLmSecureHash
	} else {
		return GenerateNtLmPasswordHash(ac.Password)
	}
}

func (ac Account) GetRealName() string {
	if "" != ac.RealName {
		return ac.RealName
	} else {
		return GenerateSoftetherEncodedString(ac.RawRealName)
	}
}

func (ac Account) GetNote() string {
	if "" != ac.Note {
		return ac.Note
	} else {
		return GenerateSoftetherEncodedString(ac.RawNote)
	}
}

func (ac Account) GetCreatedTimeInMilliseconds() int64 {
	return utils.ToUnixTimestampInMillisecond(ac.CreatedTime)
}

func (ac Account) GetUpdatedTime() int64 {
	return utils.ToUnixTimestampInMillisecond(ac.UpdatedTime)
}

func (ac Account) GetExpireTime() int64 {
	if !ac.ExpireTime.IsZero() {
		return utils.ToUnixTimestampInMillisecond(ac.ExpireTime)
	} else if !ac.RevokedTime.IsZero() {
		return utils.ToUnixTimestampInMillisecond(ac.RevokedTime.AddDate(0, 0, -1))
	} else {
		return 0
	}
}

func (ac Account) GetLastLoginTime() int64 {
	return utils.ToUnixTimestampInMillisecond(ac.LastLoginTime)
}
