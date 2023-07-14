package repositories

import (
	"fmt"
	"github.com/cyneptic/letsgo-smspanel/internal/core/entities"
	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"sync"
)

type PGRepository struct {
	DB *gorm.DB
}
type RedisDB struct {
	Client *redis.Client
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
	err = db.AutoMigrate(&entities.Template{}, &entities.BlacklistWord{}, &entities.BlacklistRegex{}, &entities.Prices{}, &entities.Contact{}, &entities.Message{}, &entities.Number{}, &entities.PhoneBook{}, &entities.Transaction{}, &entities.User{}, &entities.Wallet{})
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

func createRedisConnection() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return redisClient

}

func RedisInit() *RedisDB {
	o := sync.Once{}
	var db *RedisDB
	o.Do(func() {
		db = &RedisDB{
			Client: createRedisConnection(),
		}
	})
	return db
}
