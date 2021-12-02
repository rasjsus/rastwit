package db

import (
	"context"
	"time"

	"github.com/rasjsus/rastwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRegister(user models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	mdb := MongoCNN.Database("rastwitdb")
	col := mdb.Collection("users")

	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		register["last_name"] = user.LastName
	}

	register["birth_date"] = user.BirthDate

	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}

	if len(user.Ubication) > 0 {
		register["ubication"] = user.Ubication
	}

	if len(user.WebSite) > 0 {
		register["website"] = user.WebSite
	}

	updtString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, nil
	}

	return true, nil
}
