package teamAdministration

import (
	"data"
	"encoding/json"
	"errors"
	"net/http"
	utils "routes"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

func SubRouter(router *mux.Router) {
	subr := router.PathPrefix("/api/teamAdministration").Subrouter()

	subr.HandleFunc("/team", addTeam).Methods("POST")
	subr.HandleFunc("/team", getTeam).Methods("GET")
	subr.HandleFunc("/team", deleteTeam).Methods("DELETE")
	subr.HandleFunc("/team", editTeam).Methods("PUT")

	subr.HandleFunc("/teams", deleteTeams).Methods("DELETE")
	subr.HandleFunc("/teams", getAll).Methods("GET")
}

func addTeam(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	//check if team with name already exists
	team, err := data.GetTeam(requestBody["teamname"].(string))

	if team != nil {
		utils.RespondWithError(w, errors.New("a team with teamname '"+requestBody["teamname"].(string)+"' already exists."))
		return
	}

	var leader data.User

	mapstructure.Decode(requestBody["teamleader"], &leader)

	_, err = data.GetUser(leader.Username)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var users []data.User

	mapstructure.Decode(requestBody["teamMembers"], &users)

	var members []string

	for _, member := range users {
		_, err := data.GetUser(member.Username)
		if err != nil {
			utils.RespondWithError(w, err)
			return
		}
		members = append(members, member.Username)
	}

	team = &data.Team{
		Teamname:    requestBody["teamname"].(string),
		Teamleader:  leader.Username,
		TeamMembers: members}

	err = data.AddTeam(team)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func getTeam(w http.ResponseWriter, req *http.Request) {

	team, err := data.GetTeam(req.FormValue("teamname"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	response, err := BuildTeamResponse([]data.Team{*team})
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.TeamJson{true, response})
}

func deleteTeam(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteTeam(req.FormValue("teamname"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func editTeam(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var members []string

	for _, member := range requestBody["teamMembers"].([]interface{}) {
		members = append(members, member.(string))
	}

	team := &data.Team{
		Teamname:    requestBody["teamname"].(string),
		Teamleader:  requestBody["teamLeader"].(string),
		TeamMembers: members}

	err = data.EditTeam(team)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func deleteTeams(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteTeams()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func getAll(w http.ResponseWriter, req *http.Request) {
	teams, err := data.GetTeams()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	response, err := BuildTeamResponse(*teams)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.TeamJson{true, response})
}

func BuildTeamResponse(teams []data.Team) ([]data.TeamResponse, error) {

	var response []data.TeamResponse

	for _, team := range teams {
		leader, err := data.GetUser(team.Teamleader)
		if err != nil {
			return nil, err
		}
		var members []data.User

		for _, username := range team.TeamMembers {
			member, err := data.GetUser(username)
			if err != nil {
				return nil, err
			}
			members = append(members, *member)
		}
		response = append(response, data.TeamResponse{Teamname: team.Teamname, Teamleader: *leader, TeamMembers: members})
	}

	return response, nil
}
