
package Shared

import (
	// "time" 
	"github.com/dgrijalva/jwt-go"

)

var JwtKey = []byte("my_secret_key")

type ErrorMsg struct {
	Message string `json:"message"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims	
}