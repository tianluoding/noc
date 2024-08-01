//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		user := User{ID: 1, Name: "John Doe"}
		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	case "POST":
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Printf("Received user: %+v\n", user)
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/user", handleUserRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
