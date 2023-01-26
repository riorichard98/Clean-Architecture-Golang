package config

import (
	// "fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// db config
	DBHost      string `json:"host"`
	DBUser      string `json:"user"`
	DBPassword  string `json:"password"`
	DBName  	string `json:"databaseName"`
	DBPort      string `json:"port"`
	DBSsl		string `json:"ssl"`
	DBMongo_Uri string `json:"mongo_uri"`
	// Schema      string `json:"schema"`
	// Timeout     int    `json:"timeout"`
	// MaxPoolSize int    `json:"maxPool"`
	// MinPoolSize int    `json:"minPoolSize"`
	// DebugMode   bool   `json:"debugMode"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loading .env file")
	}
	return &Config{
		DBHost: os.Getenv("HOST"),
		DBUser: os.Getenv("USER"),
		DBPassword: os.Getenv("PASSWORD"),
		DBName: os.Getenv("DATABASE"),
		DBPort: os.Getenv("PORT"),
		DBMongo_Uri: os.Getenv("MONGO_URI"),
	}
}
