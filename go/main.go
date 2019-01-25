package main

import (
	"context"
	"fmt"
	"log"
	"html"
	"net/http"
	// "encoding/json"
	
	// "github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	// "github.com/mongodb/mongo-go-driver/mongo/options"
	// "github.com/gorilla/mux"
)

type User struct {
	Username   string
	Password   string
	AccessLevel string
}

var client *mongo.Client
var collection *mongo.Collection


func init(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))



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
	// r := mux.NewRouter()
	// r.HandleFunc("/users", userHandler)
	/*
		figure out how to create handlers, and use them as middleware
		mux should support all of this
		reference express model for functionality required
		ignore sessions for now
	*/
}
