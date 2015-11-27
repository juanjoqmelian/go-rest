package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/juanjoqmelian/go-rest/users/resources"
)

func main() {

	router := httprouter.New();

	router.GET("/users/:email", resources.GetUser)
	router.GET("/users", resources.GetUsers)
	router.POST("/users", resources.NewUser)
	router.PUT("/users/:email", resources.NewUser)


	http.ListenAndServe(":8080", router)
}