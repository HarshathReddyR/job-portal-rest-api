package database

import (
	"job-portal-api/config"
	"job-portal-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open(cfg config.Config) (*gorm.DB, error) {
	dsn := cfg.DatabaseCOnfig.DB_DSN
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmod=%s timezone=%s", cfg.DatabaseCOnfig.Host, cfg.DatabaseCOnfig.User,cfg.DatabaseCOnfig.Dbname, cfg.DatabaseCOnfig.Password, cfg.DatabaseCOnfig.Port, cfg.DatabaseCOnfig.Sslmode, cfg.DatabaseCOnfig.Timezone)
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
