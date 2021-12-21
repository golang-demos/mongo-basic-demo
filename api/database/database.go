package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Db *mongo.Database

func ConnectDB() {
	dbConnectionURI := os.Getenv("MONGO_DB_CONN_URI")
	Client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbConnectionURI))
	if err != nil {
		panic(err)
	}
	Db = Client.Database(os.Getenv("MONGO_DB_NAME"))
}

func DisconnectDB() {
	if err := Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Database disconnected")
}
