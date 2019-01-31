package main

import (
	"routes/users"
	"routes/teams"
	"routes/auth"
	"middleware"

	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	users.SubRouter(router)
	teams.SubRouter(router)
	auth.SubRouter(router)

	//attach middleware
	middleware.AttachMiddleware(router)
	
    log.Fatal(http.ListenAndServe(":8080", router))
}


