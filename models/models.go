package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID   primitive.ObjectID `bson:"id,omitempty"`
	HASH string             `json:"hash,omitempty" validate:"required"`
	URL  string             `json:"url,omitempty" validate:"required"`
}
