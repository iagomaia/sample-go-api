package data_test

import (
	"context"
	"net/http"
	"testing"

	data "github.com/iagomaia/sample-go-api/internal/data/usecases/message"
	utils "github.com/iagomaia/sample-go-api/internal/domain/models/utils"
	domainDtosMocks "github.com/iagomaia/sample-go-api/mocks/domain/dtos/message"
	repositoriesMocks "github.com/iagomaia/sample-go-api/mocks/infra/repositories/message"
)

type CreateMessageSutTypes struct {
	UseCase                     *data.CreateMessage
	CreateMessageRepositoryMock *repositoriesMocks.CreateMessageRepositoryMock
}

func GetCreateMessageSutDependencies() *CreateMessageSutTypes {
	createMessageRepositoryMock := &repositoriesMocks.CreateMessageRepositoryMock{}
	createMessageRepositoryMock.Init()

	useCase := &data.CreateMessage{
		CreateMessageRepository: createMessageRepositoryMock,
	}

	return &CreateMessageSutTypes{
		UseCase:                     useCase,
		CreateMessageRepositoryMock: createMessageRepositoryMock,
	}
}

func Test_Create(t *testing.T) {
	t.Run("should create a message calling the repository", func(t *testing.T) {
		sut := GetCreateMessageSutDependencies()
		_, err := sut.UseCase.WithCtx(
			context.Background(),
		).Create(domainDtosMocks.GetCreateMessageDomainDtoMock())

		if err != nil {
			t.Errorf("failed: %v", err)
		}
		if sut.CreateMessageRepositoryMock.CreateMethodCalledTimes != 1 {
			t.Error("Create message repository not called")
		}
	})
	t.Run("should return error if failed to create message", func(t *testing.T) {
		sut := GetCreateMessageSutDependencies()
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Internal Server Error",
			OriginalError: nil,
		}
		sut.CreateMessageRepositoryMock.CreateMethodError = cErr

		message, err := sut.UseCase.WithCtx(
			context.Background(),
		).Create(domainDtosMocks.GetCreateMessageDomainDtoMock())

		if err == nil {
			t.Errorf("expected error: %v", cErr)
		}
		if message != nil {
			t.Error("should return nil message if an error happens")
		}
		if sut.CreateMessageRepositoryMock.CreateMethodCalledTimes != 1 {
			t.Error("create message repository not called")
		}
	})
}
