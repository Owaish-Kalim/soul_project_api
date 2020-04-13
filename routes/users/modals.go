package users

import (
	"time" 
	// "github.com/dgrijalva/jwt-go"
)

type User struct {
	id 	int
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Token 	string `json:"token"`
}
