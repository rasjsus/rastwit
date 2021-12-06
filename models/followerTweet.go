package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FollowerTweet struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userid,omitempty"`
	UserRelationID string             `bson:"user_relation_id" json:"user_relation_id"`
	Tweet          struct {
		UserId  string    `bson:"_id" json:"_id,omitempty"`
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
	}
}
