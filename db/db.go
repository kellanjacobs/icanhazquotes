package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kellanjacobs/icanhazquotes/config"
	"github.com/kellanjacobs/icanhazquotes/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo Client Used here so I can write the health check
var client *mongo.Client

// Mongo Database Connection
var db *mongo.Database

// Application Config
var c config.Config

// Database Context
var ctx context.Context

func init() {
	c, err := config.LoadConfig()
	credential := options.Credential{
		AuthSource: c.DB.Authsource,
		Username:   c.DB.User,
		Password:   c.DB.Password,
	}
	client, err = mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", c.DB.Host, c.DB.Port)).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(c.App.Port)*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(c.DB.DBNAME)
}

func CheckDB() bool {
	// Ping the database
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetRandomQuote() models.Quote {
	var result models.Quote
	collection := db.Collection("quotes")

	randomStage := bson.D{{"$sample", bson.D{{"size", 1}}}}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{randomStage})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Quote Author: %s, Quote: %s\n", result.Author, result.Text)
		}
	}
	return result
}
