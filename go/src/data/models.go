package data

import (
    // "gopkg.in/mgo.v2/bson"
)

/*
{
    "username": "",
    "password": "",
    "accessLevel": ""
}
*/
type User struct {
    // ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
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
    TeamMembers []User `json:"teamMembers"`
}