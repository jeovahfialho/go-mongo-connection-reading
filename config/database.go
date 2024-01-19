package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitDB() (*mongo.Client, error) {

	// Carregar as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Obter o usuário e a senha das variáveis de ambiente
	username := os.Getenv("MONGODB_USER")
	password := os.Getenv("MONGODB_PASSWORD")

	uri := "mongodb+srv://" + username + ":" + password + "@cluster-mongo-jeovahfia.ymkjshk.mongodb.net/?retryWrites=true&w=majority"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}
