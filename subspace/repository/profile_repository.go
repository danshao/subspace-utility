package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"

	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
)

type ProfileRepository interface {
	UpdateBatch(dataSet []*model.ProfileSnapshot) (err error)
}

type MysqlProfileRepository struct {
	Host string
	Account string
	Password string
	DatabaseName string
}

func InitProfileRepositoryWithHost(host string) (repo ProfileRepository) {
	return MysqlProfileRepository{
		Host: host,
		Account: config.MYSQL_DEFAULT_ACCOUNT,
		Password: config.MYSQL_DEFAULT_PASSWORD,
		DatabaseName: config.MYSQL_DEFAULT_DATABASE,
	}
}

func (repo MysqlProfileRepository) UpdateBatch(dataSet []*model.ProfileSnapshot) (err error) {
	uri := fmt.Sprintf(config.MYSQL_URI_FORMAT,
		repo.Account,
		repo.Password,
		repo.Host,
		repo.DatabaseName,
	)
	db, err := gorm.Open("mysql", uri)

	profile := model.Profile{}
	tx := db.Begin()
	for _, data := range dataSet {
		incomingBytes := data.IncomingBroadcastTotalSize + data.IncomingUnicastTotalSize
		outgoingBytes := data.OutgoingBroadcastTotalSize + data.OutgoingUnicastTotalSize
		tx.Model(&profile).Where("hub = ? AND username = ?", data.Hub, data.UserName).Updates(map[string]interface{}{
			"incoming_bytes": incomingBytes,
			"outgoing_bytes": outgoingBytes,
			"login_count": data.NumberOfLogins,
			"last_login_date": data.LastLogin,
		})
	}
	tx.Commit()
	err = db.Error

	defer db.Close()

	return err
}

func (repo MysqlProfileRepository) Update(row *model.ProfileSnapshot) (err error) {
	uri := fmt.Sprintf(config.MYSQL_URI_FORMAT,
		repo.Account,
		repo.Password,
		repo.Host,
		repo.DatabaseName,
	)
	db, err := gorm.Open("mysql", uri)

	if nil != err {
		return err
	}

	profile := model.Profile{}
	incomingBytes := row.IncomingBroadcastTotalSize + row.IncomingUnicastTotalSize
	outgoingBytes := row.OutgoingBroadcastTotalSize + row.OutgoingUnicastTotalSize
	db.Table(profile.TableName()).Where("hub = ? AND username = ?", row.Hub, row.UserName).Updates(map[string]interface{}{
		"incoming_bytes": incomingBytes,
		"outgoing_bytes": outgoingBytes,
		"login_count":    row.NumberOfLogins,
	})

	err = db.Error

	defer db.Close()

	return err
}