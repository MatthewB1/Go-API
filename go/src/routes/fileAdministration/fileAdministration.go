package fileAdministration

import (
	"net/http"
	"strings"
	"data"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/api/fileAdministration").Subrouter()

	subr.HandleFunc("/file", addFile).Methods("POST")
	subr.HandleFunc("/file", getFile).Methods("GET")
	subr.HandleFunc("/file", deleteFile).Methods("DELETE")
	subr.HandleFunc("/file", editFile).Methods("PUT")


	subr.HandleFunc("/files", deleteFiles).Methods("DELETE")
	subr.HandleFunc("/files", getAll).Methods("GET")
}


func addFile(w http.ResponseWriter, req *http.Request) {

	var tags []string

	tagsSlice := strings.Split(req.FormValue("tags"), ",")

	for _, tag := range tagsSlice {
		tags = append(tags, tag)
	}

	user, err := data.GetUser(req.FormValue("username"))

	if err != nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}

	file := &data.File{
		Filename: req.FormValue("filename"),
		Lastsaved: req.FormValue("lastsaved"), //time formatting ??
		Lasteditor: *user,
		Editcount: req.FormValue("editcount"),
		TotaleditTime: req.FormValue("totaleditTime"),
		Tags: tags}

	err = data.AddFile(file)

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getFile(w http.ResponseWriter, req *http.Request) {

	file, err := data.GetFile(req.FormValue("filename"))

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.FileJson{true, []data.File{*file}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteFile(w http.ResponseWriter, req *http.Request) {

	err := data.DeleteFile(req.FormValue("filename"))

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func editFile(w http.ResponseWriter, req *http.Request) {

	var tags []string

	tagsSlice := strings.Split(req.FormValue("tags"), ",")

	for _, tag := range tagsSlice {
		tags = append(tags, tag)
	}

	user, err := data.GetUser(req.FormValue("username"))

	if err != nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}

	file := &data.File{
		Filename: req.FormValue("filename"),
		Lastsaved: req.FormValue("lastsaved"), //time formatting ??
		Lasteditor: *user,
		Editcount: req.FormValue("editcount"),
		TotaleditTime: req.FormValue("totaleditTime"),
		Tags: tags}

	err = data.EditFile(file)

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteFiles(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteFiles()

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	files, err := data.GetFiles()

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.FileJson{true, *files})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}
