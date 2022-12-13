package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID  primitive.ObjectID `bson:"id,omitempty"`
	URL string             `json:"url,omitempty" validate:"required"`
}
