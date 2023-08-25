package infraCredentialsRepositoryMocks

import (
	"context"

	data "github.com/iagomaia/sample-go-api/internal/data/protocols/message"
	models "github.com/iagomaia/sample-go-api/internal/domain/models/message"
	mocks "github.com/iagomaia/sample-go-api/mocks/domain/models"
)

var (
	_ data.ICreateMessage = (*CreateMessageRepositoryMock)(nil)
)

type CreateMessageRepositoryMock struct {
	CreateMethodCalledTimes int
	CreateMethodReturn      *models.Message
	CreateMethodError       error
}

func (r *CreateMessageRepositoryMock) Init() {
	r.CreateMethodCalledTimes = 0
	r.CreateMethodError = nil
	r.CreateMethodReturn = nil
}

func (r *CreateMessageRepositoryMock) WithCtx(ctx context.Context) data.ICreateMessage {
	return r
}

func (r *CreateMessageRepositoryMock) Create(dto *data.CreateMessageDto) (*models.Message, error) {
	r.CreateMethodCalledTimes++
	if r.CreateMethodError != nil {
		return nil, r.CreateMethodError
	}
	if r.CreateMethodReturn != nil {
		return r.CreateMethodReturn, nil
	}
	return mocks.GetMessageModelMock(), nil
}
