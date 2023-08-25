package factories

import (
	factories "github.com/iagomaia/sample-go-api/internal/factories/usecases/message"
	controllers "github.com/iagomaia/sample-go-api/internal/presentation/controllers/message"
	presentation "github.com/iagomaia/sample-go-api/internal/presentation/protocols"
)

var createMessageController presentation.IHandler

func GetCreateMessageController() presentation.IHandler {
	if createMessageController != nil {
		return createMessageController
	}

	createMessageController = &controllers.CreateMessageController{UseCase: factories.GetCreateMessageUseCase()}
	return createMessageController
}
