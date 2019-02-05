package teams

import (
	"net/http"
	"strings"
	"data"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/api/teamAdministration").Subrouter()

	subr.HandleFunc("/team", addTeam).Methods("POST")
	subr.HandleFunc("/team", getTeam).Methods("GET")
	subr.HandleFunc("/team", deleteTeam).Methods("DELETE")
	subr.HandleFunc("/team", editTeam).Methods("PUT")

	subr.HandleFunc("/usersInTeam", usersInTeam).Methods("GET")

	subr.HandleFunc("/teams", deleteTeams).Methods("DELETE")
	subr.HandleFunc("/teams", getAll).Methods("GET")
}


func addTeam(w http.ResponseWriter, req *http.Request) {

	var users []data.User

	usernames := strings.Split(req.FormValue("teamMembers"), ",")

	for _, username := range usernames {
		user, err := data.GetUser(username)
		if err == nil{
			users = append(users, *user)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}	
	}

	team := &data.Team{
		Teamname: req.FormValue("teamname"), 
		Teamleader: req.FormValue("teamleader"),
		TeamMembers: users}

	err := data.AddTeam(team)

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getTeam(w http.ResponseWriter, req *http.Request) {

	team, err := data.GetTeam(req.FormValue("teamname"))

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.TeamJson{true, []data.Team{*team}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteTeam(w http.ResponseWriter, req *http.Request) {

	err := data.DeleteTeam(req.FormValue("teamname"))

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func editTeam(w http.ResponseWriter, req *http.Request) {

	var users []data.User

	usernames := strings.Split(req.FormValue("teamMembers"), ",")

	for _, username := range usernames {
		user, err := data.GetUser(username)
		if err == nil{
			users = append(users, *user)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}	
	}
	
	team := &data.Team{
		Teamname: req.FormValue("teamname"), 
		Teamleader: req.FormValue("teamleader"),
		TeamMembers: users}

	err := data.EditTeam(team)

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteTeams(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteTeams()

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	teams, err := data.GetTeams()

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.TeamJson{true, *teams})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func usersInTeam(w http.ResponseWriter, req *http.Request) {

	team, err := data.GetTeam(req.FormValue("teamname"))

	if team != nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.UserJson{true, team.TeamMembers})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}