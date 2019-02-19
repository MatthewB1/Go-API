package auth

import (
	"data"
	"encoding/json"
	"errors"
	"net/http"
	utils "routes"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	user, err := data.GetUser(requestBody["username"].(string), requestBody["password"].(string))
	if err != nil {
		utils.RespondWithError(w, errors.New("Unable to authenticate"))
		return
	}

	tokenString, err := newToken(user.AccessLevel, user.Username)

	w.WriteHeader(http.StatusOK)
	// w.Write([]byte(tokenString))
	json.NewEncoder(w).Encode(data.DataJson{true, bson.M{"token": tokenString}})
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

var mySigningKey = []byte("secret")

func newToken(accessLevel string, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	//store claims in map
	claims := token.Claims.(jwt.MapClaims)

	//set claims
	claims["admin"] = (accessLevel == "admin")
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	//sign with secret
	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}
