package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/subspace/subspace-utility/subspace/model"

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

func (repo *MysqlProfileSnapshotRepository) InsertBatch(dataSet []*model.ProfileSnapshot) (err error) {
	uri := fmt.Sprintf(MYSQL_URI_FORMAT,
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

func (repo *MysqlProfileSnapshotRepository) Insert(row *model.ProfileSnapshot) (err error) {
	uri := fmt.Sprintf(MYSQL_URI_FORMAT,
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