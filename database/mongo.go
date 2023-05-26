package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pradumnasaraf/go-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var Ctx context.Context

// MongoDB is a struct that holds the MongoDB client
type MongoDB struct {
	Client *mongo.Client
}

// NewMongoDB creates a new MongoDB client and returns it
func NewMongoDB() *MongoDB {
	// Create a context
	Ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Collection = client.Database("opensource").Collection("contributors")

	return &MongoDB{
		Client: client,
	}
}

func (db *MongoDB) Add(contributor *model.Contributor) {
	result, err := Collection.InsertOne(Ctx, contributor)

	CheckNilErr(err)
	fmt.Println("Inserted a single document: ", result.InsertedID)
}

func (db *MongoDB) GetAll() []*model.Contributor {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	CheckNilErr(err)
	defer cursor.Close(Ctx)
	var result []*model.Contributor

	for cursor.Next(context.Background()) {
		var contributor *model.Contributor
		err := cursor.Decode(&contributor)
		CheckNilErr(err)

		result = append(result, contributor)
	}

	return result
}

func (db *MongoDB) GetByID(id string) *model.Contributor {
	filter := bson.M{"_id": id}
	var contributor *model.Contributor
	err := Collection.FindOne(Ctx, filter).Decode(&contributor)
	CheckNilErr(err)
	return contributor
}

func (db *MongoDB) UpdateByID(contributor *model.Contributor) {
	filter := bson.M{"_id": contributor.ID}
	update := bson.M{"$set": bson.M{"githubUsername": contributor.GithubUsername, "name": contributor.Name, "email": contributor.Email}}
	result, err := Collection.UpdateOne(context.Background(), filter, update)

	CheckNilErr(err)
	fmt.Println("Updated the document: ", result.UpsertedID)
}

func (db *MongoDB) DeleteByID(id string) {
	filter := bson.M{"_id": id}
	result, err := Collection.DeleteOne(context.Background(), filter)

	CheckNilErr(err)
	fmt.Println("Deleted the document: ", result.DeletedCount)
}

func CheckNilErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
