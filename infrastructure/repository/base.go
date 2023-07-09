package repositories

import (
	"fmt"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type PGRepository struct {
	DB *gorm.DB
}

func NewGormDatabase() *PGRepository {
	db, _ := GormInit()
	return &PGRepository{DB: db}
}

func GormInit() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&entities.Contact{}, &entities.Message{}, &entities.Number{}, &entities.PhoneBook{}, &entities.Transaction{}, &entities.User{}, &entities.Wallet{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
