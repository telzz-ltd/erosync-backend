package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func GetEnv[T any](key string, fallback T) (result T) {
	defer func() {
		if r := recover(); r != nil {
			result = fallback
		}
	}()

	return MustGetEnv[T](key)
}

func MustGetEnv[T any](key string) T {
	valStr := os.Getenv(key)
	if valStr == "" {
		panic(fmt.Sprintf("environment variable %q is required but not set", key))
	}

	var zero T
	var result any
	var err error

	// Inspect the underlying type using a zero value of T
	switch any(zero).(type) {
	case string:
		return any(valStr).(T)

	case int:
		result, err = strconv.Atoi(valStr)

	case bool:
		result, err = strconv.ParseBool(valStr)

	case time.Duration:
		result, err = time.ParseDuration(valStr)

	default:
		panic(fmt.Sprintf("unsupported environment variable type: %T", zero))
	}

	if err != nil {
		panic(fmt.Sprintf("failed to parse environment variable %q (%q): %v", key, valStr, err))
	}

	return result.(T)
}
