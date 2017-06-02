package administration

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/administration/configuration"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/utils"
)

func GenerateConfig(dbUri string) (string, error) {
	return generateConfigV1(dbUri)
}

func generateConfigV1(dbUri string) (string, error) {

	db, err := gorm.Open("mysql", dbUri)
	defer db.Close()
	if nil != err {
		return "", err
	}

	// Lock table write
	utils.LockTableWrite(db,
		model.User{}.TableName(),
		model.Profile{}.TableName(),
		model.System{}.TableName(),
	)
	defer utils.UnlockTable(db)

	// Get system data
	system, err := getSystem(db)
	if nil != err {
		return "", err
	}

	// Get user data
	users, err := getUsers(db)
	if nil != err {
		return "", err
	}

	// Get profile data
	profiles, err := getProfiles(db)
	if nil != err {
		return "", err
	}

	// Construct config data
	config := configuration.ToConfigV1(system, users, profiles)

	// Add checksum
	chkSum := config.CalculateCheckSum()
	config.CheckSum = chkSum

	if data, err := yaml.Marshal(&config); nil == err {
		return string(data), nil
	} else {
		return "", err
	}
}

func getSystem(db *gorm.DB) (model.System, error) {
	systemData := model.System{}

	sql := fmt.Sprintf("SELECT " +
					"`restriction`, " +
					"`subspace_version`, " +
					"`subspace_build_number`, " +
					"`vpn_server_version`, " +
					"`vpn_server_build_number`, " +
					"INET6_NTOA(ip) AS `ip`, " +
					"`ip_updated_date`, `host`, " +
					"`host_updated_date`, " +
					"`pre_shared_key`, " +
					"`pre_shared_key_updated_date`, " +
					"`uuid`, `uuid_updated_date`, " +
					"`smtp_host`, `smtp_port`, " +
					"`smtp_username`, " +
					"`smtp_password`, " +
					"`smtp_valid`, " +
					"`user_schema_version`, " +
					"`profile_schema_version`, " +
					"`config_schema_version`, " +
					"`updated_date`, " +
					"`created_at` " +
					"FROM %s", systemData.TableName())
	db.Raw(sql).Scan(&systemData)
	if nil != db.Error {
		return systemData, db.Error
	}
	return systemData, nil
}

func getUsers(db *gorm.DB) ([]configuration.UserV1, error) {
	var users []model.User
	db.Find(&users)
	if nil != db.Error {
		return nil, db.Error
	}

	var configUsers = make([]configuration.UserV1, 0)
	for _, user := range users {
		u := configuration.ToUserV1(user)
		configUsers = append(configUsers, u)
	}
	return configUsers, nil
}

func getProfiles(db *gorm.DB) ([]configuration.ProfileV1, error) {
	var profiles []model.Profile
	db.Find(&profiles)
	if nil != db.Error {
		return nil, db.Error
	}

	var configProfiles = make([]configuration.ProfileV1, 0)
	for _, profile := range profiles {
		p := configuration.ToProfileV1(profile)
		configProfiles = append(configProfiles, p)
	}
	return configProfiles, nil
}
