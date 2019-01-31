package data

import (
	"context"
	"fmt"
	// "encoding/json"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/bson"
)


//********************define collection**********************
var teamsCollection *mongo.Collection

func defineTeamsCollection(client *mongo.Client){
	teamsCollection = client.Database("db").Collection("teams")
}

//********************Team CRUD operations**********************
func AddTeam(record *Team) int{
	_, err := teamsCollection.InsertOne(context.TODO(), record)
	if err != nil {
		fmt.Println(err)
		return 1
	} else{
		return 0
	}
}

func GetTeam(teamname string) *Team{
	filter := bson.M{"teamname":teamname}

	var result Team

	err := teamsCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil{
		fmt.Println(err)
		return nil
	} else {
		return &result
	}
}

func DeleteTeam(teamname string) int{
	filter := bson.M{"teamname":teamname}

	_, err := teamsCollection.DeleteOne(context.TODO(), filter)
	if err != nil{
		fmt.Println(err)
		return 1
	} else {
		return 0
	}
}

func EditTeam(new *Team) int{
	filter := bson.M{"teamname":new.Teamname}

	
	update := bson.D{{"$set", bson.M{"teamname":new.Teamname,"teamleader":new.Teamleader,"teamMembers":new.TeamMembers}}}

	//doesn't store changes to users[] in team :~(

	_, err := teamsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return 1

	} else{
		return 0
	}
}

func DeleteTeams() int{								//empty bson object is like a wildcard
	res,err := teamsCollection.DeleteMany(context.TODO(), bson.M{})
	if  err != nil{
		fmt.Println("error deleting records: ", err)
		return 1
	} else {
		fmt.Printf("Deleted %v documents in collection \"teams\"", res.DeletedCount)
		return 0
	}
}

func GetTeams() *[]Team{
	var teams []Team

	cursor, err := teamsCollection.Find(context.TODO(),  bson.M{}, options.Find())
	defer cursor.Close(context.TODO())
	if err != nil {
		fmt.Println(err)
	} else{
		var elem Team
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&elem)
			if err != nil {
				fmt.Println(err)
			} else {
				teams = append(teams, elem)
			}
		}
	}
	return &teams
}