package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/GoesToEleven/golang-web-dev/040_mongodb/09_solution/controllers"
	"github.com/GoesToEleven/golang-web-dev/040_mongodb/09_solution/models"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}