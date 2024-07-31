package producer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"messagio/pkg/config/kafka"
	"messagio/pkg/domain/message"
	interfaces "messagio/pkg/kafka/producer/interface"
)

type producer struct {
	producer sarama.SyncProducer
	config   kafka.Config
}

func NewProducer(cfg kafka.Config) (interfaces.ProducerUseCase, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	produce, err := sarama.NewSyncProducer([]string{cfg.KafkaServerAddress}, config)
	log.Printf(cfg.KafkaServerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}

	return &producer{
		producer: produce,
		config:   cfg,
	}, nil
}

func (p *producer) SendKafkaMessage(currentMessage message.Message) error {
	messageJSON, err := json.Marshal(currentMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: p.config.KafkaTopic,
		Key:   sarama.StringEncoder(currentMessage.Name),
		Value: sarama.StringEncoder(messageJSON),
	}

	_, _, err = p.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	log.Printf("Success sending message to Kafka: %v (topic: %s, key: %s, message: %v)", err, p.config.KafkaTopic, currentMessage.Name, currentMessage)
	return err
}
