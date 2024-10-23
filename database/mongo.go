package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Pradumnasaraf/Contributors/config"
	"github.com/Pradumnasaraf/Contributors/graph/model"
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

// creates a new MongoDB client and returns it
func NewMongoDB() *MongoDB {
	config.Config()

	Ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(Ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
	Collection = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION"))

	return &MongoDB{
		Client: client,
	}
}

// ADD a new contributor
func (db *MongoDB) Add(contributor *model.Contributor) error {
	_, err := Collection.InsertOne(Ctx, contributor)

	if err != nil {
		return errors.New("error while adding a new document. Document with the given ID may already exist")
	}

	return nil
}

// GET all contributors
func (db *MongoDB) GetAll() ([]*model.Contributor, error) {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, errors.New("error while getting the documents")
	}

	defer cursor.Close(Ctx)
	var result []*model.Contributor

	for cursor.Next(context.Background()) {
		var contributor *model.Contributor
		err := cursor.Decode(&contributor)
		if err != nil {
			return nil, errors.New("error while decoding the document")
		}

		result = append(result, contributor)
	}

	return result, nil
}

// GET a contributor by ID
func (db *MongoDB) GetByID(userId string) (*model.Contributor, error) {
	filter := bson.M{"_id": userId}
	var contributor *model.Contributor
	err := Collection.FindOne(Ctx, filter).Decode(&contributor)

	if err != nil {
		return nil, errors.New("error while getting the document. Document with the given ID may not exist")
	}

	return contributor, nil
}

// UPDATE a contributor by ID
func (db *MongoDB) UpdateByID(contributor *model.Contributor) error {
	filter := bson.M{"_id": contributor.UserID}
	update := bson.M{"$set": bson.M{"githubusername": contributor.GithubUsername, "name": contributor.Name, "email": contributor.Email}}
	result, _ := Collection.UpdateOne(context.Background(), filter, update)

	if result.MatchedCount == 0 {
		return errors.New("document not found. Document with the given ID may not exist")
	}

	return nil
}

// DELETE a contributor by ID
func (db *MongoDB) DeleteByID(userId string) error {
	filter := bson.M{"_id": userId}
	result, _ := Collection.DeleteOne(context.Background(), filter)

	if result.DeletedCount == 0 {
		return errors.New("error while deleting the document. Document with the given ID may not exist")
	}

	return nil

}

// DELETE contribution by ID
func (db *MongoDB) DeleteContributionByID(userId string, contributionID string) error {
	filter := bson.M{"_id": userId, "contributions.contributionid": contributionID}
	update := bson.M{"$pull": bson.M{"contributions": bson.M{"contributionid": contributionID}}}
	result, _ := Collection.UpdateOne(context.Background(), filter, update)

	if result.MatchedCount == 0 {
		return errors.New("document not found. Document with the given ID may not exist or contribution with the given ID may not exist")
	}

	return nil
}

// ADD contribution by ID
func (db *MongoDB) AddContributionByID(userId string, contribution *model.Contribution) error {
	filter := bson.M{"_id": userId, "contributions.contributionid": bson.M{"$ne": contribution.ContributionID}}
	update := bson.M{"$push": bson.M{"contributions": contribution}}
	result, _ := Collection.UpdateOne(Ctx, filter, update)

	if result.MatchedCount == 0 {
		return errors.New("could not add contribution. User with the given ID may not exist or contribution with the given ID may already exist")
	}

	return nil
}
