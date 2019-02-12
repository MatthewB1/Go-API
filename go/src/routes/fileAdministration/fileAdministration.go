package fileAdministration

import (
	"data"
	"encoding/json"
	"net/http"

	utils "routes"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

func SubRouter(router *mux.Router) {
	subr := router.PathPrefix("/api/fileAdministration").Subrouter()

	subr.HandleFunc("/file", addFile).Methods("POST")
	subr.HandleFunc("/file", getFile).Methods("GET")
	subr.HandleFunc("/file", deleteFile).Methods("DELETE")

	subr.HandleFunc("/addFileVersion", addFileVersion).Methods("PUT")

	subr.HandleFunc("/files", deleteFiles).Methods("DELETE")
	subr.HandleFunc("/files", getAll).Methods("GET")
}

func addFile(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	//map[string]interface{}
	versionsBody := requestBody["versions"]

	var versions []data.Version
	var version data.Version

	var editor data.User
	var tags []string

	//iterate through all version objects
	for _, versionBody := range versionsBody.([]interface{}) {

		mapstructure.Decode(versionBody.(map[string]interface{})["lasteditor"], &editor)

		for _, tag := range versionBody.(map[string]interface{})["tags"].([]interface{}) {
			tags = append(tags, tag.(string))
		}

		version = data.Version{
			Lastsaved:     versionBody.(map[string]interface{})["lastsaved"].(string), //time formatting ??
			Lasteditor:    editor.Username,
			TotaleditTime: versionBody.(map[string]interface{})["totaleditTime"].(string),
			Tags:          tags}

		versions = append(versions, version)
		//clear slice
		tags = nil
	}

	file := &data.File{
		Filename: requestBody["filename"].(string),
		Versions: versions}

	err = data.AddFile(file)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func getFile(w http.ResponseWriter, req *http.Request) {

	file, err := data.GetFile(req.FormValue("filename"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.FileJson{true, []data.File{*file}})

}

func deleteFile(w http.ResponseWriter, req *http.Request) {

	err := data.DeleteFile(req.FormValue("filename"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func addFileVersion(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	versionBody := requestBody["version"]

	var editor data.User
	var tags []string

	mapstructure.Decode(versionBody.(map[string]interface{})["lasteditor"], &editor)

	for _, tag := range versionBody.(map[string]interface{})["tags"].([]interface{}) {
		tags = append(tags, tag.(string))
	}

	version := &data.Version{
		Lastsaved:     versionBody.(map[string]interface{})["lastsaved"].(string), //time formatting ??
		Lasteditor:    editor.Username,
		TotaleditTime: versionBody.(map[string]interface{})["totaleditTime"].(string),
		Tags:          tags}

	err = data.AddFileVersion(requestBody["filename"].(string), version)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func deleteFiles(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteFiles()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func getAll(w http.ResponseWriter, req *http.Request) {
	files, err := data.GetFiles()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.FileJson{true, *files})

}
