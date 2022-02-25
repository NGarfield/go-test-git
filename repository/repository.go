package repository

import (
	"go-test-git/services"

	"github.com/jinzhu/gorm"
)

type DB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}

func (db *DB) GetAllTest() ([]services.Test, error) {
	var test []services.Test
	result := db.db.Find(&test)
	if result.Error != nil {
		return nil, result.Error
	}
	return test, nil
}

func (r *DB) InsertTestData(name, lastname string, age uint) error {
	test := services.Test{
		Name:     name,
		Lastname: lastname,
		Age:      age,
	}
	tx := r.db.Create(&test)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
