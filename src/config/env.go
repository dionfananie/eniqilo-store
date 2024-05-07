package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var isEnvLoaded = false

var (
	DB_NAME     = getEnv("DB_NAME", "")
	DB_PORT     = getEnv("DB_PORT", "")
	DB_HOST     = getEnv("DB_HOST", "")
	DB_USERNAME = getEnv("DB_USERNAME", "")
	DB_PASSWORD = getEnv("DB_PASSWORD", "")
	DB_PARAMS   = getEnv("DB_PARAMS", "sslmode=disable")

	JWT_SECRET  = getEnv("JWT_SECRET", "")
	JWT_EXP     = getEnv("JWT_EXP", "8h")
	BCRYPT_SALT = getEnvAsInt("BCRYPT_SALT", 8)
)

func loadEnv() {
	if !isEnvLoaded {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		isEnvLoaded = true
	}
}

func getEnv(name string, fallback string) string {
	loadEnv()
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}

func getEnvAsInt(name string, fallback int) int {
	loadEnv()
	if value, exists := os.LookupEnv(name); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error converting %v to int", name)
		}
		return intValue
	}

	return fallback
}
