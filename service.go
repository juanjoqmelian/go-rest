package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/juanjoqmelian/go-rest/users/resources"
	"github.com/juanjoqmelian/go-rest/users/mongo"
)

const (
	MongoDbHost = "mongodb://192.168.99.100:27017"
)

func main() {

	mongo.GetConnection()

	router := httprouter.New();

	defaultUserWebService := resources.DefaultUserWebService{}

	router.GET("/users/:email", defaultUserWebService.GetUser)
	router.GET("/users", defaultUserWebService.GetUsers)
	router.POST("/users", defaultUserWebService.NewUser)
	router.PUT("/users/:email", defaultUserWebService.NewUser)

	http.ListenAndServe(":8080", router)
}