package server

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
