package data

import (
    // "gopkg.in/mgo.v2/bson"
    // "encoding/json"
    "github.com/mongodb/mongo-go-driver/bson"
    // "time"
)

//*****************Data models******************

/*
{
    "username": "",
    "password": "",
    "accessLevel": ""
}
*/
type User struct {
    // ID         string `json:"_id",omitempty`
	Username   string	`json:"username"`
	Password   string	`json:"password"`
	AccessLevel string	`json:"accessLevel"`
}
/*
{
    "teamName": "",
    "teamLeader": "",
    "teamMembers": []
}
*/
type Team struct {
    // ID      bson.ObjectId `bson:"_id,omitempty"`
    Teamname string `json:"teamname"`
    Teamleader string `json:"teamleader"`
    TeamMembers []User`json:"teamMembers"`
}
/*
{
    "projectname": "",
    "projectlead": "",
    "teams": [],
    "users": []
}
*/

type Project struct {
    //ID    bson.ObjectId `bson:"_id, omitempty"`
    Projectname string `json:"projectname`
    Projectlead User `json:"projectlead`
    Files []File `json:"files"`
    Teams []Team `json:"teams"`
    Users []User `json:"users"`
}
/*
{
    [
        {
        "filename": "",
        "lastsaved": "",
        "lasteditor": "",
        "versionNo": "",
        "totalEditTime": "",
        "tags": []
        }
    ]
}
*/

type File struct {
    Versions []Version `json:"versions"`
}

type Version struct {
    //ID    bson.ObjectId `bson:"_id, omitempty"`
    Filename string `json:"filename"`
    Lastsaved string `json:"lastsaved"` //maybe change to time.Time
    Lasteditor User `json:"lasteditor"`
    TotaleditTime string `json:"totaleditTime"`
    Tags []string `json:"tags"`
}


//**************Json data models*****************

type Json struct{
    Success bool `json:success`
}

type ErrorJson struct{
    Success bool `json:success`
    Error string `json:error`
}

type DataJson struct {
    Success bool `json:success`
    Data bson.M `json:data`
}

type UserJson struct{
    Success bool    `json:success`
    Data []User    `json:data,omitempty`
}

type TeamJson struct{
    Success bool    `json:success`
    Data []Team    `json:data,omitempty`
}

type ProjectJson struct{
    Success bool    `json:success`
    Data []Project    `json:data,omitempty`
}

type FileJson struct{
    
    Success bool    `json:success`
    Data []File    `json:data,omitempty`
}


