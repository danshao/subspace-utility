package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"gitlab.ecoworkinc.com/subspace/subspace-utility/subspace/model"

)

type ProfileRepository interface {
	InsertBatch(dataSet []*model.ProfileSnapshot) (err error)
}

type MysqlProfileRepository struct {
	Host string
	Account string
	Password string
	DatabaseName string

}

const MYSQL_URI_FORMAT = "%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local"

func (repo *MysqlProfileRepository) InsertBatch(dataSet []*model.ProfileSnapshot) (err error) {
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

func (repo *MysqlProfileRepository) Insert(row *model.ProfileSnapshot) (err error) {
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