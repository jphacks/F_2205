package config

import (
	"fmt"
	"os"
)

// DSNはdataSourceNameを返します、もし必須の環境変数が設定されてなかった場合はerrorを返します
func DSN() (string, error) {
	if os.Getenv("ENVIROMENT") == "PRODUCTION" {
		return os.Getenv("POSTGRES_URL"), nil
	}

	// DBの環境変数
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return "", fmt.Errorf("ERROR : required environment variable not found")
	}
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	), nil
}
