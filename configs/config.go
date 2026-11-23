package configs

import (
	"os"
	"strconv"
)

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val;
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	val := os.Getenv(key);
	if val == ""{
		return  fallback
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return  fallback
	}
	return  valInt;
}