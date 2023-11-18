package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	TgID int64              `bson:"tg_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}
