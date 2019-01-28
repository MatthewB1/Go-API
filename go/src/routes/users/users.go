package users

import (
	"net/http"
	"fmt"
	// "data/DataAccess"
	
	"encoding/json"
	
	
	"github.com/gorilla/mux"
)

type User struct {
	Username   string	`json:"username"`
	Password   string	`json:"password"`
	AccessLevel string	`json:"accessLevel"`
}

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/users").Subrouter()
	subr.HandleFunc("/addUser", addUser).Methods("GET")
	subr.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	subr.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")
	subr.HandleFunc("editUser/{id}", editUser).Methods("PUT")

	subr.HandleFunc("/deleteAll", deleteAll).Methods("DELETE")
	subr.HandleFunc("/getAll", getAll).Methods("GET")
	subr.HandleFunc("/", def).Methods("GET")
}

func addUser(w http.ResponseWriter, req *http.Request) {
	user := &User{Username: "test", Password: "test", AccessLevel: "test"}
	jsonUser,err := json.Marshal(user)
	if err != nil {
        fmt.Fprintf(w, "Error: %s", err)
    }

	fmt.Fprintf(w, "%T", jsonUser)

}

func getUser(w http.ResponseWriter, req *http.Request) {
	//code...
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	//code...
}

func editUser(w http.ResponseWriter, req *http.Request) {
	//code...
}

func deleteAll(w http.ResponseWriter, req *http.Request) {
	//code...
}

func getAll(w http.ResponseWriter, req *http.Request) {
	//code...
}

func def(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"userHandler hit!")
	//code...
}