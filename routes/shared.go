package Shared

import (
	// "time"
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key")

type ErrorMsg struct {
	Message   string `json:"message"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	MobileNo  string `json:"mobileno"`
	Status    string `json:"status"`
	Password  string `json: "password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
