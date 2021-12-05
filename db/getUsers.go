package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rasjsus/rastwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(ID string, page int64, search string, typeUser string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	mdb := MongoCNN.Database("rastwit")
	col := mdb.Collection("users")

	var result []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	var found, include bool
	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return result, false
		}
		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = GetRelation(r)
		if typeUser == "new" && found == false {
			include = true
		}
		if typeUser == "follow" && found == true {
			include = true
		}
		if r.UserRelationID == ID {
			include = false
		}
		if include == true {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Ubication = ""
			s.Banner = ""
			s.Email = ""

			result = append(result, &s)
		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}
	cur.Close(ctx)
	return result, true
}
