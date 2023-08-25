package controllers

import (
	"net/http"

	contracts "github.com/iagomaia/sample-go-api/internal/domain/contracts/message"
	usecases "github.com/iagomaia/sample-go-api/internal/domain/usecases/message"
	presentation "github.com/iagomaia/sample-go-api/internal/presentation/protocols"
)

type CreateMessageController struct {
	UseCase usecases.ICreateMessage
}

func (c *CreateMessageController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	reqBody := req.Body.(*contracts.CreateMessageRequest)

	dto := &usecases.CreateMessageDto{
		Text: reqBody.Text,
	}

	message, err := c.UseCase.WithCtx(req.Ctx).Create(dto)
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusCreated,
		Body: &contracts.CreateMessageResponse{
			MessageId: message.Id,
		},
	}
	return resp, nil
}
