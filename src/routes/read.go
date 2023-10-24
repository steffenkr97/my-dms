package routes

import (
	"context"
	// "encoding/json"
	// "fmt"
	getcollection "my-dms/collection"
	database "my-dms/database"
	model "my-dms/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func GetHomePage(c *gin.Context) {
// 	c.HTML(200, "../static/index.html", nil)
// }

func ReadAllPosts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	DB := database.ConnectDB()
	postCollection := getcollection.GetCollection(DB, "Posts")

	cursor, err := postCollection.Find(ctx, bson.D{{}})

	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var results []model.Post

	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	// for _, result := range results {
	// 	cursor.Decode(&result)
	// 	output, err := json.MarshalIndent(result, "", "    ")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%s\n", output)
	// }
	c.JSON(http.StatusAccepted, gin.H{"message": "success!", "Data": results})
}

func ReadOnePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	DB := database.ConnectDB()
	postCollection := getcollection.GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var result model.Post

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}
