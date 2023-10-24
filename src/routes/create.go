package routes

import (
	"context"
	"log"
	getcollection "my-dms/collection"
	database "my-dms/database"
	model "my-dms/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveDocument(c *gin.Context) {
	DB := database.ConnectDB()
	postCollection := getcollection.GetCollection(DB, "myDocuments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.DocumentMeta)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	documentPayload := model.DocumentMeta{
		ID:       primitive.NewObjectID(),
		Title:    post.Title,
		Author:   post.Author,
		FullText: post.FullText,
		// PdfPath:  post.PdfPath,
	}

	result, err := postCollection.InsertOne(ctx, documentPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}},
	)
}

func CreatePost(c *gin.Context) {
	DB := database.ConnectDB()
	postCollection := getcollection.GetCollection(DB, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.Post)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	postPayload := model.Post{
		ID:      primitive.NewObjectID(),
		Title:   post.Title,
		Article: post.Article,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}},
	)
}
