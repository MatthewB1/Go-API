package data

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/bson"
)


//********************define collection**********************
var projectsCollection *mongo.Collection

func defineProjectsCollection(client *mongo.Client){
	projectsCollection = client.Database("db").Collection("projects")
}

//********************User CRUD operations**********************
func AddProject(record *Project) error{
	_, err := projectsCollection.InsertOne(context.TODO(), record)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func GetProject(projectname string) (*Project, error){

	filter := bson.M{"projectname":projectname}

	var result Project

	err := projectsCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil{
		return nil, err
	} else {
		return &result, nil
	}
}

func DeleteProject(projectname string) error{
	filter := bson.M{"projectname":projectname}

	_, err := projectsCollection.DeleteOne(context.TODO(), filter)
	
	if err != nil{
		return err
	} else {
		return nil
	}
}

func EditProject(new *Project) error{
	filter := bson.M{"projectname":new.Projectname}

	var result Project

	err := projectsCollection.FindOneAndReplace(context.TODO(), filter, new).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func DeleteProjects() error{								//empty bson object is like a wildcard
	_,err := projectsCollection.DeleteMany(context.TODO(), bson.M{})

	if  err != nil{
		return err
	} else {
		return nil
	}
}

func GetProjects() (*[]Project, error){
	var projects []Project

	cursor, err := projectsCollection.Find(context.TODO(),  bson.M{}, options.Find())
	defer cursor.Close(context.TODO())

	if err != nil {
		return &projects, err
	} else{
		var elem Project
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&elem)
			if err != nil {
				return &projects, err
			} else {
				projects = append(projects, elem)
			}
		}
	}
	return &projects, err
}