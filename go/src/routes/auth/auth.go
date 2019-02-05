package auth

import (
	"net/http"
	"data"
	"encoding/json"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/api/auth").Subrouter()

	subr.HandleFunc("/login", login).Methods("GET")
	subr.HandleFunc("/accessLevel", getAccessLevel).Methods("GET")
}

func login(w http.ResponseWriter, req *http.Request) {
	
	_, err := data.GetUser(req.FormValue("username"),req.FormValue("password"))

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Json{true})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}

func getAccessLevel(w http.ResponseWriter, req *http.Request) {
	user, err := data.GetUser(req.FormValue("username"))

	if err == nil{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.DataJson{true, bson.M{"accessLevel":user.AccessLevel}})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
	}
}