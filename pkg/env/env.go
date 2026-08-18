package env

import (
	"log"
	"os"
	"strconv"
)

func Get(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func GetOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("env %s not set\n", key)
	}
	return value
}

func GetInt(key string, fallback int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return fallback
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("[WARN] env %s is not a valid int \n", key)
		return fallback
	}
	return value
}
