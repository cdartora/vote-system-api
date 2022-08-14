package models

import uuid "github.com/satori/go.uuid"

type User struct {
	Id       uuid.UUID `json:"id" bson:"_id"`
	Name     string    `json:"name" bson:"name" validate:"required, min=2, max=100"`
	Password string    `json:"password" bson:"password" validate:"required, min=5, max=15"`
}

type RequestUser struct {
	Name     string `json:"name" bson:"name" validate:"required, min=2, max=100"`
	Password string `json:"password" bson:"password" validate:"required, min=5, max=15"`
}
