package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"),
				os.Getenv("DB_NAME"),
)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil{
			panic("failed to connect database")
		}
		
		fmt.Println("database connected")
		return db
}