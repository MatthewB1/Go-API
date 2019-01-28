package teams

import (
	"net/http"
	"fmt"
	
	
	"github.com/gorilla/mux"
)

func SubRouter(router *mux.Router){
	userRouter := router.PathPrefix("/teams").Subrouter()
	userRouter.HandleFunc("/", handleFunc).Methods("GET")
}

func handleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w,"teamsHandler hit!")
	//code...
}