package db

import (
	"context"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	// err := godotenv.Load("./.env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
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
	Port                     string
	MongoDBConnectionString  string
	RedisConnectionString    string
	RedisConnectionPassword  string
	RedisTTL                 int
	LocationIQHost           string
	LocationIQAccessToken    string
	LocationIQResponseFormat string
}

// GetConfiguration method basically populate configuration information from .env and return Configuration model
func GetConfiguration() Configuration {

	// err := godotenv.Load("./.env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	ttl := os.Getenv("REDIS_TTL")
	intTTL, err := strconv.Atoi(ttl)
	if err != nil {
		panic(err)
	}

	configuration := Configuration{
		os.Getenv("PORT"),
		os.Getenv("MONGODB_CONNECTION_STRING"),
		os.Getenv("REDIS_CONNECTION_STRING"),
		os.Getenv("REDIS_CONNECTION_PASSWORD"),
		intTTL,
		os.Getenv("LOCATION_IQ_HOST"),
		os.Getenv("LOCATION_IQ_ACCESS_TOKEN"),
		os.Getenv("LOCATION_IQ_RESPONSE_FORMAT"),
	}
	return configuration
}
