package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mapping struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Kesatuan  string             `json:"kesatuan" bson:"kesatuan"`
	Angkatan  string             `json:"angkatan" bson:"angkatan"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
