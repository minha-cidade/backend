package config

import (
	"os"
	"strings"
)

type Config struct {
	Address        string
	DatabaseInfo   string
	AllowedOrigins []string
}

var config Config

func init() {
	// Lê o endereço de listen do servidor
	config.Address = os.Getenv("BACKEND_LISTEN_ADDRESS")
	if config.Address == "" {
		config.Address = ":8080"
	}

	// Lê o endereço do banco de dados Mongo
	config.DatabaseInfo = os.Getenv("BACKEND_MONGO_CONNECTION_STRING")
	if config.DatabaseInfo == "" {
		config.DatabaseInfo = "mongodb://localhost"
	}

	// Lê os endereços cors
	cors := os.Getenv("BACKEND_CORS_ALLOWED_ORIGINS")
	if cors == "" {
		config.AllowedOrigins = []string{"*"}
	} else {
		config.AllowedOrigins = strings.Split(cors, ",")
	}
}

func Get() Config {
	return config
}
