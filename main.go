package main

import (
	"fmt"
	"net/http"

	"github.com/gowithgin/controllers"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to initial call")
}

func setupRoutes(router *gin.Engine) {

	uc := controllers.NewUserController(getSession())

	router.GET("/", welcome)

	router.GET("/users/:id", uc.GetUser)
	router.POST("/users", uc.CreateUser)
	router.DELETE("/users/:id", uc.DeleteUser)
}
func main() {

	router := gin.Default()

	setupRoutes(router)

	fmt.Println(router.Run())
}

func getSession() *mgo.Session {
	// connecting to mongodb
	s, err := mgo.Dial("<<YOUR MONGO URL>>")

	if err != nil {
		panic(err)
	}

	return s
}
