package auth

import (
	"net/http"
	"data"
	"encoding/json"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/auth").Subrouter()

	subr.HandleFunc("/login", login).Methods("GET")
	subr.HandleFunc("/accessLevel", getAccessLevel).Methods("GET")
}

func login(w http.ResponseWriter, req *http.Request) {
	
	user := data.GetUser(req.FormValue("username"),req.FormValue("password"))

	if user != nil {
		//return good
		w.WriteHeader(http.StatusOK)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getAccessLevel(w http.ResponseWriter, req *http.Request) {
	user := data.GetUser(req.FormValue("username"))

	if user != nil{
		//return good
		w.WriteHeader(http.StatusOK)

		ret := bson.M{"accessLevel":user.AccessLevel}
		
		json.NewEncoder(w).Encode(ret)
	} else {
		//return bad
		w.WriteHeader(http.StatusBadRequest)
	}
}