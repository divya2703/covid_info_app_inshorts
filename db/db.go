package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := GetConfiguration()
	clientOptions := options.Client().ApplyURI(config.MongoDBConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("inshorts").Collection("covid_state_report_remote")
	log.Print("Connected to MongoDB! " + collection.Name())
	return collection
}

// Configuration model
type Configuration struct {
	Port                    string
	MongoDBConnectionString string
}

// GetConfiguration method basically populate configuration information from .env and return Configuration model
func GetConfiguration() Configuration {
	configuration := Configuration{
		os.Getenv("PORT"),
		os.Getenv("CONNECTION_STRING"),
	}
	return configuration
}
