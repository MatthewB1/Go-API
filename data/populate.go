package main

import (
	"fmt"
	"context"
	"log"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main(){
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil{log.Fatal(err)}

	//insert users
	usersCollection := client.Database("db").Collection("users")

	matthew := &User{Username: "Matthew", Password: "admin", AccessLevel: "admin"}
	dan := &User{Username: "Dan", Password: "password", AccessLevel: "user"}
	john := &User{Username: "John", Password: "password", AccessLevel: "user"}
	steve := &User{Username: "Steve", Password: "password", AccessLevel: "user"}
	
	users := []interface{} {matthew, dan, john, steve}

	insertManyResult, err := usersCollection.InsertMany(context.TODO(), users)
	if err != nil {
    	log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//inser teams
	//code to insert teams...
}




















//******************************data models****************************
/*
{
    "username": "",
    "password": "",
    "accessLevel": ""
}
*/
type User struct {
	Username   string	`json:"username"`
	Password   string	`json:"password"`
	AccessLevel string	`json:"accessLevel"`
}
/*
{
    "teamName": "",
    "teamLeader": "",
    "teamMembers": []
}
*/
type Team struct {

}