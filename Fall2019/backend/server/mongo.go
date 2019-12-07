package main

import (
	"context"
	"github.com/aut-ce/Web101/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	MongoUrl     = "mongodb://82.102.10.78:27017"
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
		house.CreatedAt = house.Ca
		house.Ca = 0
		return &house
	}
	return nil
}

func GetOccasion(client *mongo.Client, limit int, skip int) *models.Occasion {
	var occasion models.Occasion

	opts := options.Find().
		SetSort(bson.D{{"created_at", -1}}).
		SetSkip(int64(skip)).
		SetLimit(int64(limit))
	cur, err := client.Database(DatabaseName).Collection(Houses).Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		log.Fatal("Error on Finding the document ", err)
	}

	for cur.Next(context.TODO()) {
		var house models.OccasionHouse
		err = cur.Decode(&house)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		house.CreatedAt = house.Ca
		house.Ca = 0

		occasion.Items = append(occasion.Items, house)
	}
	return &occasion
}

func GetAllMagazine(client *mongo.Client) *models.MagazineResponse {
	var mags models.MagazineResponse

	cur, err := client.Database(DatabaseName).Collection(Mags).Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error on Finding the documents", err)
	}
	for cur.Next(context.TODO()) {
		var mag models.Magazine
		err = cur.Decode(&mag)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		mags.Items = append(mags.Items, mag)
	}
	return &mags
}

func GetAllChartHouse(client *mongo.Client) *models.Chart {
	var chart models.Chart

	cur, err := client.Database(DatabaseName).Collection(Houses).Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error on Finding the document ", err)
	}

	for cur.Next(context.TODO()) {
		var house models.ChartHouse
		err = cur.Decode(&house)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}

		chart.Data = append(chart.Data, house)
	}

	return &chart
}
