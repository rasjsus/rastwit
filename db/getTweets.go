package db

import (
	"context"
	"log"
	"time"

	"github.com/rasjsus/rastwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.Tweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	mdb := MongoCNN.Database("rastwitdb")
	col := mdb.Collection("tweets")

	var tweets []*models.Tweets

	condition := bson.M{
		"userid": ID,
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}

	for cursor.Next(context.TODO()) {
		var tweet models.Tweets
		err := cursor.Decode(&tweet)
		if err != nil {
			return tweets, false
		}
		tweets = append(tweets, &tweet)
	}

	return tweets, true
}
