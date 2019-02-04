package middleware

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


func AttachMiddleware(router *mux.Router) {
	router.Use(headerMW)
}

//middleware

func headerMW(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//might be fussy about encoding
		fmt.Println("request: ", req.Method, req.RequestURI)
		next.ServeHTTP(w,req)
	})
}

func authMW(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//code...
		next.ServeHTTP(w,req)
	})
}