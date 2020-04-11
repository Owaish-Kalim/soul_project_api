package team

import (
	Shared "soul_api/routes"
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
	Role      string
	Address   string
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type TeamRole struct {
	Team_Has_Role_Id int    `json:"team_has_role_id"`
	TeamId           int    `json:"teamid"`
	Status           string `json:"status"`
	UpdatedAt        time.Time
	CreatedAt        time.Time
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
}

func BuildResponse(response *Response, team Team) Response {
	response.TeamId = team.TeamId
	response.FirstName = team.FirstName
	response.LastName = team.LastName
	response.Email = team.Email
	response.Address = team.Address
	response.MobileNo = team.MobileNo
	response.Status = team.Status
	response.Role = team.Role
	response.Joining_Date = team.Joining_Date
	return *response
}

func BuildUpdateResponse(response *UpdateResponse, team Team) UpdateResponse {
	response.FirstName = team.FirstName
	response.LastName = team.LastName
	response.Email = team.Email
	response.Address = team.Address
	response.MobileNo = team.MobileNo
	response.Status = team.Status
	// response.Role = team.Role
	response.Joining_Date = team.Joining_Date
	return *response
}

func BuildLoginResponse(response *LoginResponse, team Team) LoginResponse {
	response.TeamId = team.TeamId
	response.FirstName = team.FirstName
	response.LastName = team.LastName
	response.Email = team.Email
	response.Address = team.Address
	response.MobileNo = team.MobileNo
	response.Status = team.Status
	// response.Role = team.Role
	response.Joining_Date = team.Joining_Date
	response.Token = team.Token
	return *response
}

func CheckEmpty(team Team, res *Shared.ErrorMsg) {

	if team.FirstName == "" {
		res.FirstName = "FirstName cannot be empty."
		res.Message = "Error"
	}

	if team.LastName == "" {
		res.LastName = "LastName cannot be empty."
		res.Message = "Error"
	}

	if team.Email == "" {
		res.Email = "Email cannot be empty."
		res.Message = "Error"
	}

	if team.Password == "" {
		res.Password = "Password cannot be empty."
		res.Message = "Error"
	}

	if team.Address == "" {
		res.Address = "Address cannot be empty."
		res.Message = "Error"
	}

	if team.MobileNo == "" {
		res.MobileNo = "MobileNo cannot be empty."
		res.Message = "Error"
	}

	if team.Status == "" {
		res.Status = "Status cannot be empty."
		res.Message = "Error"
	}

}

func CheckEmptyUp(team Team, res *Shared.ErrorMsg) {

	if team.FirstName == "" {
		res.FirstName = "FirstName cannot be empty."
		res.Message = "Error"
	}

	if team.LastName == "" {
		res.LastName = "LastName cannot be empty."
		res.Message = "Error"
	}

	if team.Address == "" {
		res.Address = "Address cannot be empty."
		res.Message = "Error"
	}

	if team.MobileNo == "" {
		res.MobileNo = "MobileNo cannot be empty."
		res.Message = "Error"
	}

	if team.Status == "" {
		res.Status = "Status cannot be empty."
		res.Message = "Error"
	}

}
