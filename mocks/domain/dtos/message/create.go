package mocks

import (
	domain "github.com/iagomaia/sample-go-api/internal/domain/usecases/message"
)

func GetCreateMessageDomainDtoMock() *domain.CreateMessageDto {
	return &domain.CreateMessageDto{
		Text: "some-text",
	}
}
