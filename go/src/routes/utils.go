package utils

import (
	"data"
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
}

func UserInTeam(username string, teamname string) bool {
	team, err := data.GetTeam(teamname)
	if err != nil {
		return false
	}
	for _, user := range team.TeamMembers {
		if user == username {
			return true
		}
	}
	return false
}
