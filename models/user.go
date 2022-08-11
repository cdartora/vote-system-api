package models

type User struct {
	Id       int    `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}
