package data

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

//********************define collection**********************
var teamsCollection *mongo.Collection

func defineTeamsCollection(client *mongo.Client) {
	teamsCollection = client.Database("db").Collection("teams")
}

//********************Team CRUD operations**********************
func AddTeam(record *Team) error {
	_, err := teamsCollection.InsertOne(context.TODO(), record)
	if err != nil {
		return errors.New("error adding data for team : " + record.Teamname)
	} else {
		return nil
	}
}

func GetTeam(teamname string) (*Team, error) {
	filter := bson.M{"teamname": teamname}

	var result Team

	err := teamsCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, errors.New("error getting data for team : " + teamname)
	} else {
		return &result, nil
	}
}

func DeleteTeam(teamname string) error {
	filter := bson.M{"teamname": teamname}

	_, err := teamsCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("error deleting data for team : " + teamname)
	} else {
		return nil
	}
}

func EditTeam(new *Team) error {
	filter := bson.M{"teamname": new.Teamname}

	var result Team

	err := teamsCollection.FindOneAndReplace(context.TODO(), filter, new).Decode(&result)

	if err != nil {
		return errors.New("error updating data for team : " + new.Teamname)
	} else {
		return nil
	}
}

func DeleteTeams() error { //empty bson object is like a wildcard
	_, err := teamsCollection.DeleteMany(context.TODO(), bson.M{})

	if err != nil {
		return errors.New("error deleting teams data")
	} else {
		return nil
	}
}

func GetTeams() (*[]Team, error) {
	var teams []Team

	cursor, err := teamsCollection.Find(context.TODO(), bson.M{}, options.Find())
	defer cursor.Close(context.TODO())

	if err != nil {
		return &teams, errors.New("error getting data for all teams")
	} else {
		var elem Team
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&elem)
			if err != nil {
				return &teams, errors.New("error while getting data for teams : error decoding team")
			} else {
				teams = append(teams, elem)
			}
		}
	}
	return &teams, nil
}
