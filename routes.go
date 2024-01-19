package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, client *mongo.Client) {

	// Rota para inserir um fornecedor no banco de dados
	router.POST("/inserir-fornecedor", func(c *gin.Context) {
		var fornecedor Fornecedor
		if err := c.ShouldBindJSON(&fornecedor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := InserirFornecedor(client, fornecedor); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Fornecedor inserido com sucesso"})
	})

	router.GET("/get-data", func(c *gin.Context) {
		collection := client.Database("sample_airbnb").Collection("listingsAndReviews")
		filter := bson.D{{}}
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cur.Close(context.TODO())

		var results []interface{}
		for cur.Next(context.TODO()) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			results = append(results, result)
		}

		c.JSON(http.StatusOK, gin.H{"data": results})
	})

	router.GET("/get-item/:id", func(c *gin.Context) {
		id := c.Param("id")

		collection := client.Database("sample_airbnb").Collection("listingsAndReviews")
		filter := bson.M{"_id": id}
		var result bson.M
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	})

	router.GET("/search-by-name", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'name' not provided"})
			return
		}

		collection := client.Database("sample_airbnb").Collection("listingsAndReviews")
		filter := bson.M{"name": name}
		cur, err := collection.Find(context.TODO(), filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cur.Close(context.TODO())

		var results []interface{}
		for cur.Next(context.TODO()) {
			var result bson.M
			err := cur.Decode(&result)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			results = append(results, result)
		}

		c.JSON(http.StatusOK, gin.H{"data": results})
	})
}
