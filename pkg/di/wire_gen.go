package di

import (
	"context"
	http "messagio/pkg/api"
	"messagio/pkg/api/handler"
	"messagio/pkg/config/kafka"
	"messagio/pkg/config/pgsql"
	"messagio/pkg/db"
	"messagio/pkg/kafka/producer"
	"messagio/pkg/repo/message"
	"messagio/pkg/service"
)

func InitializeAPI(pgCfg pgsql.Config, kafkaCfg kafka.Config) (*http.ServerHTTP, error) {
	bd, err := db.ConnectToBD(pgCfg)
	if err != nil {
		return nil, err
	}
	// Repository
	messageRepository := message.NewMessageDataBase(bd)

	produce, err := producer.NewProducer(kafkaCfg)
	if err != nil {
		return nil, err
	}

	//service - logic
	userService := service.NewService(messageRepository, produce)

	// Init Migrate
	err = userService.Migrate(context.Background())
	if err != nil {
		return nil, err
	}

	userHandler := handler.NewHandler(userService)
	serverHTTP := http.NewServerHTTP(userHandler)

	return serverHTTP, nil
}
