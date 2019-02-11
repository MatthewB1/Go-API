package projectAdministration

import (
	"data"
	"encoding/json"
	"net/http"
	"strings"

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
}

func addProject(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	decErr := json.NewDecoder(req.Body).Decode(&requestBody)
	if decErr == nil {

		var leader data.User

		mapstructure.Decode(requestBody["projectlead"], &leader)

		project := &data.Project{
			Projectname: requestBody["projectname"].(string),
			Projectlead: leader,
			Files:       nil,
			Teams:       nil,
			Users:       nil}

		err := data.AddProject(project)

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

func getProject(w http.ResponseWriter, req *http.Request) {

	project, err := data.GetProject(req.FormValue("projectname"))

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ProjectJson{true, []data.Project{*project}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteProject(w http.ResponseWriter, req *http.Request) {

	err := data.DeleteProject(req.FormValue("projectname"))

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func editProject(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	decErr := json.NewDecoder(req.Body).Decode(&requestBody)
	if decErr == nil {

		var leader data.User

		mapstructure.Decode(requestBody["projectlead"], &leader)

		project := &data.Project{
			Projectname: requestBody["projectname"].(string),
			Projectlead: leader,
			Files:       nil,
			Teams:       nil,
			Users:       nil}

		err := data.EditProject(project)

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

func addFiles(w http.ResponseWriter, req *http.Request) {

	var files []data.File

	filenames := strings.Split(req.FormValue("files"), ",")

	for _, filename := range filenames {
		file, err := data.GetFile(filename)
		if err == nil {
			files = append(files, *file)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	err := data.AddFiles(req.FormValue("projectname"), &files)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func addTeams(w http.ResponseWriter, req *http.Request) {

	var teams []data.Team

	teamnames := strings.Split(req.FormValue("teams"), ",")

	for _, teamname := range teamnames {
		team, err := data.GetTeam(teamname)
		if err == nil {
			teams = append(teams, *team)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	err := data.AddTeams(req.FormValue("projectname"), &teams)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func addUsers(w http.ResponseWriter, req *http.Request) {

	var users []data.User

	usernames := strings.Split(req.FormValue("users"), ",")

	for _, username := range usernames {
		user, err := data.GetUser(username)
		if err == nil {
			users = append(users, *user)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	err := data.AddUsers(req.FormValue("projectname"), &users)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func removeFiles(w http.ResponseWriter, req *http.Request) {

	var files []data.File

	filenames := strings.Split(req.FormValue("files"), ",")

	for _, filename := range filenames {
		file, err := data.GetFile(filename)
		if err == nil {
			files = append(files, *file)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	err := data.RemoveFiles(req.FormValue("projectname"), &files)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func removeTeams(w http.ResponseWriter, req *http.Request) {

	var teams []data.Team

	teamnames := strings.Split(req.FormValue("teams"), ",")

	for _, teamname := range teamnames {
		team, err := data.GetTeam(teamname)
		if err == nil {
			teams = append(teams, *team)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	err := data.RemoveTeams(req.FormValue("projectname"), &teams)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func removeUsers(w http.ResponseWriter, req *http.Request) {

	var users []data.User

	usernames := strings.Split(req.FormValue("users"), ",")

	for _, username := range usernames {
		user, err := data.GetUser(username)
		if err == nil {
			users = append(users, *user)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	err := data.RemoveUsers(req.FormValue("projectname"), &users)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteProjects(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteProjects()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	projects, err := data.GetProjects()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ProjectJson{true, *projects})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}
