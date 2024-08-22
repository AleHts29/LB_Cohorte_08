package database

import (
	"c10_practica/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(databaseUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	return err
}
