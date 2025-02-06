package database

import (
	"fmt"
	"os"

	"github.com/charitan-go/profile-server/internal/donor/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("Connect to database success")
	}

	return err
}

func migrate() error {
	if err := DB.AutoMigrate(&model.Donor{}); err != nil {
		fmt.Println("Migrate failed")
		return err
	}

	return nil
}

func SetupDatabase() error {
	if err := connect(); err != nil {
		return err
	}

	if err := migrate(); err != nil {
		return err
	}

	return nil
}
