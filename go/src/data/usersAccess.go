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
func AddUser(record *User) error{
	_, err := usersCollection.InsertOne(context.TODO(), record)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func GetUser(strings ...string) (*User, error){

	var filter bson.M

	if len(strings) == 1{
		filter = bson.M{"username":strings[0]}
	}
	if len(strings) == 2{
		filter = bson.M{"username":strings[0], "password":strings[1]}
	}

	var result User

	res := usersCollection.FindOne(context.TODO(), filter)

	err := res.Decode(&result)

	// raw, decerr := res.DecodeBytes()

	// if decerr != nil {fmt.Println("error decoding")} 

	if err != nil{
		fmt.Println(err)
		return nil, err
	} else {
		return &result, nil
	}
}

func DeleteUser(username string) error{
	filter := bson.M{"username":username}

	_, err := usersCollection.DeleteOne(context.TODO(), filter)
	
	if err != nil{
		return err
	} else {
		return nil
	}
}

func EditUser(new *User) error{
	filter := bson.M{"username":new.Username}

	var result User

	err := usersCollection.FindOneAndReplace(context.TODO(), filter, new).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func DeleteUsers() error{								//empty bson object is like a wildcard
	_,err := usersCollection.DeleteMany(context.TODO(), bson.M{})

	if  err != nil{
		return err
	} else {
		return nil
	}
}

func GetUsers() (*[]User, error){
	var users []User

	cursor, err := usersCollection.Find(context.TODO(),  bson.M{}, options.Find())
	defer cursor.Close(context.TODO())

	if err != nil {
		return &users, err
	} else{
		var elem User
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&elem)
			if err != nil {
				return &users, err
			} else {
				users = append(users, elem)
			}
		}
	}
	return &users, err
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