package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type documentMetadata struct {
	fileID   int    `bson:"fileId"`
	filename string `bson:"fileName"`
	author   string `bson:"author"`
	filePath string `bson:"filePath"`
}

var docs = []documentMetadata{
	{fileID: 1, filename: "Rechnung-123", author: "Autor Name", filePath: "path/to/file1.pdf"},
	{fileID: 2, filename: "REchnung-999", author: "Autor Name", filePath: "path/to/file2.pdf"},
	{fileID: 3, filename: "Rechnung-123", author: "Autor Name", filePath: "path/to/file3.pdf"},
	{fileID: 4, filename: "Document-12345", author: "Autor Name", filePath: "path/to/file4.pdf"},
}

func getAllDocuments(c *gin.Context)
