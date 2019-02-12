package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	usersCollection := client.Database("db").Collection("users")

	teamsCollection := client.Database("db").Collection("teams")

	dropErr := usersCollection.Drop(context.TODO())
	if dropErr != nil {
		log.Fatal(err)
	}
	dropErr = teamsCollection.Drop(context.TODO())
	if dropErr != nil {
		log.Fatal(err)
	}

	//insert users

	matthew := &User{Username: "Matthew", Password: "admin", AccessLevel: "admin"}
	dan := &User{Username: "Dan", Password: "password", AccessLevel: "user"}
	john := &User{Username: "John", Password: "password", AccessLevel: "user"}
	steve := &User{Username: "Steve", Password: "password", AccessLevel: "user"}

	users := []interface{}{matthew, dan, john, steve}

	insertManyResult, err := usersCollection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//insert teams

	petrels := &Team{Teamname: "petrels", Teamleader: "Dan", TeamMembers: []string{dan.Username, steve.Username}}
	puffins := &Team{Teamname: "puffins", Teamleader: "Steve", TeamMembers: []string{matthew.Username, dan.Username}}

	teams := []interface{}{petrels, puffins}

	insertManyResult, err = teamsCollection.InsertMany(context.TODO(), teams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

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
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessLevel string `json:"accessLevel"`
}

/*
{
    "teamName": "",
    "teamLeader": "",
    "teamMembers": []
}
*/
type Team struct {
	Teamname    string   `json:"teamname"`
	Teamleader  string   `json:"teamleader"`
	TeamMembers []string `json:"teamMembers"`
}
