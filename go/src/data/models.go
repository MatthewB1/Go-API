package data

import (
    // "gopkg.in/mgo.v2/bson"
    // "encoding/json"
	"github.com/mongodb/mongo-go-driver/bson"
)

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



