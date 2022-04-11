package quotes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Quote struct {
	Text   string `json:"quoteText"`
	Author string `json:"quoteAuthor"`
}

// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// // Connect to MongoDB
// client, err := mongo.Connect(context.TODO(), clientOptions)
// if err != nil {
// 	log.Fatal(err)
// }

// // Check the connection
// err = client.Ping(context.TODO(), nil)
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("Connected to MongoDB!")

func GetRandomQuote() Quote {
	var result Quote

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("icanhazquotes").Collection("quotes")

	return result
}

// client, err := mongo.NewClient(options.Client().ApplyURI("<ATLAS_URI_HERE>"))
// if err != nil {
// 	log.Fatal(err)
// }
// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// err = client.Connect(ctx)

// quickstartDatabase := client.Database("quickstart")
// podcastsCollection := quickstartDatabase.Collection("podcasts")
// episodesCollection := quickstartDatabase.Collection("episodes")
