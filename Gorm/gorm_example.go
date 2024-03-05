package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	Code  string
	Price uint
}

func main() {
	/**
	type Option interface {
		Apply(*Config) error
		AfterInitialize(*DB) error
	}
	*/
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "12345678", Price: 200})

	var storedProduct Product
	db.First(&storedProduct, 1) // find product with integer primary key
	db.First(&storedProduct, "code = ?", "123456")

	var products []Product
	db.Find(&products)
	for _, product := range products {
		log.Printf("Code %s, Price: %d", product.Code, product.Price)
	}
}
