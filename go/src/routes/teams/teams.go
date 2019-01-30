package teams

import (
	"net/http"
	"fmt"
	// "data"
	// "encoding/json"
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/users").Subrouter()
	// subr.HandleFunc("/addTeam", addTeam).Methods("POST")
	// subr.HandleFunc("/getTeam", getTeam).Methods("GET")
	// subr.HandleFunc("/deleteTeam", deleteTeam).Methods("DELETE")
	// subr.HandleFunc("/editTeam", editTeam).Methods("PUT")

	// subr.HandleFunc("/deleteTeams", deleteTeams).Methods("DELETE")
	// subr.HandleFunc("/getAll", getAll).Methods("GET")
	subr.HandleFunc("/", def).Methods("GET")
}



// func addTeam(w http.ResponseWriter, req *http.Request) {
	
// 	//build object from request
// 	team := &data.Team{
// 		TeamName: req.FormValue("teamName"), 
// 		TeamLeader: req.FormValue("teamLeader"),
// 		TeamMembers: req.FormValue("teamMembers")}

// 	responseCode := data.AddTeam(user)

// 	if responseCode == 0 {
// 		//return good
// 		w.WriteHeader(http.StatusOK)
// 	} else {
// 		//return bad
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }

// func getUser(w http.ResponseWriter, req *http.Request) {

// 	user := data.GetUser(req.FormValue("username"))

// 	if user != nil{
// 		//return good
// 		w.WriteHeader(http.StatusOK)
// 		//return user as json
// 		json.NewEncoder(w).Encode(user)
// 	} else {
// 		//return bad
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }

// func deleteUser(w http.ResponseWriter, req *http.Request) {
// 	responseCode := data.DeleteUser(req.FormValue("username"))

// 	if responseCode == 0{
// 		//return good
// 		w.WriteHeader(http.StatusOK)
// 	} else {
// 		//return bad
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }

// func editUser(w http.ResponseWriter, req *http.Request) {
// 	//build object from request
// 	user := &data.User{
// 		Username: req.FormValue("username"), 
// 		Password: req.FormValue("password"),
// 		AccessLevel: req.FormValue("accessLevel")}

// 	responseCode := data.EditUser(user)

// 	if responseCode == 0 {
// 		//return good
// 		w.WriteHeader(http.StatusOK)
// 	} else {
// 		//return bad
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }

// func deleteUsers(w http.ResponseWriter, req *http.Request) {
// 	responseCode := data.DeleteUsers()

// 	if responseCode == 0 {
// 		//return good
// 		fmt.Fprintf(w, "Deleted all users")
// 	} else {
// 		//return bad
// 		fmt.Fprintf(w, "Error deleting all users!")
// 	}
// }

// func getAll(w http.ResponseWriter, req *http.Request) {
// 	users := data.GetUsers()

// 	if len(*users) > 0{
// 		//return good
// 		w.WriteHeader(http.StatusOK)
// 		//return user as json
// 		for _, user := range *users{
// 			json.NewEncoder(w).Encode(user)
// 		}
// 	} else {
// 		//return bad
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// }

func def(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"teamHandler hit!")
}


