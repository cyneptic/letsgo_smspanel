package repositories

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGRepository struct {
	DB *gorm.DB
}

func NewGormDatabase() *PGRepository {
	db, _ := GormInit()
	return &PGRepository{DB: db}
}

func GormInit() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&entities.Prices{}, &entities.Contact{}, &entities.Message{}, &entities.Number{}, &entities.PhoneBook{}, &entities.Transaction{}, &entities.User{}, &entities.Wallet{})
	if err != nil {
		return nil, err
	}
	var pricesCount int64
	db.Model(&entities.Prices{}).Count(&pricesCount)

	if pricesCount == 0 {
		single, err := strconv.Atoi(os.Getenv("DEFAULT_PRICE_SINGLE"))
		if err != nil {
			return nil, err
		}
		group, err := strconv.Atoi(os.Getenv("DEFAULT_PRICE_GROUP"))
		if err != nil {
			return nil, err
		}
		db.Create(&entities.Prices{
			SingleMessage: single,
			GroupMessage:  group,
		})
	}

	return db, nil
}
