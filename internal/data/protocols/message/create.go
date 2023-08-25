package data

import (
	"context"

	models "github.com/iagomaia/sample-go-api/internal/domain/models/message"
)

type CreateMessageDto struct {
	Text string
}

type ICreateMessage interface {
	Create(dto *CreateMessageDto) (*models.Message, error)
	WithCtx(ctx context.Context) ICreateMessage
	Init()
}
