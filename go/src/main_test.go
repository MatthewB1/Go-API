//main_test.go

package main

import (
	"data"
	"fmt"
	"log"
	"os/exec"

	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client
var usersCollection *mongo.Collection
var teamsCollection *mongo.Collection
var filesCollection *mongo.Collection
var projectsCollection *mongo.Collection

func TestMain(m *testing.M) {
	// SETUP
	fmt.Printf("running setup....\n")
	err := setup()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("running tests....\n")
	//RUN TESTS
	m.Run()

	fmt.Printf("running teardown....\n")
	//TEARDOWN
	err = teardown()
	if err != nil {
		log.Fatal(err)
	}
}

func setup() error {
	_, err := exec.Command("go", "run", "../../data/populate.go").Output()
	if err != nil {
		return err
	}

	//*************DB connection setup*****************************
	client, err = mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		return err
	}

	usersCollection = client.Database("db").Collection("users")
	teamsCollection = client.Database("db").Collection("teams")
	filesCollection = client.Database("db").Collection("files")
	projectsCollection = client.Database("db").Collection("projects")

	return nil
}

func teardown() error {
	//drop collections
	err := usersCollection.Drop(context.TODO())
	if err != nil {
		return err
	}
	err = teamsCollection.Drop(context.TODO())
	if err != nil {
		return err
	}
	err = filesCollection.Drop(context.TODO())
	if err != nil {
		return err
	}
	err = projectsCollection.Drop(context.TODO())
	if err != nil {
		return err
	}
	client.Disconnect(context.TODO())
	return nil
}

//****************CRUD tests*****************************
//data model to decode results into
type Data struct {
	Name string
	Food string
}

func TestAddData(t *testing.T) {

	//get count
	preInsertUsers, err := data.GetUsers()

	if err != nil {
		t.Errorf("Error getting count before insert : %s", err)
	}

	//insert a document, and then check count again
	err = data.AddUser(&data.User{Username: "test", Password: "password", AccessLevel: "user"})

	if err != nil {
		t.Errorf("Error inserting document : %s", err)
	}

	postInsertUsers, err := data.GetUsers()

	if err != nil {
		t.Errorf("Error getting count after insert : %s", err)
	}

	difference := len(*postInsertUsers) - len(*preInsertUsers)

	if difference != 1 {
		t.Errorf("Expected count 1, got count : %v", difference)
	}
}

func TestDeleteData(t *testing.T) {

	//get count
	preInsertTeam, err := data.GetTeams()

	if err != nil {
		t.Errorf("Error getting count before insert : %s", err)
	}

	//insert a document, and then check count again
	err = data.DeleteTeam((*preInsertTeam)[0].Teamname)

	if err != nil {
		t.Errorf("Error deleting document : %s", err)
	}

	postInsertTeam, err := data.GetTeams()

	if err != nil {
		t.Errorf("Error getting count after insert : %s", err)
	}

	difference := len(*preInsertTeam) - len(*postInsertTeam)

	if difference != 1 {
		t.Errorf("Expected count 1, got count : %v", difference)
	}
}

func TestUpdateData(t *testing.T) {

	//get count
	preEditUsers, err := data.GetUsers()

	if err != nil {
		t.Errorf("Error getting users : %s", err)
	}

	//change a field
	(*preEditUsers)[0].Password = "crisps"

	//put the edited user back into db
	err = data.EditUser(&(*preEditUsers)[0])

	if err != nil {
		t.Errorf("Error editing document : %s", err)
	}

	postEditUsers, err := data.GetUsers()

	if err != nil {
		t.Errorf("Error getting users : %s", err)
	}

	if !cmp.Equal((*preEditUsers)[0], (*postEditUsers)[0]) {
		t.Errorf("Expected users 'Password' field to be different, got : %s & %s", (*preEditUsers)[0].Password, (*postEditUsers)[0].Password)
	}
}

func TestBuildResponse(t *testing.T) {
	teams, err := data.GetTeams()
	if err != nil {
		t.Errorf("Error getting teams : %s", err)
	}

	var teamnames []string

	for _, team := range *teams {
		teamnames = append(teamnames, team.Teamname)
	}

	project := &data.Project{Projectname: "test", Projectlead: "test", Files: nil, Teams: teamnames, Users: nil}

	response, err := data.BuildProjectResponse([]data.Project{*project})

	if err != nil {
		t.Errorf("Error building project response : %s", err)
	}
	teamResponse, err := data.BuildTeamResponse(*teams)
	if err != nil {
		t.Errorf("Error build team response : %s", err)
	}

	if !cmp.Equal(response[0].Teams, teamResponse) {
		t.Errorf("Expected team response objects and projects response object teams field to be equal")
	}
}
