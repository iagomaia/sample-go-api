package factories

import (
	data "github.com/iagomaia/sample-go-api/internal/data/protocols/message"
	messageRepository "github.com/iagomaia/sample-go-api/internal/infra/repositories/message"
)

var createMessageRepository data.ICreateMessage

func GetCreateMessageRepository() data.ICreateMessage {
	if createMessageRepository == nil {
		createMessageRepository = new(messageRepository.CreateMessageRepository)
		createMessageRepository.Init()
	}
	return createMessageRepository
}
