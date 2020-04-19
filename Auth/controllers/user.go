package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Auth/structs"
)

// func for generate random password
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

// login user handler function
func (idb *InDB) LoginUser(c *gin.Context) {
	var (
		user   structs.User
		result gin.H
	)

	phone := c.PostForm("phone")
	password := c.PostForm("password")

	phone_check := idb.DB.Where("phone = ?", phone).Find(&user).RecordNotFound()

	if phone_check {
		result = gin.H{
			"message": "Phone not Found!",
		}
	} else {
		idb.DB.Where("phone = ?", phone).First(&user)

		if (user.Phone == phone) && (user.Password == password) {

			tk := &structs.Token{
				Name:     user.Name,
				Phone:    user.Phone,
				Role:     user.Role,
				Password: user.Password,
				Time:     time.Now(),
			}

			sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
			token, err := sign.SignedString([]byte("secret"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "some error, fail to convert jwt",
				})
				c.Abort()
			}
			result = gin.H{
				"token":   token,
				"message": "logged in",
			}
		} else {
			result = gin.H{
				"message": "failed log in",
			}
		}
	}
	c.JSON(http.StatusOK, result)
}
