package helpers

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectToDB() *mongo.Client{
	err := godotenv.Load()

	if err != nil{
		fmt.Println("Error Loading the .env File")
	}

	Mongo_Url := os.Getenv("MONGODB_URL")
	clientOpt := options.Client().ApplyURI(Mongo_Url)
	client , _ = mongo.Connect(context.Background(), clientOpt)
	if client != nil{
		fmt.Println("Connected to DB")
	} else {
		fmt.Println("Error Connecting to DB")
	}
	return client
}