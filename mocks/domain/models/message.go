package mocks

import (
	"time"

	models "github.com/iagomaia/sample-go-api/internal/domain/models/message"
)

func GetMessageModelMock() *models.Message {
	return &models.Message{
		Id:        "some-id",
		Text:      "some-text",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}
