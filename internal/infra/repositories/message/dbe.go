package repositories

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MessageCollection = "messages"
)

type MessageDbe struct {
	Id        *primitive.ObjectID `bson:"_id,omitempty"`
	Text      string              `bson:"text"`
	CreatedAt time.Time           `bson:"createdAt"`
	UpdatedAt *time.Time          `bson:"updatedAt"`
	DeletedAt *time.Time          `bson:"deletedAt"`
}
