package main

import (
	"fmt"
	routes "my-dms/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Load Home Page
	router.LoadHTMLGlob("static/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)

		fmt.Print("Ich wurde als Startseite aufgerufen!!!!!")
	})
	// router.POST("/", routes.CreatePost)

	router.GET("getAll", routes.ReadAllPosts)

	// called as localhost:3000/getOne/{id}
	router.GET("getOne/:postId", routes.ReadOnePost)

	// called as localhost:3000/update/{id}
	router.PUT("/update/:postId", routes.UpdatePost)

	// called as localhost:3000/delete/{id}
	router.DELETE("/delete/:postId", routes.DeletePost)

	router.POST("/document", routes.SaveDocument)

	router.Run("localhost: 3000")
}
