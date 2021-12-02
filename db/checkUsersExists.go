package db

import (
	"context"
	"time"

	"github.com/rasjsus/rastwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUsersExists(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mdb := MongoCNN.Database("rastwitdb")
	col := mdb.Collection("users")

	condition := bson.M{"email": email}
	var result models.Usuario

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
