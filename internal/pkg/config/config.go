package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	HttpPort   string
	HttpHost   string
}

func LoadConfig() (Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
		panic(err)
	}

	DbPort, err := strconv.Atoi(os.Getenv("AB_DB_PORT"))
	if err != nil {
		return Config{}, fmt.Errorf("error loading db port: %w", err)
	}

	config := Config{
		DBHost:     os.Getenv("AB_DB_HOST"),
		DBPort:     DbPort,
		DBUser:     os.Getenv("AB_DB_USER"),
		DBPassword: os.Getenv("AB_DB_PASSWORD"),
		DBName:     os.Getenv("AB_DB_NAME"),
		HttpPort:   os.Getenv("AB_HTTP_PORT"),
		HttpHost:   os.Getenv("HTTP_HOST"),
	}

	log.Printf("config: %#v\n", config)
	return config, nil
}
