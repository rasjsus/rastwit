package db

import (
	"context"
	"time"

	"github.com/rasjsus/rastwit/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(user *models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mdb := MongoCNN.Database("rastwitdb")
	col := mdb.Collection("users")

	user.Password, _ = PasswordCrypt(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
