package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quote struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Text   string             `bson:"quoteText,omitempty"`
	Author string             `bson:"quoteAuthor,omitempty"`
}
