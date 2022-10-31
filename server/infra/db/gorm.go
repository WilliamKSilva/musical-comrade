package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=docker password=admin dbname=musical_comrade port=5542"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	return db, nil
}
