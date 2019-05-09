package models

import (
	"github.com/globalsign/mgo/bson"
)

const (
	CollectionUsers = "users"
)

type User struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username    string        `json:"username" bson:"username" binding:"required,alpha"`
	DateOfBirth string        `json:"dateOfBirth" bson:"dateOfBirth" binding:"required"`
}
