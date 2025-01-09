package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("testdb").Collection("users")
    filter := bson.D{{"$text", bson.D{{"$search", "John"}}}}

    cursor, err := collection.Find(context.TODO(), filter)
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.TODO())

    var results []bson.M
    if err = cursor.All(context.TODO(), &results); err != nil {
        log.Fatal(err)
    }

    for _, result := range results {
        fmt.Printf("Found document: %v\n", result)
    }
}

// db.users.find({$text: {$search: "John"}})