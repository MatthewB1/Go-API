package data

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
)


//********************define collection**********************
var usersCollection *mongo.Collection

func defineUsersCollection(client *mongo.Client){
	usersCollection = client.Database("db").Collection("users")
}

//********************User CRUD operations**********************
func AddUser(record *User) int{
	res, err := usersCollection.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println("error inserting record")
		return 1
	} else{
	fmt.Println("Inserted document: ", res.InsertedID)
	return 0
	}
}

func DeleteUsers() int{
	res,err := usersCollection.DeleteMany(context.TODO(), nil)
	if  err != nil{
		fmt.Println("error deleting records")
		return 1
	} else {
		fmt.Printf("Deleted a%v documents in collection \"users\"", res.DeletedCount)
		return 0
	}
}


/*
	subr.HandleFunc("/addUser", addUser).Methods("GET")
	subr.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	subr.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")
	subr.HandleFunc("editUser/{id}", editUser).Methods("PUT")

	subr.HandleFunc("/deleteAll", deleteAll).Methods("DELETE")
	subr.HandleFunc("/getAll", getAll).Methods("GET")
	subr.HandleFunc("/", def).Methods("GET")
*/