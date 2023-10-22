package main

import (
	routes "my-dms/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", routes.CreatePost)

	router.GET("getAll", routes.ReadAllPosts)

	// called as localhost:3000/getOne/{id}
	router.GET("getOne/:postId", routes.ReadOnePost)

	// called as localhost:3000/update/{id}
	router.PUT("/update/:postId", routes.UpdatePost)

	// called as localhost:3000/delete/{id}
	router.DELETE("/delete/:postId", routes.DeletePost)

	router.Run("localhost: 3000")
}
