package models

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
    dsn := "user=postgres password=belelik04 dbname=golang4 host=localhost port=5432 sslmode=disable TimeZone=UTC"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

	db.AutoMigrate(&User{}, &Product{})

    return db, nil
}