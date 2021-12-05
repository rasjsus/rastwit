package db

import (
	"context"
	"time"

	"github.com/rasjsus/rastwit/models"
)

func DeleteRelation(t *models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	mdb := MongoCNN.Database("rastwitdb")
	col := mdb.Collection("relations")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
