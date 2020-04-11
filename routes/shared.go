package Shared

import (
	// "time"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key")

type ErrorMessage struct {
	Message string `json:"message"`
}

type ErrorMsg struct {
	Message   string `json:"message"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	MobileNo  string `json:"mobileno"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

func ParseBool(s string, dest *bool) error {
	// assume error = false
	*dest, _ = strconv.ParseBool(s)
	return nil
}

func ParseInt(s string, dest *int) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*dest = n
	return nil
}
