package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/susilowibowo/eFishery_SE_BE_Task/Fetching/structs"
)

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

//function show all user
func (idb *InDB) ShowUser(c *gin.Context) {
	var (
		// user   structs.User
		result gin.H
	)

	result = gin.H{
		"message": "showing all user",
	}

	c.JSON(http.StatusOK, result)
}
