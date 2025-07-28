package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No env file")
	}
	log.Println(".env file loaded")
}

type DatabaseConfig struct {
	url string
}

func getInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	i, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}
	return i
}

func getString(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultValue
	}
	return val
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		url: getString("DATABASE_URL", ""),
	}
}
