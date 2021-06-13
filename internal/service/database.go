package service

import (
	"log"

	"github.com/jimmiepr/bank-transaction/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=root dbname=my-db port=15433 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Wallet{})
	return db
}
