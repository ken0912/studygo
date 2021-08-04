package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tner struct {
	Id   int
	Name string
}

func main() {
	//Setup client
	clientOptions := options.Client().ApplyURI("mongodb://salarysql38.salarynet.local:27016,mongodb://salarysql38.salarynet.local:27017")

	//Connection mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//Specify the collection
	collection := client.Database("test").Collection("t")
	// fmt.Println("collection:", collection)

	var results []Tner
	// filter := bson.D{{"name", "KenTest4656"}}
	filter := options.Find().SetProjection(bson.M{"_id": 0})
	filter.SetSort(bson.D{{"id", 1}})
	filter.SetLimit(10)

	cursor, err := collection.Find(context.TODO(), bson.D{{}}, filter)
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result.Id, result.Name)
	}
	res, err := json.Marshal(results)
	fmt.Println("res:", string(res))

}
