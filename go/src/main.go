package main

import (
	"routes/users"
	"routes/teams"

	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	users.SubRouter(router)
	teams.SubRouter(router)

	
    log.Fatal(http.ListenAndServe(":8080", router))
}
