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
	Role         string `json:"role"`
	Gender       string `json:"gender"`
	Joining_Date time.Time
	CreatedAt    time.Time
}

type Response struct {
	TeamId       int    `json:"teamid"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Role         string `json:"role"`
	MobileNo     string `json:"mobileno"`
	Status       string `json:"status"`
	Joining_Date time.Time
}

type UpdateResponse struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	// Role         string `json:"role"`
	MobileNo     string `json:"mobileno"`
	Status       string `json:"status"`
	Joining_Date time.Time
}

type LoginResponse struct {
	TeamId    int    `json:"teamid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Address   string `json:"address"`
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

type query struct {
	Limit     int
	Page      int
	TeamId    int
	FirstName string
	LastName  string
	Email     string
	MobileNo  string
	Status    string
}

type TeamRole struct {
	Team_Has_Role_Id int    `json:"team_has_role_id"`
	TeamId           int    `json:"teamid"`
	Status           string `json:"status"`
	UpdatedAt        time.Time
	CreatedAt        time.Time
}

type RoleUp struct {
	Team_Has_Role_Id int    `json:"team_has_role_id"`
	TeamId           int    `json:"teamid"`
	Role_Id          int    `json:"role_id"`
	Role_Name        string `json:"role_name"`
}

type Roles struct {
	Role_Id   int    `json:"role_id"`
	Role_Name string `json:"role_name"`
	Status    bool   `json:"status"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
