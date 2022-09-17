package dao

import (
	"testing"

	"github.com/surzia/crane-be/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var conn *gorm.DB

func TestInitDatabase(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Errorf("database init error: %v", err)
	}

	err = db.AutoMigrate(&models.Story{}, &models.Category{}, &models.Tag{})
	if err != nil {
		t.Errorf("database init error: %v", err)
	}
}

func TestConnectDatabase(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Errorf("cannot connect to database, error: %v", err)
	}

	conn = db
}
