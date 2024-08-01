//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tianluoding/noc"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func RegisterRoutes(server *noc.ExampleServer) {
	server.GET("/user", getUser)
	server.POST("/user", registerUser)
}

func getUser(ctx *noc.Context) error {
	user := User{ID: 1, Name: "John Doe"}
	return ctx.WriteJSON(http.StatusOK, user)
}

func registerUser(ctx *noc.Context) error {
	var user User
	err := ctx.ReadJSON(&user)
	if err != nil {
		log.Printf("Error reading JSON: %v", err)
		return err
	}
	fmt.Printf("Received user: %+v\n", user)
	return ctx.WriteJSON(http.StatusOK, nil)
}

func main() {
	router := noc.NewMapRouter()
	server := noc.NewExampleServer("example", router)
	RegisterRoutes(server)
	log.Fatal(server.Start(":8080"))
}
