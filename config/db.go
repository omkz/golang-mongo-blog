package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB *mongo.Client

func ConnectDB(uri string) {
	var err error
	DB, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic(err)
	}

	err = DB.Connect(context.Background())
	if err != nil {
		log.Panic(err)
	}

}