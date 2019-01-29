package data

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
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
													//empty bson object is like a wildcard
	res,err := usersCollection.DeleteMany(context.TODO(), bson.M{})
	if  err != nil{
		fmt.Println("error deleting records: ", err)
		return 1
	} else {
		fmt.Printf("Deleted %v documents in collection \"users\"", res.DeletedCount)
		return 0
	}
}

// func GetUser(id string) *User int{

// 	filter := bson.D{{"_id":, id}}

// 	var result User

// 	err := usersCollection.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil{
// 		fmt.Println("error finding user: ", err)
// 		return result, 1
// 	} else {
// 		fmt.Printf("Found user in collection \"users\": ", result._id)
// 	}
// }


/*
	subr.HandleFunc("/addUser", addUser).Methods("GET")
	subr.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	subr.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")
	subr.HandleFunc("editUser/{id}", editUser).Methods("PUT")

	subr.HandleFunc("/deleteUsers", deleteAll).Methods("DELETE")
	subr.HandleFunc("/getAll", getAll).Methods("GET")
	subr.HandleFunc("/", def).Methods("GET")
*/