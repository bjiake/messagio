package interfaces

import (
	"context"
	"messagio/pkg/domain/message"
)

type MessageRepository interface {
	Migrate(ctx context.Context) error
	Post(ctx context.Context, newPeople message.Message) (int64, error)
	PUT(ctx context.Context, id int64, updatedMessage message.Message) error
	GetAll(ctx context.Context) ([]message.Message, error)
}
