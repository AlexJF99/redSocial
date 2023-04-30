package bd

import (
	"context"
	"log"
	"redSocial/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweet(ID string, page int64) ([]*models.ObtainedTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var results []*models.ObtainedTweets

	condition := bson.M{
		"userid": ID,
	}

	option := options.Find()
	option.SetLimit(20)
	option.SetSort(bson.D{{Key: "Date", Value: -1}})
	option.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, option)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var regist models.ObtainedTweets
		err := cursor.Decode(&regist)
		if err != nil {
			return results, false
		}
		results = append(results, &regist)
	}
	return results, true
}
