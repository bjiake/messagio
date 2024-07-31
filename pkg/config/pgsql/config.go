package pgsql

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PsqlUser   string
	PsqlPass   string
	PsqlHost   string
	PsqlPort   string
	PsqlDBName string
}

func LoadConfig() (Config, error) {
	var config Config

	err := godotenv.Load("./app.env")
	if err != nil {
		return config, err
	}
	// Получение значений переменных окружения
	config.PsqlUser = os.Getenv("POSTGRES_USER")
	config.PsqlPass = os.Getenv("POSTGRES_PASSWORD")
	config.PsqlHost = os.Getenv("POSTGRES_HOST")
	config.PsqlDBName = os.Getenv("POSTGRES_DB")
	config.PsqlPort = os.Getenv("POSTGRES_PORT")

	return config, err
}
