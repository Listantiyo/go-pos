package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("pos.db"), &gorm.Config{})
}