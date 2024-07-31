package kafka

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	KafkaServerAddress string
	KafkaTopic         string
}

func LoadConfig() (Config, error) {
	var config Config
	err := godotenv.Load("./app.env")
	if err != nil {
		return config, err
	}
	config.KafkaServerAddress = os.Getenv("KAFKA_SERVER_ADDRESS")
	config.KafkaTopic = os.Getenv("KAFKA_TOPIC")
	return config, nil
}
