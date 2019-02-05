package main

import(
	"context"
	// "fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/bson"
)

type User struct {
 	ID      	string `json:"id", omitempty` 
	Username   string	`json:"username"`
	Password   string	`json:"password"`
	}


func main(){
	client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil{log.Fatal(err)}

	collection := client.Database("db").collection("bsontest")	
}


	
