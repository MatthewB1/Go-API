package users

import (
	"net/http"
	"fmt"
	"data"
	// "encoding/json"
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/users").Subrouter()
	subr.HandleFunc("/addUser", addUser).Methods("POST")
	subr.HandleFunc("/getUser/{id}", getUser).Methods("GET")
	subr.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")
	subr.HandleFunc("editUser/{id}", editUser).Methods("PUT")

	subr.HandleFunc("/deleteUsers", deleteUsers).Methods("DELETE")
	subr.HandleFunc("/getAll", getAll).Methods("GET")
	subr.HandleFunc("/", def).Methods("GET")

}



func addUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("passed to addUser handler...")
	user := &data.User{Username: "test", Password: "test", AccessLevel: "test"}
	
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
	//code...
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	//code...
}

func editUser(w http.ResponseWriter, req *http.Request) {
	//code...
}

func deleteUsers(w http.ResponseWriter, req *http.Request) {
	responseCode := data.DeleteUsers()

	if responseCode == 0 {
		//return good
		fmt.Fprintf(w, "Deleted all users")
	} else {
		//return bad
		fmt.Fprintf(w, "Error deleting all users!")
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	//code...
}

func def(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"userHandler hit!")
	//code...
}


