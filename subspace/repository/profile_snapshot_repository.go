package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"

	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/config"
)

type ProfileSnapshotRepository interface {
	InsertBatch(dataSet []*model.ProfileSnapshot) (err error)
}

type MysqlProfileSnapshotRepository struct {
	Host string
	Account string
	Password string
	DatabaseName string
}

func InitProfileSnapshotRepositoryWithHost(host string) (repo ProfileSnapshotRepository) {
	return MysqlProfileSnapshotRepository{
		Host: host,
		Account: config.MYSQL_DEFAULT_ACCOUNT,
		Password: config.MYSQL_DEFAULT_PASSWORD,
		DatabaseName: config.MYSQL_DEFAULT_DATABASE,
	}
}

func (repo MysqlProfileSnapshotRepository) InsertBatch(dataSet []*model.ProfileSnapshot) (err error) {
	uri := fmt.Sprintf(config.MYSQL_URI_FORMAT,
		repo.Account,
		repo.Password,
		repo.Host,
		repo.DatabaseName,
	)
	db, err := gorm.Open("mysql", uri)

	tx := db.Begin()
	for _, data := range dataSet {
		tx.Create(data)
	}
	tx.Commit()
	err = db.Error

	defer db.Close()

	return err
}

func (repo MysqlProfileSnapshotRepository) Insert(row *model.ProfileSnapshot) (err error) {
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

	db.NewRecord(&row)
	db.Save(&row)

	err = db.Error

	defer db.Close()

	return err
}