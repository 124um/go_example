package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name,omitempty"`
}

var (
	users = []User{{1, "Bonny"}, {2, "Maya"}, {33, ""}}
)

func mainServer() {
	http.HandleFunc("/users", authMiddleware(loggerMiddleware(handleUsers)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server listing :8080")
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("x-id")
		if userID == "" {
			log.Printf("[%s] %s - error: userID is not provided\n", r.Method,r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", userID)

		r= r.WithContext(ctx)
		next(w,r)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCtx := r.Context().Value("id")
		userID, ok := idFromCtx.(string)
		if !ok {
			log.Printf("[%s] %s - error: userID is invalid \n", r.Method,r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("[%s] %s - by userID %s \n", r.Method,r.URL, userID)
		next(w,r)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getUsers(w,r)
	case http.MethodPost:
		addUsers(w,r)
	default:
		w.WriteHeader(http.StatusNotImplemented)

	}


}

func getUsers(w http.ResponseWriter, r *http.Request){
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func addUsers(w http.ResponseWriter, r *http.Request){
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	if err = json.Unmarshal(reqBytes, &user); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users = append(users, user)
}