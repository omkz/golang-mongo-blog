package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/omkz/golang-mongo-rest/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Content     string             `json:"content"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

func PostAll() (posts []*Post, err error) {

	collection := config.DB.Database("blog").Collection("posts")

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}})
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

func PostCreate(post *Post) error {

	collection := config.DB.Database("blog").Collection("posts")

	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	return nil
}

func PostFindById(id string) (*Post, error) {
	post := new(Post)
	filter := bson.D{{"id", id}}
	err := config.DB.Database("blog").Collection("posts").FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		log.Println("Error!", err.Error())
	}

	return post, nil
}
