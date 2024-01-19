package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InserirFornecedor(client *mongo.Client, fornecedor Fornecedor) error {
	collection := client.Database("quemindica").Collection("fornecedores")

	_, err := collection.InsertOne(context.TODO(), fornecedor)
	if err != nil {
		return err
	}

	return nil
}

func ListarFornecedores(client *mongo.Client) ([]Fornecedor, error) {
	collection := client.Database("quemindica").Collection("fornecedores")

	filter := bson.D{}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var fornecedores []Fornecedor
	for cur.Next(context.TODO()) {
		var fornecedor Fornecedor
		err := cur.Decode(&fornecedor)
		if err != nil {
			return nil, err
		}
		fornecedores = append(fornecedores, fornecedor)
	}

	return fornecedores, nil
}
