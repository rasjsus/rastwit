package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rasjsus/rastwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LookPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	mdb := MongoCNN.Database("rastwitdb")
	col := mdb.Collection("users")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("User not found" + err.Error())
		return perfil, err
	}

	return perfil, nil
}
