package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo-options"
	// "github.com/mongodb/mongo-go-driver/bson"
)




var client *mongo.Client


func init(){
	//*************DB connection setup*****************************
	var err error
	client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil{log.Fatal(err)}
	//*************************************************************
}

// func addRecord(collectionName string, user User) int{
// 	res,err := collection.InsertOne(context.TODO(), user)
// 	if err != nil{log.Fatal(err)}

// 	fmt.Println("Inserted document: ", res.InsertedID)

// 	return err
// }

// func deleteUsers() error{
// 	res,err := collection.DeleteMany(context.TODO(), nil)
// 	if  err != nil{
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("Deleted a%v documents in collection \"users\"", res.DeletedCount);

// 	return err 
// }
