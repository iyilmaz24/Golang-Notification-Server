package config

import (
	"log"
	"os"
	"strings"
	"sync"
)

var (
	once *sync.Once
	instance *Config
)

type Config struct {
	Port string
	Cors map[string]bool
	AccessSecret string
}

func LoadConfig() *Config {

	once.Do(func() {

		port := ":8400"

		corsString := os.Getenv("CORS_ORIGIN_NOTIFICATION")
		if corsString == "" {
			log.Fatal("CORS_ORIGIN_NOTIFICATION env variable not set")
		}
		corsUrls := strings.Split(corsString, ",")

		corsOrigin := make(map[string]bool, len(corsUrls))
		for _, url := range corsUrls {
			corsOrigin[url] = true
		}

		adminSecret := os.Getenv("ADMIN_PASSWORD")
		if adminSecret == "" {
			log.Fatal("ADMIN_PASSWORD env variable not set")
		}

		instance = &Config {
			Port: port,
			Cors: corsOrigin,
			AccessSecret: adminSecret,
		}

	})

	return instance
}