package helper

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/pradumnasaraf/go-api/config"
	"github.com/pradumnasaraf/go-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {

	config.Config() // loading .env file
	mongoURL := os.Getenv("MONGO_URL")
	databaseName := os.Getenv("MONGO_DB")
	collectionName := os.Getenv("MONGO_COLLECTION")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	CheckNilErr(err)

	fmt.Println("MongoDB Connected sucessfully")

	Collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Collection instance is ready")

}

func InsertOneMovie(movie model.Netflix) {
	result, err := Collection.InsertOne(context.Background(), movie)

	CheckNilErr(err)
	fmt.Println("Inserted 1 movie in DB with id:", result.InsertedID)
}

func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := Collection.UpdateOne(context.Background(), filter, update)

	CheckNilErr(err)
	fmt.Println("Modified count:", result.UpsertedCount)
}

func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := Collection.DeleteOne(context.Background(), filter)

	CheckNilErr(err)
	fmt.Println("Modified count:", result.DeletedCount)
}

func DeleteAllMovies() int64 {
	filter := bson.D{{}}
	result, err := Collection.DeleteMany(context.Background(), filter)

	CheckNilErr(err)
	fmt.Println("Modified count:", result.DeletedCount)
	return result.DeletedCount
}

func GetOneMovie(movieId string) model.Netflix {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	var movie model.Netflix
	err := Collection.FindOne(context.Background(), filter).Decode(&movie)
	CheckNilErr(err)
	return movie
}

func GetAllMovies() []primitive.M {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	CheckNilErr(err)
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		CheckNilErr(err)

		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

func CheckNilErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
