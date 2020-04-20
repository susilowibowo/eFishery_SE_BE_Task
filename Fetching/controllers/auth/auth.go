package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	// if token.Valid && err == nil {
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

// function verified for role "admin" only
func IsAdmin(c *gin.Context) {
	hmacSecretString := "secret"
	hmacSecret := []byte(hmacSecretString)
	tokenStr := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if token != nil && err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			if claims["Role"] == "admin" {
				result := gin.H{
					"Claims": claims,
				}
				c.JSON(http.StatusUnauthorized, result)
				fmt.Println("token verified")
			} else {
				result := gin.H{
					"message": "not Admin,  not authorized",
				}
				c.JSON(http.StatusUnauthorized, result)
				c.Abort()
			}
		}

	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

}
