package users

import (
	"net/http"
	"fmt"
	"data"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/users").Subrouter()
	subr.HandleFunc("/addUser", addUser).Methods("POST")
	subr.HandleFunc("/getUser", getUser).Methods("GET")
	subr.HandleFunc("/deleteUser", deleteUser).Methods("DELETE")
	subr.HandleFunc("/editUser", editUser).Methods("PUT")

	subr.HandleFunc("/deleteUsers", deleteUsers).Methods("DELETE")
	subr.HandleFunc("/getAll", getAll).Methods("GET")
	subr.HandleFunc("/", def).Methods("GET")
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
		fmt.Fprintf(w, "Deleted all users")
	} else {
		//return bad
		fmt.Fprintf(w, "Error deleting all users!")
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

func def(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"userHandler hit!")
}


