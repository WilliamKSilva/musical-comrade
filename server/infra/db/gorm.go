package db

import (
	"github.com/WilliamKSilva/musical-comrade/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=docker password=admin dbname=musical_comrade_db port=5432"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.User{})

	return db, nil
}
