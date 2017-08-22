package main

import (
	"encoding/json"
	"net/http"
)

func requestParser(next func(http.ResponseWriter, User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user, error := fetchJsonParams(r)

		if error != nil {

			Error.Println("Unable to parse json from request body, URL: " + r.URL.Path)

			errorJson, _ := json.Marshal(ServerError{"Server Error"})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorJson)

			return
		}

		next(w, user)
	}
}

func postRequestHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {

			Error.Println("Only POST request allowed on this URL: ", r.URL.Path)

			errorJson, _ := json.Marshal(ServerError{"Not Found"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorJson)

			return
		}
		next(w, r)
	}
}

func headerSetter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func requestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Info.Println("URL: ", r.URL.Path, "Request Type: ", r.Method)
		next(w, r)
	}
}
