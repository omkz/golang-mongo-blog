package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = func() context.Context {
	return context.Background()
}()

type post struct {
	Title       string `bson:"name"`
	Description string `bson:"description"`
	Content     string `bson: "content`
}

func connect() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("blog"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("post").InsertOne(ctx, post{"the title", "the desciption", "the content"})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func main() {
	insert()
}
