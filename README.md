# Noc
a golang simple web server

## Usage

```go
func getUser(ctx *noc.Context) error {
	user := User{ID: 1, Name: "John Doe"}
	return ctx.WriteJSON(http.StatusOK, user)
}

func getUserByID(ctx *noc.Context) error {
	id := ctx.Param("id")
	numId, _ := strconv.Atoi(id)
	user := User{ID: numId, Name: "Jo Jo"}
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
    	// init logger
	logger.InitLogger()
    	// create noc server
	server := noc.NewNoc("example web")
    	// add route
	server.GET("/user", getUser)
    	// add route use :param
    	server.GET("/user/:id", getUserByID)
    	// post route
    	server.POST("/user", registerUser)
    	// start server
	if err := server.Start(":8080"); err != nil {
        	logger.Logger.Fatal(err)
    	}
}
```
