package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort             string
	DBDSN               string
	JWTSecret           string
	JWTRefreshSecret    string
	JWTExpireMinute     int
	JWTRefreshExpireDay int
}

func LoadConfig() AppConfig {
	_ = godotenv.Load()

	accessMinute, err := strconv.Atoi(getEnv("JWT_EXPIRE_MINUTE", "15"))
	if err != nil {
		accessMinute = 15
	}

	refreshDay, err := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRE_DAY", "7"))
	if err != nil {
		refreshDay = 7
	}

	return AppConfig{
		AppPort:             getEnv("APP_PORT", "3000"),
		DBDSN:               getEnv("DB_DSN", ""),
		JWTSecret:           getEnv("JWT_SECRET", "secret"),
		JWTRefreshSecret:    getEnv("JWT_REFRESH_SECRET", "refresh-secret"),
		JWTExpireMinute:     accessMinute,
		JWTRefreshExpireDay: refreshDay,
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
