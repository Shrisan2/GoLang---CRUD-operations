package main

import (
	"learnWednesday/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectMySQL() (*gorm.DB, error) {
	// Setting up a dsn
	// Structure : admin:password@tcp(host:port)/schemaName?...
	dsn := "root:root@tcp(127.0.0.1:3306)/book_library?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Checking if there is an error while connecting to the db
	if err != nil {
		panic("Failed to connect to the database!!")
	}

	// Auto Migrating DataBase - these create tables in the db schema if not exist
	db.AutoMigrate(&models.Book{})

	return db, nil
}
