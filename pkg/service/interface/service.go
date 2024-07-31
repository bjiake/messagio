package interfaces

import (
	"context"
	"messagio/pkg/domain/message"
	"messagio/pkg/domain/statistic"
)

type ServiceUseCase interface {
	Migrate(ctx context.Context) error
	PostMessage(ctx context.Context, newMessage message.Message) (int64, error)
	GetStaticsMessage(ctx context.Context) (statistic.Statistic, error)
}
