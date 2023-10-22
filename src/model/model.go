package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID      primitive.ObjectID
	Title   string `json:"title"`
	Article string `json:"article"`
}
