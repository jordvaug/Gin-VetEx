package handlers

import (
	"Go/GinVetEx/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb+srv://%s:%s@%s"
)

func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")

	connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)
	fmt.Println(connectionURI)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}

func Create(event *models.Event) (primitive.ObjectID, error) {
	fmt.Println("Create Event")
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	result, err := client.Database("GoTest").Collection("Events").InsertOne(ctx, event)
	if err != nil {
		log.Printf("Could not create Event: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func GetAllEvents() ([]*models.Event, error) {
	fmt.Print("Get All Events.")
	var events []*models.Event

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("GoTest")
	collection := db.Collection("Events")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Print("Found no results")
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &events)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return events, nil
}
