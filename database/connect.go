package database

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var (
	UserColl = DB.Database(os.Getenv("DB_NAME")).Collection("users")
)

var DB = ConnectDB()

func ConnectDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error in connecting with DB %s", err)
	}
	if err != nil {
		logrus.Error(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info("successfully connecting to database")

	return client
}
