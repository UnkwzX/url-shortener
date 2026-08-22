package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBURL      string
}

// getEnv читает и отдает из .env, или отдает defaultVal
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func Load() (*Config, error) {
	// Открывает .env и читает оттуда
	_ = godotenv.Load()

	//Берем из .env или читаем defaultVal если нету в .env
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "shortener")
	dbSSL := getEnv("DB_SSLMODE", "disable")

	//Собираю DSN - строку подключения | postgres://postgres(dbUser):postgres(dbPass)@localhost(dbHost):5432(dbPort)/shortener(dbName)?sslmode=disable(dbSSL)
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSL,
	)

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBURL:      dbURL,
	}, nil

}
