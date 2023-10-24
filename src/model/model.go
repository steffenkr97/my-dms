package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID      primitive.ObjectID
	Title   string `json:"title"`
	Article string `json:"article"`
}

type DocumentMeta struct {
	ID       primitive.ObjectID
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	FileName string   `json:"fileName"`
	Keywords []string `json:"keywords"`
	FullText string   `json:"fullText"`
}
