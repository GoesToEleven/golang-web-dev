package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}
