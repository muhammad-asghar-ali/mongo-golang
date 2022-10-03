package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json: "id" bson: "_id"`
	Name   string        `json: "name" bson: "name"`
	Gander string        `json: "gander" bson: "gander"`
	Age    int           `json: "age" bson: "age"`
}
