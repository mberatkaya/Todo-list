package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nickname string             `json:"nickname"`
	FullName string             `json:"fullName"`
}
