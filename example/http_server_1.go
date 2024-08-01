//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tianluoding/noc"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUser(ctx *noc.Context) {
	user := User{ID: 1, Name: "John Doe"}
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(ctx.W, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx.W.Header().Set("Content-Type", "application/json")
	ctx.W.Write(jsonData)
}

func registerUser(ctx *noc.Context) {
	var user User
	err := json.NewDecoder(ctx.R.Body).Decode(&user)
	if err != nil {
		http.Error(ctx.W, err.Error(), http.StatusBadRequest)
		return
	}
	defer ctx.R.Body.Close()
	fmt.Printf("Received user: %+v\n", user)
	ctx.W.WriteHeader(http.StatusOK)
}

func main() {
	server := noc.NewExampleServer("example")
	server.Route("GET", "/user", getUser)
	server.Route("POST", "/user", registerUser)
	log.Fatal(server.Start(":8080"))
}
