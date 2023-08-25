package dataDtosMocks

import (
	data "github.com/iagomaia/sample-go-api/internal/data/protocols/message"
)

func GetCreateMessageDataDtoMock() *data.CreateMessageDto {
	return &data.CreateMessageDto{
		Text: "some-text",
	}
}
