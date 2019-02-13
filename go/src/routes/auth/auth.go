package auth

import (
	"data"
	"encoding/json"
	"errors"
	"net/http"
	utils "routes"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson"
)

func SubRouter(router *mux.Router) {
	subr := router.PathPrefix("/api/auth").Subrouter()

	subr.HandleFunc("/login", login).Methods("POST")
	subr.HandleFunc("/accessLevel", getAccessLevel).Methods("GET")
}

func login(w http.ResponseWriter, req *http.Request) {

	var requestBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	_, err = data.GetUser(requestBody["username"].(string), requestBody["password"].(string))
	if err != nil {
		utils.RespondWithError(w, errors.New("Unable to authenticate"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Json{true})
}

func getAccessLevel(w http.ResponseWriter, req *http.Request) {
	user, err := data.GetUser(req.FormValue("username"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.DataJson{true, bson.M{"accessLevel": user.AccessLevel}})

}
