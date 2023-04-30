package bd

import (
	"context"
	"redSocial/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ReadTweetFollowers(ID string, page int64) ([]models.TweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (page - 1) * 20

	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{"$match": bson.M{"userid": ID}})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	condition = append(condition, bson.M{"$unwind": "$tweet"})
	condition = append(condition, bson.M{"$sort": bson.M{"date": -1}})
	condition = append(condition, bson.M{"$skip": skip})
	condition = append(condition, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, condition)
	var result []models.TweetsFollowers
	err := cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
