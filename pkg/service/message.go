package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"messagio/pkg/domain/message"
	"messagio/pkg/domain/statistic"
	"time"
)

func (s *service) PostMessage(ctx context.Context, newMessage message.Message) (int64, error) {
	result, err := s.rMessage.Post(ctx, newMessage)
	if err != nil {
		return 0, err
	}

	err = s.produce.SendKafkaMessage(newMessage)
	if err != nil {
		log.Errorf("Failed to send message: %s", err)
		return 0, err
	}
	if time.Now().Second()%2 == 0 {
		newMessage.Status = "Success"
	} else {
		newMessage.Status = "Failure"
	}
	err = s.rMessage.PUT(ctx, result, newMessage)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (s *service) GetStaticsMessage(ctx context.Context) (statistic.Statistic, error) {
	messages, err := s.rMessage.GetAll(ctx)
	if err != nil {
		return statistic.Statistic{}, err
	}
	result := statistic.Statistic{}
	for _, msg := range messages {
		switch msg.Status {
		case "Success":
			result.Success++
		case "Failure":
			result.Failure++
		default:
			result.Unknown++
		}
	}
	return result, nil
}
