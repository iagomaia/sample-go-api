package data

import (
	"context"

	data "github.com/iagomaia/sample-go-api/internal/data/protocols/message"
	models "github.com/iagomaia/sample-go-api/internal/domain/models/message"
	usecases "github.com/iagomaia/sample-go-api/internal/domain/usecases/message"
)

var (
	_ usecases.ICreateMessage = (*CreateMessage)(nil)
)

type CreateMessage struct {
	CreateMessageRepository data.ICreateMessage
	ctx                     context.Context
}

func (c *CreateMessage) WithCtx(ctx context.Context) usecases.ICreateMessage {
	return &CreateMessage{
		CreateMessageRepository: c.CreateMessageRepository,
		ctx:                     ctx,
	}
}

func (c *CreateMessage) Create(dto *usecases.CreateMessageDto) (*models.Message, error) {
	dataDto := mapDomainDtoToDataDto(dto)
	return c.CreateMessageRepository.WithCtx(c.ctx).Create(dataDto)
}

func mapDomainDtoToDataDto(dto *usecases.CreateMessageDto) *data.CreateMessageDto {
	return &data.CreateMessageDto{
		Text: dto.Text,
	}
}
