package Shared

import (
	// "time"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

func MobileVerify(Mobile string) string {
	fmt.Println(len(Mobile))

	if len(Mobile) > 10 || len(Mobile) < 10 {
		return "MobileNo Should be 10 digits"
	}

	fmt.Println(1)

	for i := 0; i < len(Mobile); i++ {
		// fmt.Println(Mobile[i])
		if Mobile[i] < '0' || Mobile[i] > '9' {
			return "MobileNo. should be in Number Form"
		}
	}

	fmt.Println(1)
	return "correct"

}

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
