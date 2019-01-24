package main

import (
	"context"
	"fmt"
	"log"

	// "encoding/json"
	// "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
    // "github.com/mongodb/mongo-go-driver/mongo/options"
)

type User struct {
	Username   string
	Password   string
	AccessLevel string
}

var client *mongo.Client
var collection *mongo.Collection


func init(){
	var err error
	client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil{log.Fatal(err)}
	collection = client.Database("db").Collection("users")
}

func addUser(user User) error{
	res,err := collection.InsertOne(context.TODO(), user)
	if err != nil{log.Fatal(err)}

	fmt.Println("Inserted document: ", res.InsertedID)

	return err
}

func deleteUsers() error{

	res,err := collection.DeleteMany(context.TODO(), nil)
	if  err != nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted a%v documents in collection \"users\"", res.DeletedCount);

	return err 
}


func main() {

}
