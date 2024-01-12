package config

// la unica funcion de este archivo es la conexion a la base de datos
// lo unico que tienen que hacer aqui es importar la variable "Database" a donde quieran usar la base de datos
import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func ConnectDB() {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		panic(err)
	}

	Database = db.Database("yconnect")
}
