package data

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/bson"
)


//********************define collection**********************
var filesCollection *mongo.Collection

func defineFilesCollection(client *mongo.Client){
	filesCollection = client.Database("db").Collection("files")
}

//********************User CRUD operations**********************
func AddFile(record *File) error{
	_, err := filesCollection.InsertOne(context.TODO(), record)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func GetFile(filename string) (*File, error){

	filter := bson.M{"versions.filename":filename}

	var result File

	err := filesCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil{
		return nil, err
	} else {
		return &result, nil
	}
}

func DeleteFile(filename string) error{
	filter := bson.M{"versions.filename":filename}

	_, err := filesCollection.DeleteOne(context.TODO(), filter)
	
	if err != nil{
		return err
	} else {
		return nil
	}
}

func EditFile(new *File) error{
	filter := bson.M{"versions.filename":new.Versions[0].Filename}

	var result File

	err := filesCollection.FindOneAndReplace(context.TODO(), filter, new).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}
}

func AddFileVersion(new *Version) error{

	filter := bson.M{"versions.filename":new.Filename}

	update, err := GetFile(new.Filename)

	update.Versions = append(update.Versions, *new)

	var result File

	err = filesCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		return err
	} else{
		return nil
	}	
}

func DeleteFiles() error{								//empty bson object is like a wildcard
	_,err := filesCollection.DeleteMany(context.TODO(), bson.M{})

	if  err != nil{
		return err
	} else {
		return nil
	}
}

func GetFiles() (*[]File, error){
	var files []File

	cursor, err := filesCollection.Find(context.TODO(),  bson.M{}, options.Find())
	defer cursor.Close(context.TODO())

	if err != nil {
		return &files, err
	} else{
		var elem File
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&elem)
			if err != nil {
				return &files, err
			} else {
				files = append(files, elem)
			}
		}
	}
	return &files, err
}