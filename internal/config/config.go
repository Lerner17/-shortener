package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL       string `env:"BASE_URL,required"`
}

var instance *Config

var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("Load config...")
		instance = &Config{}
		if err := env.Parse(instance); err != nil {
			log.Fatalf("Cannot parse env variables: %v", err)
		}
	})
	log.Println("Successfully load config from env variables")
	return instance
}