package interfaces

import (
	"messagio/pkg/domain/message"
)

type ProducerUseCase interface {
	SendKafkaMessage(currentMessage message.Message) error
}
