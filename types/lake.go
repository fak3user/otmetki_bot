package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lake struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty"`
	UserId  primitive.ObjectID   `bson:"user_id,omitempty"`
	Title   string               `bson:"title,omitempty"`
	Records []primitive.ObjectID `bson:"records,omitempty"`
}
