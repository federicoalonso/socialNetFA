package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN variable de conexion a la base de datos
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://socialFede:socialFede555@cluster0.dxysw.mongodb.net/<dbname>?retryWrites=true&w=majority")

// ConectarBD - funcion de conexion a BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la BD")
	return client
}

// CheckConnection se fija si existe conexion
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
