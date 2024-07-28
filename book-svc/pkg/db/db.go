package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("POSTGRE_HOST")
	user := os.Getenv("POSTGRE_USER")
	password := os.Getenv("POSTGRE_PASSWORD")
	dbname := os.Getenv("POSTGRE_DB")
	port := os.Getenv("POSTGRE_PORT")
	sslmode := os.Getenv("POSTGRE_SSL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	postgresCon := postgres.Open(dsn)

	DB, err = gorm.Open(postgresCon, &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	DB.AutoMigrate(&Book{}, &Stock{})
}
