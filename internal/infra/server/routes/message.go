package routes

import (
	"github.com/go-chi/chi/v5"
	contracts "github.com/iagomaia/sample-go-api/internal/domain/contracts/message"
	factories "github.com/iagomaia/sample-go-api/internal/factories/controllers/message"
	"github.com/iagomaia/sample-go-api/internal/infra/adapters"
)

func GetMessageRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", adapters.AdaptRoute[contracts.CreateMessageRequest](factories.GetCreateMessageController(), nil))
	return r
}
