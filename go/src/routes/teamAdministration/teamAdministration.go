package teamAdministration

import (
	"data"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

func SubRouter(router *mux.Router) {
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
	var requestBody map[string]interface{}

	decErr := json.NewDecoder(req.Body).Decode(&requestBody)
	if decErr == nil {

		var leader data.User

		mapstructure.Decode(requestBody["teamLeader"], &leader)

		var users []data.User

		mapstructure.Decode(requestBody["teamMembers"], &users)

		team := &data.Team{
			Teamname:    requestBody["teamname"].(string),
			Teamleader:  leader,
			TeamMembers: users}

		err := data.AddTeam(team)

		if err == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.Json{true})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, decErr.Error()})
	}
}

func getTeam(w http.ResponseWriter, req *http.Request) {

	team, err := data.GetTeam(req.FormValue("teamname"))

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.TeamJson{true, []data.Team{*team}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteTeam(w http.ResponseWriter, req *http.Request) {

	err := data.DeleteTeam(req.FormValue("teamname"))

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func editTeam(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	decErr := json.NewDecoder(req.Body).Decode(&requestBody)
	if decErr == nil {

		var leader data.User

		mapstructure.Decode(requestBody["teamLeader"], &leader)

		var users []data.User

		mapstructure.Decode(requestBody["teamMembers"], &users)

		team := &data.Team{
			Teamname:    requestBody["teamname"].(string),
			Teamleader:  leader,
			TeamMembers: users}

		err := data.EditTeam(team)

		if err == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.Json{true})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, decErr.Error()})
	}
}

func deleteTeams(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteTeams()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	teams, err := data.GetTeams()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.TeamJson{true, *teams})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func usersInTeam(w http.ResponseWriter, req *http.Request) {

	team, err := data.GetTeam(req.FormValue("teamname"))

	if team != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.UserJson{true, team.TeamMembers})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}
