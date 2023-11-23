package database

import (
	"job-portal-api/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// err = db.Migrator().DropTable(&models.Company{}, &models.User{}, &models.Job{})
	// if err != nil {
	// 	return nil, err
	// }
	err = db.Migrator().AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Company{})
	if err != nil {
		return nil, err
	}
	err = db.Migrator().AutoMigrate(&models.Job{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
