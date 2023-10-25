package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	var creds options.Credential

	creds.Username = "root"
	creds.Password = "rootpassword"

	Mongo_URL := "mongodb://127.0.0.1:27017"

	// client, err := mongo.Connect(
	// 	context.Background(),
	// 	options.Client().ApplyURI(Mongo_URL).SetAuth(creds),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL).SetAuth(creds))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB")
	return client
}
