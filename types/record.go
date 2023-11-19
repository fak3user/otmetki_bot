package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Rate      string              `bson:"rate,omitempty"`
	Note      string              `bson:"note,omitempty"`
	CreatedAt primitive.Timestamp `bson:"created_at,omitempty"`
}
