package routes

import (
	"context"
	"fmt"
	"io"
	"log"
	getcollection "my-dms/collection"
	database "my-dms/database"
	model "my-dms/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveDocument(c *gin.Context) {
	fmt.Print("Ich wurde aufgerufen!!")

	// PDF-Datei aus Formular erhalten
	file, err := c.FormFile("pdf_file")
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	title := c.PostForm("title")
	author := c.PostForm("author")
	description := c.PostForm("description")
	// keywords := c.PostForm("keywords")
	// fileName := header.Filename
	fileName := file.Filename

	// ab hier DB und speicherung

	// DB Connection herstellen
	dbClient := database.ConnectDB()
	fmt.Print("db connected")
	// Collection beziehen
	collection := dbClient.Database("dms").Collection("docs")
	fmt.Print("collection izz da")
	// gridFs Bucket erstellen etc.
	fs, err := gridfs.NewBucket(
		dbClient.Database("dms"),
		options.GridFSBucket().SetName("documentsBucket"),
	)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Speichere die PDF-Datei in GridFS
	uploadStream, err := fs.OpenUploadStream(file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer uploadStream.Close()

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	_, err = io.Copy(uploadStream, src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	meta := model.DocumentMeta{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Author:      author,
		FileName:    fileName,
		Description: description,
		Keywords:    []string{"k1", "k2"},
		FullText:    "myFullText still TBD",
	}

	// in DB speichern
	_, err = collection.InsertOne(context.Background(), meta)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File and metadata uploaded successfully"})

	// DB := database.ConnectDB()
	// postCollection := getcollection.GetCollection(DB, "myDocuments")
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// post := new(model.DocumentMeta)
	// defer cancel()
	//
	// if err := c.BindJSON(&post); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": err})
	// 	log.Fatal(err)
	// 	return
	// }
	//
	// documentPayload := model.DocumentMeta{
	// 	ID:       primitive.NewObjectID(),
	// 	Title:    post.Title,
	// 	Author:   post.Author,
	// 	FullText: post.FullText,
	// 	// PdfPath:  post.PdfPath,
	// }
	//
	// result, err := postCollection.InsertOne(ctx, documentPayload)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	// 	return
	// }
	//
	// c.JSON(
	// 	http.StatusCreated,
	// 	gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}},
	// )
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
