package users

import (
	"net/http"
	"data"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/api/userAdministration").Subrouter()

	subr.HandleFunc("/user", addUser).Methods("POST")
	subr.HandleFunc("/user", getUser).Methods("GET")
	subr.HandleFunc("/user", deleteUser).Methods("DELETE")
	subr.HandleFunc("/user", editUser).Methods("PUT")

	subr.HandleFunc("/users", deleteUsers).Methods("DELETE")
	subr.HandleFunc("/users", getAll).Methods("GET")
}



func addUser(w http.ResponseWriter, req *http.Request) {
	
	//build object from request
	user := &data.User{
		Username: req.FormValue("username"), 
		Password: req.FormValue("password"),
		AccessLevel: req.FormValue("accessLevel")}

	responseCode := data.AddUser(user)

	if responseCode == 0 {
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getUser(w http.ResponseWriter, req *http.Request) {

	user := data.GetUser(req.FormValue("username"))

	if user != nil{
		//return good
		w.WriteHeader(http.StatusOK)
		//return user as json
		json.NewEncoder(w).Encode(user)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	responseCode := data.DeleteUser(req.FormValue("username"))

	if responseCode == 0{
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func editUser(w http.ResponseWriter, req *http.Request) {
	//build object from request
	user := &data.User{
		Username: req.FormValue("username"), 
		Password: req.FormValue("password"),
		AccessLevel: req.FormValue("accessLevel")}

	responseCode := data.EditUser(user)

	if responseCode == 0 {
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteUsers(w http.ResponseWriter, req *http.Request) {
	responseCode := data.DeleteUsers()

	if responseCode == 0 {
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	users := data.GetUsers()

	if len(*users) > 0{
		//return good
		w.WriteHeader(http.StatusOK)
		//return user as json
		for _, user := range *users{
			json.NewEncoder(w).Encode(user)
		}
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}
