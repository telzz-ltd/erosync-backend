package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port int

	AppName string
	AppUrl  string

	DatabaseUrl string

	MailHost     string
	MailPort     int
	MailUsername string
	MailPassword string
	MailFrom     string
}

func New() *Config {
	return &Config{
		Port: GetEnv("PORT", 8080),

		AppName: GetEnv("APP_NAME", "Erosync"),
		AppUrl:  MustGetEnv[string]("APP_URL"),

		DatabaseUrl: MustGetEnv[string]("DATABASE_URL"),

		MailHost:     MustGetEnv[string]("MAIL_HOST"),
		MailPort:     MustGetEnv[int]("MAIL_PORT"),
		MailUsername: GetEnv("MAIL_USERNAME", ""),
		MailPassword: GetEnv("MAIL_PASSWORD", ""),
		MailFrom:     GetEnv("MAIL_FROM", "Erosync Support <support@erosyncng.com>"),
	}
}

func GetEnv[T any](key string, fallback T) (value T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			value = fallback
		}
	}()

	value = MustGetEnv[T](key)

	return value
}

func MustGetEnv[T any](key string) T {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		log.Panicln(key, "variable not set")
	}

	var zero T

	switch any(zero).(type) {
	case int:
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Panicf("Error converting %s variable to int %s\n", key, err)
		}
		return any(v).(T)
	case string:
		return any(value).(T)
	default:
		log.Panicf("%T type not supported\n", zero)
		return zero
	}
}
