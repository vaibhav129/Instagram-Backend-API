package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Email string             `json:"email" bson:"email,omitempty"`
	Post   []Posts           `json:"post" bson:"post,omitempty"`
	Password string            `json:"password" bson:"password,omitempty"`

}

type Posts struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"Caption,omitempty" bson:"Caption,omitempty"`
	ImageUrl  string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	Timestamp string             `json:"timestamp,omitempty" bson:"timestamp,omite"`
}
