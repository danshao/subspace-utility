package utils

import (
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration/configuration"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
)

func ToUserV1(u model.User) configuration.UserV1 {
	userV1 := configuration.UserV1{
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

func ToUser(u configuration.UserV1) model.User {
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

func ToProfileV1(p model.Profile) configuration.ProfileV1 {
	profileV1 := configuration.ProfileV1{
		Id:           p.Id,
		Hub:          p.Hub,
		UserName:     p.UserName,
		Description:  p.Description,
		UserId:       p.UserId,
		Enabled:      p.Enabled,
		LoginCount:   p.LoginCount,
		PasswordHash: p.PasswordHash,
		VpnHost:      p.VpnHost,
		PreSharedKey: p.PreSharedKey,

		//IncomingBytes: p.IncomingBytes,
		//OutgoingBytes: p.OutgoingBytes,

		RevokedDate:   p.RevokedDate,
		LastLoginDate: p.LastLoginDate,
		UpdatedDate:   p.UpdatedDate,
		CreatedDate:   p.CreatedDate,
	}
	return profileV1
}

func ToProfile(p configuration.ProfileV1) model.Profile {
	profile := model.Profile{
		Id:           p.Id,
		Hub:          p.Hub,
		UserName:     p.UserName,
		Description:  p.Description,
		UserId:       p.UserId,
		Enabled:      p.Enabled,
		LoginCount:   p.LoginCount,
		PasswordHash: p.PasswordHash,
		VpnHost:      p.VpnHost,
		PreSharedKey: p.PreSharedKey,

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