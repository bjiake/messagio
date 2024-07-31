package service

import (
	"context"
	"messagio/pkg/db"
	produceI "messagio/pkg/kafka/producer/interface"
	messageI "messagio/pkg/repo/message/interface"
	interfaces "messagio/pkg/service/interface"
	"strconv"
)

type service struct {
	rMessage messageI.MessageRepository
	produce  produceI.ProducerUseCase
}

func NewService(
	messageRepository messageI.MessageRepository,
	produce produceI.ProducerUseCase,
) interfaces.ServiceUseCase {
	return &service{
		rMessage: messageRepository,
		produce:  produce,
	}
}

func (s *service) Migrate(ctx context.Context) error {
	if err := s.rMessage.Migrate(ctx); err != nil {
		return err
	}

	return nil
}

func (s *service) checkIdParam(id string) (int64, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil || idInt <= 0 {
		return 0, db.ErrParamNotFound
	}
	return idInt, nil
}
