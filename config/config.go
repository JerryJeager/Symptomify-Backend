package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Session *gorm.DB
var RedisClient *redis.Client


func GetSession() *gorm.DB {
	return Session
}

func ConnectToDB() {
	environment := os.Getenv("ENVIRONMENT")
	var db *gorm.DB
	var err error
	if environment == "development" {
		//local development DB config:::
		host := os.Getenv("HOST")
		username := os.Getenv("USER")
		password := os.Getenv("PASSWORD")
		port := os.Getenv("DBPORT")
		dbName := os.Getenv("DBNAME")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		//production DB config:::
		connectionString := os.Getenv("CONNECTION_STRING")
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	}

	if err != nil {
		log.Fatal(err)
	}

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Code{})
	// db.AutoMigrate(&models.Customer{})
	// db.AutoMigrate(&models.Photographer{})
	// db.AutoMigrate(&models.Token{})

	Session = db.Session(&gorm.Session{SkipDefaultTransaction: true, PrepareStmt: false})
	if Session != nil {
		log.Print("success: created db session")
	}
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Print(err)
		log.Print("failed to load envirionment variables")
		// log.Fatal("failed to load environment variables")
	}
}

func ConnectToRedis() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "development" {
		//local development redis config:::
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	} else {
		//production redis config:::
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
	}

	pong, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Print("failed to connect to redis instance")
		return
	}
	log.Printf("%s: redis instance connected", pong)
}
