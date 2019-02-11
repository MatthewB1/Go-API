package auth

import (
	"net/http"
	"data"
	"encoding/json"
	"fmt"
	
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	subr := router.PathPrefix("/api/auth").Subrouter()

	subr.HandleFunc("/login", login).Methods("POST")
	subr.HandleFunc("/accessLevel", getAccessLevel).Methods("GET")
}

func login(w http.ResponseWriter, req *http.Request) {
	
	var requestBody map[string]interface{}

	decErr := json.NewDecoder(req.Body).Decode(&requestBody)
	if decErr != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.ErrorJson{false, decErr.Error()})
		fmt.Printf("Error decoding : %v", decErr)
	} else {

		_, err := data.GetUser(requestBody["username"].(string), requestBody["password"].(string))
	
		if err == nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.Json{true})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data.ErrorJson{false, err.Error()})
			fmt.Printf("Error authenticating : %v", err)
		}
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