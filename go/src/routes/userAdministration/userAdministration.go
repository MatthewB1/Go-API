package userAdministration

import (
	"errors"
	"net/http"

	"data"
	"encoding/json"
	utils "routes"

	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router) {
	subr := router.PathPrefix("/api/userAdministration").Subrouter()

	subr.HandleFunc("/user", addUser).Methods("POST")
	subr.HandleFunc("/user", getUser).Methods("GET")
	subr.HandleFunc("/user", deleteUser).Methods("DELETE")
	subr.HandleFunc("/user", editUser).Methods("PUT")

	subr.HandleFunc("/users", deleteUsers).Methods("DELETE")
	subr.HandleFunc("/users", getAll).Methods("GET")
}

func addUser(w http.ResponseWriter, req *http.Request) {
	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	user, err := data.GetUser(requestBody["username"].(string))

	if user != nil {
		utils.RespondWithError(w, errors.New("a user with username '"+requestBody["username"].(string)+"' already exists."))
		return
	}

	user = &data.User{
		Username:    requestBody["username"].(string),
		Password:    requestBody["password"].(string),
		AccessLevel: requestBody["accessLevel"].(string)}

	err = data.AddUser(user)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})

}

func getUser(w http.ResponseWriter, req *http.Request) {

	user, err := data.GetUser(req.FormValue("username"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.UserJson{true, []data.User{*user}})

}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteUser(req.FormValue("username"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func editUser(w http.ResponseWriter, req *http.Request) {
	var request map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	user := &data.User{
		Username:    request["username"].(string),
		Password:    request["password"].(string),
		AccessLevel: request["accessLevel"].(string)}

	err = data.EditUser(user)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func deleteUsers(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteUsers()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func getAll(w http.ResponseWriter, req *http.Request) {
	users, err := data.GetUsers()
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.UserJson{true, *users})
}
