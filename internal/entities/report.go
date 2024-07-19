package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	Id          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"Title"`
	Content     string             `bson:"Content"`
	Responsible string             `bson:"Responsible"`
}
