//go:build ignore

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/tianluoding/noc"
	"github.com/tianluoding/noc/filter"
	"github.com/tianluoding/noc/logger"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func RegisterRoutes(server *noc.Noc) {
	server.GET("/user", getUser)
	server.POST("/user", registerUser)
	// 路径参数使用:key的形式
	server.GET("/user/:id", getUserByID)
}

func getUserByID(ctx *noc.Context) error {
	id := ctx.Param("id")
	numId, _ := strconv.Atoi(id)
	user := User{ID: numId, Name: "Jo Jo"}
	return ctx.WriteJSON(http.StatusOK, user)
}

func getUser(ctx *noc.Context) error {
	user := User{ID: 1, Name: "John Doe"}
	return ctx.WriteJSON(http.StatusOK, user)
}

func registerUser(ctx *noc.Context) error {
	var user User
	err := ctx.ReadJSON(&user)
	if err != nil {
		logger.Logger.Errorf("Error reading JSON: %v", err)
		return err
	}
	logger.Logger.Infof("Received user: %+v", user)
	return ctx.WriteJSON(http.StatusOK, nil)
}

func main() {
	logger.InitLogger()
	server := noc.NewNoc("example web")
	server.AddFilters(filter.MetricFilter)
	RegisterRoutes(server)

	log.Fatal(server.Start(":8080"))
}
