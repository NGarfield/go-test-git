package repository

import (
	"go-test-git/services"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertTestData(name, lastname string, age uint) error {
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
