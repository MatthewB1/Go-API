package data

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/bson"
)


//********************define collection**********************
var usersCollection *mongo.Collection

func defineUsersCollection(client *mongo.Client){
	usersCollection = client.Database("db").Collection("users")
}

//********************User CRUD operations**********************
func AddUser(record *User) int{
	_, err := usersCollection.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return 1
	} else{
		return 0
	}
}

func GetUser(username string) *User{
	filter := bson.M{"username":username}

	var result User

	err := usersCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil{
		fmt.Println(err)
		return &result
	} else {
		return &result
	}
}

func DeleteUser(username string) int{
	filter := bson.M{"username":username}

	_, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil{
		fmt.Println(err)
		return 1
	} else {
		return 0
	}
}

func EditUser(new *User) int{
	filter := bson.M{"username":new.Username}

	update:= bson.D{{"$set", bson.M{
						"username" : new.Username,
						"password" : new.Password,
						"accessLevel" : new.AccessLevel}}}

	_, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return 1
	} else{
		return 0
	}
}

func DeleteUsers() int{								//empty bson object is like a wildcard
	res,err := usersCollection.DeleteMany(context.TODO(), bson.M{})
	if  err != nil{
		fmt.Println("error deleting records: ", err)
		return 1
	} else {
		fmt.Printf("Deleted %v documents in collection \"users\"", res.DeletedCount)
		return 0
	}
}

func GetUsers() *[]User{
	var users []User

	cursor, err := usersCollection.Find(context.TODO(),  bson.M{}, options.Find())
	defer cursor.Close(context.TODO())
	if err != nil {
		fmt.Println(err)
	} else{
		var elem User
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&elem)
			if err != nil {
				fmt.Println(err)
			} else {
				users = append(users, elem)
			}
		}
	}
	return &users
}


/*
	subr.HandleFunc("/addUser?", addUser).Methods("POST")
	subr.HandleFunc("/getUser?", getUser).Methods("GET")
	subr.HandleFunc("/deleteUser?", deleteUser).Methods("DELETE")
	subr.HandleFunc("editUser?", editUser).Methods("PUT")

	subr.HandleFunc("/deleteUsers", deleteAll).Methods("DELETE")
	subr.HandleFunc("/getUsers", getAll).Methods("GET")
	subr.HandleFunc("/", def).Methods("GET")
*/