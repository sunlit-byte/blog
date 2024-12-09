package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost     string // host
	DBPort     string // port
	DBUser     string // user
	DBPassword string // password
	DBName     string // name
	LogLevel   string // log level
	LogUrl     string // log url
)

func LoadConfig(envFile *string) error {
	// 加载env文件
	if envFile == nil {
		return fmt.Errorf("ENV file is nil")
	}
	err := godotenv.Load(*envFile)
	if err != nil {
		return err
	}
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	LogLevel = os.Getenv("LOG_LEVEL")
	LogUrl = os.Getenv("LOG_URL")
	return nil

}
