package main

import (
	"middleware"
	"routes/auth"
	"routes/fileAdministration"
	"routes/projectAdministration"
	"routes/teamAdministration"
	"routes/userAdministration"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	//subrouters for routes
	userAdministration.SubRouter(router)
	teamAdministration.SubRouter(router)
	fileAdministration.SubRouter(router)
	projectAdministration.SubRouter(router)

	//attach middlewares
	middleware.AttachMiddleware(router)

	auth.SubRouter(router)

	//accept CORS requests
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
