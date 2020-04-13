package Shared

import (
	// "time"
	"strconv"

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

type ErrorMesg struct{
	Message   string `json:"message"`
	Customer_Souls_Id       string `json:"customers_souls_id"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	Status                  string `json:"status"`
	Commission_Amount       string    `json:"commission_amount"`
	Created_By              string `json:"created_by"`
	Updated_By              string `json:"updated_by"`
}
	

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
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
