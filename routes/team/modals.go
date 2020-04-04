package team

import (
	"time"
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
	// Role         string `json:"role"`
	Joining_Date time.Time
	CreatedAt    time.Time
}

type Response struct {
	TeamId       int    `json:"teamid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	// Role         string `json:"role"`
	MobileNo     string `json:"mobileno"`
	Status       string `json:"status"`
	Joining_Date time.Time
}

type UpdateResponse struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	// Role         string `json:"role"`
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
	// Role         string `json:"role"`
	MobileNo     string `json:"mobileno"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	Joining_Date time.Time
}

type StatusResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}
