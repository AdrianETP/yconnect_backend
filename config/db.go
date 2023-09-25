package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var Database *mongo.Database
func ConnectDB()  {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		panic(err)
	}


	Database = db.Database("yconnect")
}

