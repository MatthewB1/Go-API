package data

import (
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
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessLevel string `json:"accessLevel"`
}

/*
{
    "teamName": "",
    "teamLeader": "",
    "teamMembers": []
}
*/
type Team struct {
	Teamname    string   `json:"teamname"`
	Teamleader  string   `json:"teamleader"`
	TeamMembers []string `json:"teamMembers"`
}

type TeamResponse struct {
	Teamname    string `json:"teamname"`
	Teamleader  User   `json:"teamleader"`
	TeamMembers []User `json:"teamMembers"`
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
	Projectname string   `json:"projectname`
	Projectlead string   `json:"projectlead`
	Files       []string `json:"files"`
	Teams       []string `json:"teams"`
	Users       []string `json:"users"`
}

type ProjectResponse struct {
	//ID    bson.ObjectId `bson:"_id, omitempty"`
	Projectname string         `json:"projectname`
	Projectlead User           `json:"projectlead`
	Files       []File         `json:"files"`
	Teams       []TeamResponse `json:"teams"`
	Users       []User         `json:"users"`
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
	Filename string    `json:"filename"`
	Versions []Version `json:"versions"`
}

type Version struct {
	//ID    bson.ObjectId `bson:"_id, omitempty"`
	Lastsaved     string   `json:"lastsaved"` //maybe change to time.Time
	Lasteditor    string   `json:"lasteditor"`
	TotaleditTime string   `json:"totaleditTime"`
	Tags          []string `json:"tags"`
}

//**************Json data models*****************

type Json struct {
	Success bool `json:success`
}

type ErrorJson struct {
	Success bool   `json:success`
	Error   string `json:error`
}

type DataJson struct {
	Success bool   `json:success`
	Data    bson.M `json:data`
}

type UserJson struct {
	Success bool   `json:success`
	Data    []User `json:data,omitempty`
}

type TeamJson struct {
	Success bool           `json:success`
	Data    []TeamResponse `json:data,omitempty`
}
type ProjectJson struct {
	Success bool              `json:success`
	Data    []ProjectResponse `json:data,omitempty`
}

type FileJson struct {
	Success bool   `json:success`
	Data    []File `json:data,omitempty`
}

//***************************functions***************************
func BuildTeamResponse(teams []Team) ([]TeamResponse, error) {

	var response []TeamResponse
	var members []User

	for _, team := range teams {
		leader, err := GetUser(team.Teamleader)
		if err != nil {
			return nil, err
		}

		for _, username := range team.TeamMembers {
			member, err := GetUser(username)
			if err != nil {
				return nil, err
			}
			members = append(members, *member)
		}
		response = append(response, TeamResponse{Teamname: team.Teamname, Teamleader: *leader, TeamMembers: members})
		members = nil
	}

	return response, nil
}

func BuildProjectResponse(projects []Project) ([]ProjectResponse, error) {
	var response []ProjectResponse
	var users []User
	var teams []Team
	var files []File

	for _, project := range projects {
		leader, err := GetUser(project.Projectlead)
		if err != nil {
			return nil, err
		}

		for _, filename := range project.Files {
			file, err := GetFile(filename)
			if err != nil {
				return nil, err
			}
			files = append(files, *file)
		}

		for _, teamname := range project.Teams {
			team, err := GetTeam(teamname)
			if err != nil {
				return nil, err
			}
			teams = append(teams, *team)
		}

		for _, username := range project.Users {
			user, err := GetUser(username)
			if err != nil {
				return nil, err
			}
			users = append(users, *user)
		}

		teamsResponse, err := BuildTeamResponse(teams)

		response = append(response,
			ProjectResponse{Projectname: project.Projectname,
				Projectlead: *leader,
				Files:       files,
				Teams:       teamsResponse,
				Users:       users})

		//clear slices
		files, teams, users = nil, nil, nil
	}

	return response, nil
}
