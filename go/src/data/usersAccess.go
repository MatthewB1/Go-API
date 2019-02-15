package data

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

//********************define collection**********************
var usersCollection *mongo.Collection

func defineUsersCollection(client *mongo.Client) {
	usersCollection = client.Database("db").Collection("users")
}

//********************User CRUD operations**********************
func AddUser(record *User) error {
	_, err := usersCollection.InsertOne(context.TODO(), record)
	if err != nil {
		return errors.New("error inserting data for user : '" + record.Username + "'")
	} else {
		return nil
	}
}

func GetUser(strings ...string) (*User, error) {

	var filter bson.M

	if len(strings) == 1 {
		filter = bson.M{"username": strings[0]}
	}
	if len(strings) == 2 {
		filter = bson.M{"username": strings[0], "password": strings[1]}
	}

	var result User

	err := usersCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, errors.New("error finding data for user : '" + strings[0] + "'")
	} else {
		return &result, nil
	}
}

func DeleteUser(username string) error {
	filter := bson.M{"username": username}

	_, err := usersCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("error trying to delete user : '" + username + "'")
	} else {
		return nil
	}
}

func EditUser(new *User) error {
	filter := bson.M{"username": new.Username}

	var result User

	err := usersCollection.FindOneAndReplace(context.TODO(), filter, new).Decode(&result)
	if err != nil {
		return errors.New("error updating user : '" + new.Username + "'")
	} else {
		return nil
	}
}

func DeleteUsers() error { //empty bson object is like a wildcard
	_, err := usersCollection.DeleteMany(context.TODO(), bson.M{})

	if err != nil {
		return errors.New("error deleting all users")
	} else {
		return nil
	}
}

func GetUsers() (*[]User, error) {
	var users []User

	cursor, err := usersCollection.Find(context.TODO(), bson.M{}, options.Find())
	defer cursor.Close(context.TODO())

	if err != nil {
		return &users, errors.New("error getting all users")
	} else {
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
	return &users, nil
}
