package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/omkz/golang-mongo-rest/config"
	"context"
	"log"
)

type Post struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	Content     string             `json:"content,omitempty"`
}

func PostAll() (posts []*Post, err error) {

	collection := config.DB.Database("blog").Collection("posts")

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Here's an array in which you can store the decoded documents
	var results []*Post

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		var post Post
		err = cur.Decode(&post)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		results = append(results, &post)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, nil
}