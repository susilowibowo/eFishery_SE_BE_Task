package structs

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Token struct declaration
type Token struct {
	Name     string    `json: "name" `
	Phone    string    `json: "phone" `
	Role     string    `json: "role" `
	Password string    `json: "password" `
	Time     time.Time `json: "timestamp" `
	*jwt.StandardClaims
}
