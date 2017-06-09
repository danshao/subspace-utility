package configuration

import (
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"time"
)

/**
	Convert System
 */
func ToConfigV1(system model.System, users []UserV1, profiles []ProfileV1) ConfigV1 {
	return ConfigV1{
		ConfigSchemaVersion:             system.ConfigSchemaVersion,
		CreatedTime:                     time.Now(),
		InstanceId:                      system.InstanceId,
		SubspaceVersion:                 system.SubspaceVersion,
		SubspaceBuildNumber:             system.SubspaceBuildNumber,
		VpnServerVersion:                system.VpnServerVersion,
		VpnServerBuildNumber:            system.VpnServerBuildNumber,
		VpnServerAdministrationPassword: system.VpnServerAdministrationPassword,
		VpnServerAdministrationPort:     system.VpnServerAdministrationPort,
		VpnHubName:                      system.VpnHubName,
		Ip:                              system.Ip,
		VpnHost:                         system.Host,
		PreSharedKey:                    system.PreSharedKey,
		Uuid:                            system.Uuid,
		SmtpHost:                        system.SmtpHost,
		SmtpPort:                        system.SmtpPort,
		SmtpAuthentication:              system.SmtpAuthentication,
		SmtpUsername:                    system.SmtpUsername,
		SmtpPassword:                    system.SmtpPassword,
		SmtpValid:                       system.SmtpValid,
		SmtpSenderName:                  system.SmtpSenderName,
		SmtpSenderEmail:                 system.SmtpSenderEmail,

		UserSchemaVersion: system.UserSchemaVersion,
		Users:             users,

		ProfileSchemaVersion: system.ProfileSchemaVersion,
		Profiles:             profiles,
	}
}

func ParseSystemFromConfigV1(c ConfigV1) model.System {
	return model.System{
		Restriction:                     "",
		InstanceId:                      c.InstanceId,
		SubspaceVersion:                 c.SubspaceVersion,
		SubspaceBuildNumber:             c.SubspaceBuildNumber,
		VpnServerVersion:                c.VpnServerVersion,
		VpnServerBuildNumber:            c.SubspaceBuildNumber,
		VpnServerAdministrationPassword: c.VpnServerAdministrationPassword,
		VpnServerAdministrationPort:     c.VpnServerAdministrationPort,
		VpnHubName:                      c.VpnHubName,
		Ip:                              c.Ip,
		//IpUpdatedDate: c.IpUpdatedDate,
		Host: c.VpnHost,
		//HostUpdatedDate: c.HostUpdatedDate,
		PreSharedKey: c.PreSharedKey,
		//PreSharedKeyUpdatedDate: c.PreSharedKeyUpdatedDate,
		Uuid: c.Uuid,
		//UuidUpdatedDate: c.UuidUpdatedDate,
		ConfigSchemaVersion:  c.ConfigSchemaVersion,
		UserSchemaVersion:    c.UserSchemaVersion,
		ProfileSchemaVersion: c.ProfileSchemaVersion,
		//UpdatedDate: c.UpdatedDate,
		//CreatedAt: c.CreatedAt,
	}
}

/**
	Convert User
 */
func ParseUserV1(u UserV1) model.User {
	user := model.User{
		Id:               u.Id,
		Email:            u.Email,
		Alias:            u.Alias,
		Role:             u.Role,
		EmailVerified:    u.EmailVerified,
		Enabled:          u.Enabled,
		PasswordHash:     u.PasswordHash,
		SetPasswordToken: u.SetPasswordToken,

		RevokedDate:   u.RevokedDate,
		LastLoginDate: u.LastLoginDate,
		UpdatedDate:   u.UpdatedDate,
		CreatedDate:   u.CreatedDate,
	}
	return user
}

func ToUserV1(u model.User) UserV1 {
	userV1 := UserV1{
		Id:               u.Id,
		Email:            u.Email,
		Alias:            u.Alias,
		Role:             u.Role,
		EmailVerified:    u.EmailVerified,
		Enabled:          u.Enabled,
		PasswordHash:     u.PasswordHash,
		SetPasswordToken: u.SetPasswordToken,

		RevokedDate:   u.RevokedDate,
		LastLoginDate: u.LastLoginDate,
		UpdatedDate:   u.UpdatedDate,
		CreatedDate:   u.CreatedDate,
	}
	return userV1
}

/**
	Convert profile
 */
func ToProfileV1(p model.Profile) ProfileV1 {
	profileV1 := ProfileV1{
		Id:             p.Id,
		Hub:            p.Hub,
		UserName:       p.UserName,
		FullName:       p.FullName,
		Description:    p.Description,
		UserId:         p.UserId,
		Enabled:        p.Enabled,
		LoginCount:     p.LoginCount,
		PasswordHash:   p.PasswordHash,
		NtLmSecureHash: p.NtLmSecureHash,
		VpnHost:        p.VpnHost,
		PreSharedKey:   p.PreSharedKey,

		//IncomingBytes: p.IncomingBytes,
		//OutgoingBytes: p.OutgoingBytes,

		RevokedDate:   p.RevokedDate,
		LastLoginDate: p.LastLoginDate,
		UpdatedDate:   p.UpdatedDate,
		CreatedDate:   p.CreatedDate,
	}
	return profileV1
}

func ParseProfileV1(p ProfileV1) model.Profile {
	profile := model.Profile{
		Id:             p.Id,
		Hub:            p.Hub,
		UserName:       p.UserName,
		FullName:       p.FullName,
		Description:    p.Description,
		UserId:         p.UserId,
		Enabled:        p.Enabled,
		LoginCount:     p.LoginCount,
		PasswordHash:   p.PasswordHash,
		NtLmSecureHash: p.NtLmSecureHash,
		VpnHost:        p.VpnHost,
		PreSharedKey:   p.PreSharedKey,

		//We do not keep transfer bytes now.
		IncomingBytes: 0,
		OutgoingBytes: 0,

		RevokedDate:   p.RevokedDate,
		LastLoginDate: p.LastLoginDate,
		UpdatedDate:   p.UpdatedDate,
		CreatedDate:   p.CreatedDate,
	}
	return profile
}
