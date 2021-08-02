package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Type    string             `bson:"type, omitempty"`
	Message string             `bson:"message, omitempty"`
	Date    time.Time          `bson:"date, omitempty"`
}
