package middleware

import (
	"fmt"
	"net/http"

	utils "routes"

	"github.com/go-chi/jwtauth"
	"github.com/gorilla/mux"
)

func AttachMiddleware(router *mux.Router) {
	router.Use(headerMW)

	//can't get the error handling right on this :~(
	//causes : runtime error: invalid memory address or nil pointer dereference
	//from returning I think
	// router.Use(authToken)
}

//middleware

func headerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("request: ", req.Method, req.RequestURI)
		next.ServeHTTP(w, req)
	})
}

//checks for valid token on request
func authToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			fmt.Printf("some kind of error :~(  : %s", err)
			utils.RespondWithError(w, err)
			return
		}

		if token == nil || !token.Valid {
			fmt.Printf("oh dear!!!")
			utils.RespondWithError(w, err)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}
