package routes

import (
	"context"
	getcollection "my-dms/collection"
	database "my-dms/database"
	model "my-dms/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	DB := database.ConnectDB()
	postCollection := getcollection.GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var post model.Post

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"title": post.Title, "article": post.Article}

	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}
