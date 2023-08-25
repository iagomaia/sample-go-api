package factories

import (
	data "github.com/iagomaia/sample-go-api/internal/data/usecases/message"
	usecases "github.com/iagomaia/sample-go-api/internal/domain/usecases/message"
	repoFactories "github.com/iagomaia/sample-go-api/internal/factories/repositories/message"
)

var createMessageUseCase usecases.ICreateMessage

func GetCreateMessageUseCase() usecases.ICreateMessage {
	if createMessageUseCase != nil {
		return createMessageUseCase
	}

	createMessageUseCase = &data.CreateMessage{
		CreateMessageRepository: repoFactories.GetCreateMessageRepository(),
	}
	return createMessageUseCase
}
