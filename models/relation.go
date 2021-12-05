package models

type Relation struct {
	UserID         string `bson:"userid" json:"userid"`
	UserRelationID string `bson:"user_relation_id" json:"user_relation_id"`
}
