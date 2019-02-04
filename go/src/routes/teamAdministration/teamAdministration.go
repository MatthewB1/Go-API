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
		userPointer := data.GetUser(username)
		if userPointer != nil{
			users = append(users, *userPointer)
		}
	}

	team := &data.Team{
		Teamname: req.FormValue("teamname"), 
		Teamleader: req.FormValue("teamleader"),
		TeamMembers: users}

	responseCode := data.AddTeam(team)

	if responseCode == 0 {
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getTeam(w http.ResponseWriter, req *http.Request) {

	team := data.GetTeam(req.FormValue("teamname"))

	if team != nil{
		//return good
		w.WriteHeader(http.StatusOK)
		//return user as json
		json.NewEncoder(w).Encode(team)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteTeam(w http.ResponseWriter, req *http.Request) {
	responseCode := data.DeleteTeam(req.FormValue("teamname"))

	if responseCode == 0{
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func editTeam(w http.ResponseWriter, req *http.Request) {

	var users []data.User

	usernames := strings.Split(req.FormValue("teamMembers"), ",")

	for _, username := range usernames {
		users = append(users, *data.GetUser(username))
	}
	
	team := &data.Team{
		Teamname: req.FormValue("teamname"), 
		Teamleader: req.FormValue("teamleader"),
		TeamMembers: users}

	responseCode := data.EditTeam(team)

	if responseCode == 0 {
		//return good
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(team)

	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteTeams(w http.ResponseWriter, req *http.Request) {
	responseCode := data.DeleteTeams()

	if responseCode == 0 {
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	teams := data.GetTeams()

	if len(*teams) > 0{
		//return good
		w.WriteHeader(http.StatusOK)
		//return user as json
		for _, team := range *teams{
			json.NewEncoder(w).Encode(team)
		}
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func usersInTeam(w http.ResponseWriter, req *http.Request) {
	team := data.GetTeam(req.FormValue("teamname"))

	if team != nil{
		//return good
		w.WriteHeader(http.StatusOK)
		//return user as json
		for _, user := range team.TeamMembers{
			json.NewEncoder(w).Encode(user)
		}
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}