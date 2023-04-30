package bd

import (
	"context"
	"fmt"
	"redSocial/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{
			"$regex": `(?i)` + search,
		},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var find, include bool

	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = user.ID.Hex()
		include = false

		find, _ = GetRelation(r)
		if tipo == "new" && !find {
			include = true
		}
		if tipo == "follow" && find {
			include = true
		}
		if r.UserRelationID == ID {
			include = false
		}

		if include {
			user.Password = ""
			user.Biography = ""
			user.WebSite = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""

			results = append(results, &user)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
