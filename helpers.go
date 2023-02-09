package lungo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(name string) *mongo.Collection {
	collection := connection.Database(configuration.Name).Collection(name)

	return collection
}

func NewObjectID() ObjectID {
	return primitive.NewObjectID()
}
