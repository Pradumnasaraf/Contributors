package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Pradumnasaraf/Contributors/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB struct holds the client and collection
type MongoDB struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// MongoInit creates a new MongoDB client and returns it
func MongoInit() *MongoDB {
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
	collection := client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COLLECTION"))

	return &MongoDB{
		Client:     client,
		Collection: collection,
	}
}

// Add a new contributor
func (db *MongoDB) Add(contributor *model.Contributor) error {
	Ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Collection.InsertOne(Ctx, contributor)

	if err != nil {
		return errors.New("error while adding a new document. Document with the given ID may already exist")
	}

	return nil
}

// GetAll contributors
func (db *MongoDB) GetAll() ([]*model.Contributor, error) {
	Ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.Collection.Find(Ctx, bson.D{{}})
	if err != nil {
		return nil, errors.New("error while getting the documents")
	}

	defer cursor.Close(Ctx)
	var result []*model.Contributor

	for cursor.Next(Ctx) {
		var contributor model.Contributor
		if err := cursor.Decode(&contributor); err != nil {
			return nil, errors.New("error while decoding the document")
		}
		result = append(result, &contributor)
	}

	return result, nil
}

// GetByID retrieves a contributor by ID
func (db *MongoDB) GetByID(userId string) (*model.Contributor, error) {
	Ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": userId}
	var contributor model.Contributor
	err := db.Collection.FindOne(Ctx, filter).Decode(&contributor)

	if err != nil {
		return nil, errors.New("error while getting the document. Document with the given ID may not exist")
	}

	return &contributor, nil
}

// UpdateByID updates a contributor by ID
func (db *MongoDB) UpdateByID(contributor *model.Contributor) error {
	Ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": contributor.UserID}
	update := bson.M{"$set": bson.M{
		"githubusername": contributor.GithubUsername,
		"name":           contributor.Name,
		"email":          contributor.Email,
	}}
	result, err := db.Collection.UpdateOne(Ctx, filter, update)

	if err != nil {
		return errors.New("error while updating the document")
	}
	if result.MatchedCount == 0 {
		return errors.New("document not found. Document with the given ID may not exist")
	}

	return nil
}

// DeleteByID deletes a contributor by ID
func (db *MongoDB) DeleteByID(userId string) error {
	Ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": userId}
	result, err := db.Collection.DeleteOne(Ctx, filter)

	if err != nil {
		return errors.New("error while deleting the document")
	}
	if result.DeletedCount == 0 {
		return errors.New("document not found. Document with the given ID may not exist")
	}

	return nil
}

// DeleteContributionByID removes a specific contribution from a contributor's record
func (db *MongoDB) DeleteContributionByID(userId string, contributionID string) error {
	Ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": userId, "contributions.contributionid": contributionID}
	update := bson.M{"$pull": bson.M{"contributions": bson.M{"contributionid": contributionID}}}
	result, err := db.Collection.UpdateOne(Ctx, filter, update)

	if err != nil {
		return errors.New("error while deleting the contribution")
	}
	if result.MatchedCount == 0 {
		return errors.New("document not found. Document with the given ID may not exist or contribution with the given ID may not exist")
	}

	return nil
}

// AddContributionByID adds a contribution by ID
func (db *MongoDB) AddContributionByID(userId string, contribution *model.Contribution) error {
	Ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"_id": userId, "contributions.contributionid": bson.M{"$ne": contribution.ContributionID}}
	update := bson.M{"$push": bson.M{"contributions": contribution}}
	result, err := db.Collection.UpdateOne(Ctx, filter, update)

	if err != nil {
		return errors.New("error while adding contribution")
	}

	if result.MatchedCount == 0 {
		return errors.New("could not add contribution. User with the given ID may not exist or contribution with the given ID may already exist")
	}

	return nil
}
