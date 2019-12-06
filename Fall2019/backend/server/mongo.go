package main

import (
	"context"
	"github.com/aut-ce/Web101/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoUrl     = "mongodb://localhost:27017"
	DatabaseName = "kilid"
	// collections
	Houses = "houses"
	Mags   = "mags"
)

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(MongoUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetHouse(client *mongo.Client, id string) *models.House {
	cur, err := client.Database(DatabaseName).Collection(Houses).Find(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Fatal("Error on Finding the document", id, err)
	}
	for cur.Next(context.TODO()) {
		var house models.House
		err = cur.Decode(&house)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		return &house
	}
	return nil
}
