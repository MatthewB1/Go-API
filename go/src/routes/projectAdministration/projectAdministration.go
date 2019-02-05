package projectAdministration

import (
	"net/http"
	"strings"
	"data"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/api/projectAdministration").Subrouter()

	subr.HandleFunc("/project", addProject).Methods("POST")
	subr.HandleFunc("/project", getProject).Methods("GET")
	subr.HandleFunc("/project", deleteProject).Methods("DELETE")
	subr.HandleFunc("/project", editProject).Methods("PUT")


	subr.HandleFunc("/projects", deleteProjects).Methods("DELETE")
	subr.HandleFunc("/projects", getAll).Methods("GET")
}


func addProject(w http.ResponseWriter, req *http.Request) {

	var files []data.File

	filenames := strings.Split(req.FormValue("files"), ",")

	for _, filename := range filenames {
		file, err := data.GetFile(filename)
		if err == nil{
			files = append(files, *file)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	var teams []data.Team

	teamnames := strings.Split(req.FormValue("teams"), ",")

	for _, teamname := range teamnames {
		team, err := data.GetTeam(teamname)
		if err == nil{
			teams = append(teams, *team)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}	
	}

	var users []data.User

	usernames := strings.Split(req.FormValue("users"), ",")

	for _, username := range usernames {
		user, err := data.GetUser(username)
		if err == nil{
			users = append(users, *user)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}	
	}

	user, err := data.GetUser(req.FormValue("projectlead"))

	if err != nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}	

	project := &data.Project{
		Projectname: req.FormValue("projectname"),
		Projectlead: *user,
		Files: files,
		Teams: teams,
		Users: users}
	

	err = data.AddProject(project)

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getProject(w http.ResponseWriter, req *http.Request) {

	project, err := data.GetProject(req.FormValue("projectname"))

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ProjectJson{true, []data.Project{*project}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteProject(w http.ResponseWriter, req *http.Request) {

	err := data.DeleteProject(req.FormValue("projectname"))

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func editProject(w http.ResponseWriter, req *http.Request) {

	var files []data.File

	filenames := strings.Split(req.FormValue("files"), ",")

	for _, filename := range filenames {
		file, err := data.GetFile(filename)
		if err == nil{
			files = append(files, *file)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}
	}

	var teams []data.Team

	teamnames := strings.Split(req.FormValue("teams"), ",")

	for _, teamname := range teamnames {
		team, err := data.GetTeam(teamname)
		if err == nil{
			teams = append(teams, *team)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}	
	}

	var users []data.User

	usernames := strings.Split(req.FormValue("users"), ",")

	for _, username := range usernames {
		user, err := data.GetUser(username)
		if err == nil{
			users = append(users, *user)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
		}	
	}

	user, err := data.GetUser(req.FormValue("projectlead"))

	if err != nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}	

	project := &data.Project{
		Projectname: req.FormValue("projectname"),
		Projectlead: *user,
		Files: files,
		Teams: teams,
		Users: users}
	

	err = data.EditProject(project)

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteProjects(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteProjects()

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	projects, err := data.GetProjects()

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ProjectJson{true, *projects})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}
