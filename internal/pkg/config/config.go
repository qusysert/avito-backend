package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
}

func LoadConfig() (Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
		panic(err)
	}

	config := Config{}

	log.Printf("config: %#v\n", config)
	return config, nil
}
