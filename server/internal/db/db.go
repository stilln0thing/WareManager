package db

import (
   "os"

   "gorm.io/driver/postgres"
   "gorm.io/gorm"
   "github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() error {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB , err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the DB")
	} 
	return nil
}

func GetDB() *gorm.DB {
	return DB
}