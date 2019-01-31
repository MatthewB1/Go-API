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
		fmt.Println("request: ", req.Method, req.RequestURI)
		next.ServeHTTP(w,req)
	})
}