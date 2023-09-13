package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowithgin/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not Found",
		})
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("gowithgin").C("users").FindId(oid).One(&u); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not Found",
		})
		return
	}
	user, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, user)
	fmt.Println(user)
}
func (uc UserController) CreateUser(c *gin.Context) {
	u := models.User{}

	json.NewDecoder(c.Request.Body).Decode(&u)

	u.ID = bson.NewObjectId()

	uc.session.DB("gowithgin").C("users").Insert(&u)

	user, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, user)
	fmt.Println(user)
}
func (uc UserController) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not Found",
		})
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("gowithgin").C("users").RemoveId(oid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not Found",
		})
		return
	}
	c.JSON(http.StatusOK, "Deleted user successfully")

}
