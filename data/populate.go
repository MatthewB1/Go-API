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

	//drop collections ahead of populating with test data
	usersCollection := client.Database("db").Collection("users")
	teamsCollection := client.Database("db").Collection("teams")
	filesCollection := client.Database("db").Collection("files")
	projectsCollection := client.Database("db").Collection("projects")

	dropErr := usersCollection.Drop(context.TODO())
	if dropErr != nil {
		log.Fatal(err)
	}
	dropErr = teamsCollection.Drop(context.TODO())
	if dropErr != nil {
		log.Fatal(err)
	}
	dropErr = filesCollection.Drop(context.TODO())
	if dropErr != nil {
		log.Fatal(err)
	}
	dropErr = projectsCollection.Drop(context.TODO())
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

	fmt.Println("Inserted multiple user documents: ", insertManyResult.InsertedIDs)

	//insert teams

	petrels := &Team{Teamname: "petrels", Teamleader: "Dan", TeamMembers: []string{dan.Username, steve.Username}}
	puffins := &Team{Teamname: "puffins", Teamleader: "Steve", TeamMembers: []string{matthew.Username, dan.Username}}

	teams := []interface{}{petrels, puffins}

	insertManyResult, err = teamsCollection.InsertMany(context.TODO(), teams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple team documents: ", insertManyResult.InsertedIDs)

	//insert files

	coffee := &File{Filename: "coffee.jpg", Versions: []Version{Version{Lastsaved: "13/02/2019", Lasteditor: "Dave", TotaleditTime: "12312", Tags: []string{"tasty", "black", "hot"}}}}
	bovril := &File{Filename: "bovril.jpg", Versions: []Version{Version{Lastsaved: "13/02/2019", Lasteditor: "Steve", TotaleditTime: "45451", Tags: []string{"disgustang", "dunno what it is", "james may"}}, Version{Lastsaved: "13/02/2019", Lasteditor: "Steve", TotaleditTime: "3234", Tags: []string{"disgustang", "dunno what it is", "james may", "gravy?"}}}}

	files := []interface{}{coffee, bovril}

	insertManyResult, err = filesCollection.InsertMany(context.TODO(), files)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple file documents: ", insertManyResult.InsertedIDs)

	//insert projects
	project := &Project{Projectname: "lunch", Projectlead: "Dan", Files: []string{coffee.Filename, bovril.Filename}, Teams: []string{petrels.Teamname}, Users: []string{matthew.Username}}

	projects := []interface{}{project}

	insertManyResult, err = projectsCollection.InsertMany(context.TODO(), projects)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple project documents: ", insertManyResult.InsertedIDs)

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

/*
{
    "projectname": "",
    "projectlead": "",
    "teams": [],
    "users": []
}
*/

type Project struct {
	//ID    bson.ObjectId `bson:"_id, omitempty"`
	Projectname string   `json:"projectname`
	Projectlead string   `json:"projectlead`
	Files       []string `json:"files"`
	Teams       []string `json:"teams"`
	Users       []string `json:"users"`
}

/*
{
	"filename": "",
	"versions" :
    [
        {
        "lastsaved": "",
        "lasteditor": "",
        "versionNo": "",
        "totalEditTime": "",
        "tags": []
        }
    ]
}
*/

type File struct {
	Filename string    `json:"filename"`
	Versions []Version `json:"versions"`
}

type Version struct {
	//ID    bson.ObjectId `bson:"_id, omitempty"`
	Lastsaved     string   `json:"lastsaved"` //maybe change to time.Time
	Lasteditor    string   `json:"lasteditor"`
	TotaleditTime string   `json:"totaleditTime"`
	Tags          []string `json:"tags"`
}
