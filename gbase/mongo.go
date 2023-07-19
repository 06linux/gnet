package gbase

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type baseMongo struct{}

var Mongo = &baseMongo{}

func (baseMongo) Test() {
	fmt.Println("baseMongo Test")

	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to a mongodb server.
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	type Pet struct {
		Type string `bson:"type"`
		Name string `bson:"name"`
	}

	// Create a slice of documents to insert. We will lookup a subset of
	// these documents using regex.
	toInsert := []interface{}{
		Pet{Type: "cat", Name: "Mo"},
		Pet{Type: "dog", Name: "Loki"},
	}

	coll := client.Database("test").Collection("test")

	if _, err := coll.InsertMany(ctx, toInsert); err != nil {
		panic(err)
	}

	// Create a filter to find a document with key "name" and any value that
	// starts with letter "m". Use the "i" option to indicate
	// case-insensitivity.
	filter := bson.D{{"name", bson.D{{"$regex", "^m"}, {"$options", "i"}}}}

	_, err = coll.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
}
