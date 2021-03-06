package data

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

//********************define collection**********************
var filesCollection *mongo.Collection

func defineFilesCollection(client *mongo.Client) {
	filesCollection = client.Database("db").Collection("files")
}

//********************User CRUD operations**********************
func AddFile(record *File) error {
	_, err := filesCollection.InsertOne(context.TODO(), record)
	if err != nil {
		return errors.New("error adding data for file : " + record.Filename)
	} else {
		return nil
	}
}

func GetFile(filename string) (*File, error) {

	filter := bson.M{"filename": filename}

	var result File

	err := filesCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, errors.New("error getting data for file : " + filename)
	} else {
		return &result, nil
	}
}

func DeleteFile(filename string) error {
	filter := bson.M{"filename": filename}

	_, err := filesCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("error deleting data for file : " + filename)
	} else {
		return nil
	}
}

func EditFile(new *File) error {
	filter := bson.M{"filename": new.Filename}

	var result File

	err := filesCollection.FindOneAndReplace(context.TODO(), filter, new).Decode(&result)
	if err != nil {
		return errors.New("error updating data for file : " + new.Filename)
	} else {
		return nil
	}
}

func AddFileVersion(filename string, new *Version) error {

	filter := bson.M{"filename": filename}

	update, err := GetFile(filename)

	update.Versions = append(update.Versions, *new)

	var result File

	err = filesCollection.FindOneAndReplace(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		return errors.New("error adding new version to file : " + filename)
	} else {
		return nil
	}
}

func DeleteFiles() error { //empty bson object is like a wildcard
	_, err := filesCollection.DeleteMany(context.TODO(), bson.M{})

	if err != nil {
		return errors.New("error deleting data for all files")
	} else {
		return nil
	}
}

func GetFiles() (*[]File, error) {
	var files []File

	cursor, err := filesCollection.Find(context.TODO(), bson.M{}, options.Find())
	defer cursor.Close(context.TODO())

	if err != nil {
		return &files, errors.New("error getting data for all files")
	} else {
		var elem File
		for cursor.Next(context.TODO()) {
			err := cursor.Decode(&elem)
			if err != nil {
				return &files, errors.New("error while getting all files : unable to decode file")
			} else {
				files = append(files, elem)
			}
		}
	}
	return &files, err
}
