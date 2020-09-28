package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akwanmaroso/blogos/api/auth"
	"github.com/akwanmaroso/blogos/api/helpers/responses"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("")
		log.Printf("%s %s %s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

func SetMiddlewareCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Credentials", "false")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD")

		next(w, r)
	}
}
