package repository

import (
	"go-test-git/services"

	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *DB{
	return &DB{
		db: db,
	}
}

func (db *DB) GetAllTest() ([]services.Test, error){
 	var test []services.Test
	result := db.db.Find(&test)
	if result.Error != nil {
		return nil, result.Error
	}
	return test,nil
}