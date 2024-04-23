package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() {
	userName := os.Getenv("USER") //testing123
	password := os.Getenv("PW")   //6ftCi9eH*9WaYakF6ZA$6y#!$
	// fmt.Println(userName, password)
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@speedrate.rz0ap.mongodb.net/speedrate?retryWrites=true&w=majority", userName, password))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Errorf("Could not connect to Mongo : %+v\n", err)
	}
	defer client.Disconnect(ctx)
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	speedrate := client.Database("speedrate")
	categoriesCollection := speedrate.Collection("category")
	// findOptions := options.Find()
	// findOptions.SetLimit(10)
	cur, err := categoriesCollection.Find(context.TODO(), bson.D{}, nil) //findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		fmt.Println(cur, "---------------------------------")
	}
	cur.Close(context.TODO())

}

func main() {
	connect()
}
