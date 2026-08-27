package config

import (
	"log"
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func GetInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		log.Println("unable to parse env value", key, "=", value)
		return fallback
	}

	return v
}
