package projectAdministration

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
	subr := router.PathPrefix("/api/projectAdministration").Subrouter()

	subr.HandleFunc("/project", addProject).Methods("POST")
	subr.HandleFunc("/project", getProject).Methods("GET")
	subr.HandleFunc("/project", deleteProject).Methods("DELETE")
	subr.HandleFunc("/project", editProject).Methods("PUT")

	subr.HandleFunc("/addFiles", addFiles).Methods("PUT")
	subr.HandleFunc("/addTeams", addTeams).Methods("PUT")
	subr.HandleFunc("/addUsers", addUsers).Methods("PUT")

	subr.HandleFunc("/removeFiles", removeFiles).Methods("PUT")
	subr.HandleFunc("/removeTeams", removeTeams).Methods("PUT")
	subr.HandleFunc("/removeUsers", removeUsers).Methods("PUT")

	subr.HandleFunc("/projects", deleteProjects).Methods("DELETE")
	subr.HandleFunc("/projects", getAll).Methods("GET")
	subr.HandleFunc("/usersProjects", getUsersProjects).Methods("GET")
}

func addProject(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	project, err := data.GetProject(requestBody["projectname"].(string))

	if project != nil {
		utils.RespondWithError(w, errors.New("a project with projectname '"+requestBody["projectname"].(string)+"' already exists."))
		return
	}

	var leader data.User

	mapstructure.Decode(requestBody["projectlead"], &leader)

	project = &data.Project{
		Projectname: requestBody["projectname"].(string),
		Projectlead: leader.Username,
		Files:       nil,
		Teams:       nil,
		Users:       nil}

	err = data.AddProject(project)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func getProject(w http.ResponseWriter, req *http.Request) {
	project, err := data.GetProject(req.FormValue("projectname"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	//build the response
	response, err := data.BuildProjectResponse([]data.Project{*project})
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.ProjectJson{true, response})

}

func deleteProject(w http.ResponseWriter, req *http.Request) {

	err := data.DeleteProject(req.FormValue("projectname"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func editProject(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var leader data.User

	mapstructure.Decode(requestBody["projectlead"], &leader)

	project := &data.Project{
		Projectname: requestBody["projectname"].(string),
		Projectlead: leader.Username,
		Files:       nil,
		Teams:       nil,
		Users:       nil}

	err = data.EditProject(project)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func addFiles(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var files []data.File

	mapstructure.Decode(requestBody["files"], &files)

	var filenames []string

	for _, file := range files {
		filenames = append(filenames, file.Filename)
	}

	err = data.AddFiles(requestBody["projectname"].(string), &filenames)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func addTeams(w http.ResponseWriter, req *http.Request) {

	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var teams []data.Team

	mapstructure.Decode(requestBody["teams"], &teams)

	var teamnames []string

	for _, team := range teams {
		teamnames = append(teamnames, team.Teamname)
	}

	err = data.AddTeams(requestBody["projectname"].(string), &teamnames)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func addUsers(w http.ResponseWriter, req *http.Request) {

	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var users []data.User

	mapstructure.Decode(requestBody["users"], &users)

	var usernames []string

	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	err = data.AddUsers(requestBody["projectname"].(string), &usernames)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func removeFiles(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var files []data.File

	mapstructure.Decode(requestBody["files"], &files)

	var filenames []string

	for _, file := range files {
		filenames = append(filenames, file.Filename)
	}

	err = data.RemoveFiles(requestBody["projectname"].(string), &filenames)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func removeTeams(w http.ResponseWriter, req *http.Request) {

	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var teams []data.Team

	mapstructure.Decode(requestBody["teams"], &teams)

	var teamnames []string

	for _, team := range teams {
		teamnames = append(teamnames, team.Teamname)
	}

	err = data.RemoveTeams(requestBody["projectname"].(string), &teamnames)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func removeUsers(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var users []data.User

	mapstructure.Decode(requestBody["users"], &users)

	var usernames []string

	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	err = data.RemoveUsers(requestBody["projectname"].(string), &usernames)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func deleteProjects(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteProjects()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func getAll(w http.ResponseWriter, req *http.Request) {
	projects, err := data.GetProjects()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	response, err := data.BuildProjectResponse(*projects)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.ProjectJson{true, response})
}

func getUsersProjects(w http.ResponseWriter, req *http.Request) {
	projects, err := data.GetProjects()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	var newProjects []data.Project

	user, err := data.GetUser(req.FormValue("user"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	username := user.Username

	for _, project := range *projects {
		if project.Projectlead == username {
			newProjects = append(newProjects, project)
			continue
		}
		for _, user := range project.Users {
			if user == username {
				newProjects = append(newProjects, project)
				continue
			}
		}
		for _, team := range project.Users {
			if utils.UserInTeam(username, team) {
				newProjects = append(newProjects, project)
				continue
			}
		}
	}

	response, err := data.BuildProjectResponse(newProjects)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.ProjectJson{true, response})
}
