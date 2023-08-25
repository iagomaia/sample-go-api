package repositories

import (
	"context"
	"log"
	"net/http"

	data "github.com/iagomaia/sample-go-api/internal/data/protocols/message"
	models "github.com/iagomaia/sample-go-api/internal/domain/models/message"
	"github.com/iagomaia/sample-go-api/internal/domain/models/utils"
	factories "github.com/iagomaia/sample-go-api/internal/factories/clients"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ data.ICreateMessage = (*CreateMessageRepository)(nil)
)

type CreateMessageRepository struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *CreateMessageRepository) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(MessageCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	r.session = session
	r.collection = collection
}

func (r *CreateMessageRepository) WithCtx(ctx context.Context) data.ICreateMessage {
	return &CreateMessageRepository{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *CreateMessageRepository) Create(dto *data.CreateMessageDto) (*models.Message, error) {
	defer r.session.EndSession(r.ctx)
	dbe := mapMessageDtoToDbe(dto)
	result, err := r.collection.InsertOne(r.ctx, dbe)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error inserting message into DB",
			OriginalError: err,
		}
		return nil, cErr
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	dbe.Id = &id

	return mapDbeToModel(dbe), nil
}
