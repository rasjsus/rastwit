package models

type TweetResponse struct {
	Message string `bson:"message" json:"message"`
}
