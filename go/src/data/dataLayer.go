package data

import (
	"context"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client


func init(){
	//*************DB connection setup*****************************
	var err error
	client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil{log.Fatal(err)}

	//************Collections******************************
	defineUsersCollection(client)
	defineTeamsCollection(client)
	// projectsCollection := client.Database("db").Collection("projects")
	// filesCollection := client.Database("db").Collection("files")

	//*************************************************************
}