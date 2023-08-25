package repositories

import (
	"time"

	data "github.com/iagomaia/sample-go-api/internal/data/protocols/message"
	models "github.com/iagomaia/sample-go-api/internal/domain/models/message"
)

func mapMessageDtoToDbe(dto *data.CreateMessageDto) *MessageDbe {
	return &MessageDbe{
		Id:        nil,
		Text:      dto.Text,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

func mapDbeToModel(dbe *MessageDbe) *models.Message {
	return &models.Message{
		Id:        dbe.Id.Hex(),
		Text:      dbe.Text,
		CreatedAt: dbe.CreatedAt,
		UpdatedAt: dbe.UpdatedAt,
		DeletedAt: dbe.DeletedAt,
	}
}
