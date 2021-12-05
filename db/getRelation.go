package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rasjsus/rastwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	mdb := MongoCNN.Database("rastwitdb")
	col := mdb.Collection("realtions")

	condition := bson.M{
		"userid":           t.UserID,
		"user_relation_id": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)

	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
