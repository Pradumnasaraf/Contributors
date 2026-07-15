package graph

import "github.com/Pradumnasaraf/Contributors/mongo"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
}

var mongoClient *mongo.MongoDB

// GetMongoClient is a function that sets the MongoDB client.
func GetMongoClient(client *mongo.MongoDB) {
	mongoClient = client
}
