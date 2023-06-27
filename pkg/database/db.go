package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Db *mongo.Database
}

func NewDataBase(url string) (*DB, error) {

	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	fmt.Println("Created database and blog collection")
	db := client.Database("mydb")
	db.Collection("blog")

	return &DB{Db: db}, nil
}
