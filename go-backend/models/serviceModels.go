package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	username  string             `json:"username" bson:"username,omitempty"`
	password  string             `json:"password" bson:"password,omitempty"`
}