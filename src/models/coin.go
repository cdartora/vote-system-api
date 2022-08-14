package models

import (
	uuid "github.com/satori/go.uuid"
)

type Vote struct {
	UserId uuid.UUID `json:"user_id" bson:"user_id"`
	Vote   int       `json:"vote" bson:"vote"`
}

type VoteBody struct {
	Vote int `json:"vote" bson:"vote"`
}

type CoinBody struct {
	Name  string `json:"name" bson:"name"`
	Code  string `json:"code" bson:"code"`
	Votes []Vote `json:"votes" bson:"votes"`
}

type CoinVotes struct {
	Name  string `json:"name" bson:"name"`
	Code  string `json:"code" bson:"code"`
	Votes int    `json:"votes" bson:"votes"`
}

type Coin struct {
	Id    uuid.UUID `json:"id,omitempty" bson:"_id"`
	Name  string    `json:"name" bson:"name"`
	Code  string    `json:"code" bson:"code"`
	Votes []Vote    `json:"votes" bson:"votes"`
}
