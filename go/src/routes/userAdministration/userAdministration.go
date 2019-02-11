package userAdministration

import (
	"net/http"

	"data"
	"encoding/json"

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

	decErr := json.NewDecoder(req.Body).Decode(&requestBody)
	if decErr == nil {

		user := &data.User{
			Username:    requestBody["username"].(string),
			Password:    requestBody["password"].(string),
			AccessLevel: requestBody["accessLevel"].(string)}

		err := data.AddUser(user)

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

func getUser(w http.ResponseWriter, req *http.Request) {

	user, err := data.GetUser(req.FormValue("username"))

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.UserJson{true, []data.User{*user}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteUser(req.FormValue("username"))

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func editUser(w http.ResponseWriter, req *http.Request) {
	var request map[string]interface{}

	decErr := json.NewDecoder(req.Body).Decode(&request)
	if decErr == nil {

		user := &data.User{
			Username:    request["username"].(string),
			Password:    request["password"].(string),
			AccessLevel: request["accessLevel"].(string)}

		err := data.EditUser(user)

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

func deleteUsers(w http.ResponseWriter, req *http.Request) {
	err := data.DeleteUsers()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	users, err := data.GetUsers()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.UserJson{true, *users})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}
