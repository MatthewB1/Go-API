package data

import (
	"context"

	"github.com/google/go-cmp/cmp"
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

func AddFiles(projectname string, files *[]File) error{
	filter := bson.M{"projectname":projectname}

	update, err := GetProject(projectname)

	for _, file := range *files{
		update.Files = append(update.Files, file)
	}


	var result Project

	err = projectsCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func AddTeams(projectname string, teams *[]Team) error{
	filter := bson.M{"projectname":projectname}

	update, err := GetProject(projectname)

	for _, team := range *teams{
		update.Teams = append(update.Teams, team)
	}


	var result Project

	err = projectsCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func AddUsers(projectname string, users *[]User) error{
	filter := bson.M{"projectname":projectname}

	update, err := GetProject(projectname)

	for _, user := range *users{
		update.Users = append(update.Users, user)
	}


	var result Project

	err = projectsCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func RemoveFiles(projectname string, files *[]File) error{
	filter := bson.M{"projectname":projectname}

	update, err := GetProject(projectname)

	for _, newfile := range *files{
		for i, file := range update.Files {
        	if cmp.Equal(newfile, file) {
				update.Files = append(update.Files[:i], update.Files[i+1:]...)
        	}
    	}
	}
	
	var result Project

	err = projectsCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func RemoveTeams(projectname string, teams *[]Team) error{
	filter := bson.M{"projectname":projectname}

	update, err := GetProject(projectname)

	for _, newteam := range *teams{
		for i, team := range update.Teams {
        	if cmp.Equal(newteam, team) {
				update.Teams = append(update.Teams[:i], update.Teams[i+1:]...)
        	}
    	}
	}


	var result Project

	err = projectsCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func RemoveUsers(projectname string, users *[]User) error{
	filter := bson.M{"projectname":projectname}

	update, err := GetProject(projectname)

	for _, newuser := range *users{
		for i, user := range update.Users {
        	if cmp.Equal(newuser, user) {
				update.Users = append(update.Users[:i], update.Users[i+1:]...)
        	}
    	}
	}

	var result Project

	err = projectsCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
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