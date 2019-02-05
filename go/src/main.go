package main

import (
	"routes/userAdministration"
	"routes/teamAdministration"
	"routes/projectAdministration"
	"routes/fileAdministration"
	"routes/auth"
	"middleware"


	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	
)

func main() {
	router := mux.NewRouter()

	userAdministration.SubRouter(router)
	teamAdministration.SubRouter(router)
	fileAdministration.SubRouter(router)
	projectAdministration.SubRouter(router)
	auth.SubRouter(router)


	//attach middleware
	middleware.AttachMiddleware(router)

	handler := cors.Default().Handler(router)
	
	// router.PathPrefix("/api")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
