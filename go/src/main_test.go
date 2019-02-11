//main_test.go

package main

import (
	"routes/userAdministration"
	"routes/auth"

	// "fmt"
	"context"
	"bytes"
    "time"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/gorilla/mux"
)

//helper function to return a new server, reduces code duplication
func newServer(r http.Handler, addr string) *http.Server{
return &http.Server{
        Addr:         addr,
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, 
	}
}

//****************HTTP request tests*****************************

//starts a server and closes it again
//assertions:
//1. that a valid server object is created
//2. that the server begins listening
//3. that the server is closed successfully
func TestStartServerAndListen(t *testing.T){
	r := mux.NewRouter()

	srv := newServer(r, "0.0.0.0:8001")
	
	if srv == nil {
		t.Errorf("Error creating server")
	}
	
	go func() {
        if err := srv.ListenAndServe(); err != nil {
            t.Errorf("Error listening: %v", err)
		}
		err := srv.Close()

		if err != nil {
			t.Errorf("Error closing server: %v", err)
		}
	}()
}


//starts a server, and fires a GET query to get all user data
//assertions:
//1. that the server started listening ok
//2. response code is ok
//3. response could be decoded
//4. repsonse included value 'true' for key "Success"
func TestGetRequest (t *testing.T) {
	r := mux.NewRouter()
	userAdministration.SubRouter(r)

	srv :=newServer(r, "0.0.0.0:8002")
	
	go func() {
        if err := srv.ListenAndServe(); err != nil {
            t.Errorf("Error listening: %v", err)
		}
		defer srv.Close()
	}()
	
	resp,err := http.Get(httptest.NewServer(r).URL + "/api/userAdministration/users")
	if err != nil {
		t.Fatalf("error sending request : %v", err)
	}
	
	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Got bad status code: %d", resp.StatusCode)
	}

	if err != nil {
		t.Errorf("Error decoding response : %s", err)
	}

	
	if response["Success"] != true {
		t.Errorf("Unexpected response : \n%v\n", response)
	}
}

// starts a server, creates a json structure and fires a POST request with login information
//assertions:
//1. server started listening ok
//2. json structure marshalled correctly
//3. POST request sent
//4. response code is ok
//5. response could be decoded
//6. repsonse included value 'true' for key "Success"
func TestPostRequest (t *testing.T) {
	r := mux.NewRouter()
	auth.SubRouter(r)

	srv :=newServer(r, "0.0.0.0:8003")
	
	go func() {
        if err := srv.ListenAndServe(); err != nil {
            t.Errorf("Error listening: %v", err)
		}
		defer srv.Close()
	}()

	jsonData, jsonErr := json.Marshal(map[string]string{"username": "Matthew", "password": "admin"})

	if jsonErr != nil {
		t.Errorf("Error creating json structure : %v", jsonErr)
	}
	
	resp,err := http.Post(httptest.NewServer(r).URL + "/api/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil { 
		t.Fatalf("error sending request : %v", err)
	}

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Got bad status code: %d", resp.StatusCode)
	}

	if err != nil {
		t.Errorf("Error decoding response : %s", err)
	}

	if response["Success"] != true {
		t.Errorf("Unexpected response : \n%v\n", response)
	}
}


//****************CRUD tests*****************************
//data model to decode results into
type Data struct {
	Name string
	Food string
}
//defines a client and collection, drops all data from the collection
//then adds a collection, and deletes that same collection
//assertions:
//1. client is established
//2. that the collection is successfully dropped
//3. that document count is 0 ahead of inserting
//4. that document is inserted correctly
//5. count is 1 after insertion
//6. document is deleted
//7. count is 0 after deletion
func TestAddDeleteData (t *testing.T) {
	//setup client and collection
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		t.Fatalf("Error establishing client connecton to db : %v", err)
	}

	defer client.Disconnect(context.TODO())

	testCollection := client.Database("db").Collection("tests")

	//drop collection ahead of tests
	dropErr := testCollection.Drop(context.TODO())

	if dropErr != nil {
		t.Errorf("Error dropping collection in preperation of tests : %v", dropErr)
	}

	//check count is 0
	preInsertCount, preInsertErr := testCollection.Count(context.TODO(), bson.M{})

	if preInsertCount != 0 {
		t.Errorf("Expected count 0, got count : %v", preInsertCount)
	}

	if preInsertErr != nil {
		t.Errorf("Error getting count before insert : %s", preInsertErr)
	}

	//insert a document, and then check count is 1
	_, insertErr := testCollection.InsertOne(context.TODO(), bson.M{"name": "test", "food": "apples"})

	if insertErr != nil {
		t.Errorf("Error inserting document : %s", insertErr)
	}

	postInsertCount, postInsertErr := testCollection.Count(context.TODO(), bson.M{})

	if postInsertCount != 1 {
		t.Errorf("Expected count 1, got count : %v", postInsertCount)
	}

	if postInsertErr != nil {
		t.Errorf("Error getting count after insert : %s", postInsertErr)
	}

	//delete the document just inserting, searching by field "name", then check count is 0
	_, deleteErr := testCollection.DeleteOne(context.TODO(), bson.M{"name": "test"})

	if deleteErr != nil {
		t.Errorf("Error deleting document : %s", deleteErr)
	}

	postDeleteCount, postDeleteErr := testCollection.Count(context.TODO(), bson.M{})

	if postDeleteCount != 0 {
		t.Errorf("Expected count 0, got count : %v", postDeleteCount)
	}

	if postDeleteErr != nil {
		t.Errorf("Error getting count after delete : %s", postDeleteErr)
	}
}
//defines a client and collection, drops all data from the collection
//then adds a collection, and deletes that same collection
//assertions:
//1. client is established
//2. that the collection is successfully dropped
//3. that document count is 0 ahead of inserting
//4. that document is inserted correctly
//5. count is 1 after insertion
//6. document is retrieved correctly
//7. field 'food' is as expected
//8. document is updated correctly
//9. document is retrieved correctly
//10. edited field has been reflected in database
func TestUpdateData (t *testing.T) {
	//setup client and collection
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")
	if err != nil {
		t.Fatalf("Error establishing client connecton to db : %v", err)
	}

	defer client.Disconnect(context.TODO())

	testCollection := client.Database("db").Collection("tests")

	//drop collection ahead of tests
	dropErr := testCollection.Drop(context.TODO())

	if dropErr != nil {
		t.Errorf("Error dropping collection in preperation of tests : %v", dropErr)
	}

	//check count is 0
	preInsertCount, preInsertErr := testCollection.Count(context.TODO(), bson.M{})

	if preInsertCount != 0 {
		t.Errorf("Expected count 0, got count : %v", preInsertCount)
	}

	if preInsertErr != nil {
		t.Errorf("Error getting count before insert : %s", preInsertErr)
	}

	//insert a document, and then check count is 1
	_, insertErr := testCollection.InsertOne(context.TODO(), bson.M{"name": "test", "food": "apples"})

	if insertErr != nil {
		t.Errorf("Error inserting document : %s", insertErr)
	}

	postInsertCount, postInsertErr := testCollection.Count(context.TODO(), bson.M{})

	if postInsertCount != 1 {
		t.Errorf("Expected count 1, got count : %v", postInsertCount)
	}

	if postInsertErr != nil {
		t.Errorf("Error getting count after insert : %s", postInsertErr)
	}

	var doc Data

	retrieveErr := testCollection.FindOne(context.TODO(), bson.M{"name": "test"}).Decode(&doc)

	if retrieveErr != nil {
		t.Errorf("Error retrieving and decoding document from collection : %s", retrieveErr)
	}

	if doc.Food != "apples" {
		t.Errorf("Expected value \"apples\", got : %v", doc.Food)
	}

	// dummy variable to satisfy promise of collection method
	var dummy interface{}

	updateErr := testCollection.FindOneAndReplace(context.TODO(), bson.M{"name": "test"}, bson.M{"name": "test", "food": "oranges"}).Decode(&dummy)

	if updateErr != nil {
		t.Errorf("Error decoding result of update : %s", updateErr)
	}

	retrieveErr = testCollection.FindOne(context.TODO(), bson.M{"name": "test"}).Decode(&doc)

	if retrieveErr != nil {
		t.Errorf("Error retrieving and decoding document from collection : %s", retrieveErr)
	}

	if doc.Food != "oranges" {
		t.Errorf("Expected value \"oranges\", got : %v", doc.Food)
	}
}