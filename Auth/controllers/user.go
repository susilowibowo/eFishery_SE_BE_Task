package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Auth/structs"
)

// generate random password
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// func create user / check registered user
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	role := c.PostForm("role")
	// password := c.PostForm("password")

	name_check := idb.DB.Where("name = ?", name).Find(&user).RecordNotFound()

	// !(recordnotfound) = found!
	if !(name_check) {

		fmt.Println("choosed")
		result = gin.H{
			"result": user,
		}

	} else {
		rand.Seed(time.Now().UnixNano())
		user.Name = name
		user.Phone = phone
		user.Role = role
		//generated random password
		user.Password = randSeq(4)

		idb.DB.Create(&user)
		result = gin.H{
			"result": user,
		}

	}

	c.JSON(http.StatusOK, result)
}
