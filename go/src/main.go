package main

import (
	"routes/users"
	"routes/teams"
	"routes/auth"

	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	users.SubRouter(router)
	teams.SubRouter(router)


	auth.SubRouter(router)
	
	//use middleware functions
	router.Use(headerMW)
	
    log.Fatal(http.ListenAndServe(":8080", router))
}

func headerMW(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("adding headers and passing to handler for request: ", req.RequestURI)
		next.ServeHTTP(w,req)
	})
}


/*
// Define our struct
type authenticationMiddleware struct {
    tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
    amw.tokenUsers["00000000"] = "user0"
    amw.tokenUsers["aaaaaaaa"] = "userA"
    amw.tokenUsers["05f717e5"] = "randomUser"
    amw.tokenUsers["deadbeef"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Session-Token")

        if user, found := amw.tokenUsers[token]; found {
            // We found the token in our map
            log.Printf("Authenticated user %s\n", user)
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Forbidden", http.StatusForbidden)
        }
    })
}


*/