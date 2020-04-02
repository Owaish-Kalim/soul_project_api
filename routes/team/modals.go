package team

import (
	"time"
	// "github.com/dgrijalva/jwt-go"
)

type Team struct {
	TeamId       int    `json:"teamid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Password     string `json: "password", Db:"password"`
	Address      string `json:"address"`
	Token        string `json:"token"`
	MobileNo     string `json:"mobileno"`
	Status       string `json:"status"`
	Joining_Date time.Time
	CreatedAt    time.Time
}

type Response struct {
	TeamId       int    `json:"teamid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	MobileNo     string `json:"mobileno"`
	Status       string `json:"status"`
	Joining_Date time.Time
}

type LoginResponse struct {
	TeamId       int    `json:"teamid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	MobileNo     string `json:"mobileno"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	Joining_Date time.Time
}

type StatusResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}
