package teamHasRole

import (
	"time"
)

type TeamRole struct {
	Team_Has_Role_Id int    `json:"team_has_role_id"`
	TeamId           int    `json:"teamid"`
	Status           string `json:"status"`
	UpdatedAt        time.Time
	CreatedAt        time.Time
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
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

type RoleUp struct {
	Team_Has_Role_Id int    `json:"team_has_role_id"`
	TeamId           int    `json:"teamid"`
	Role_Id          int    `json:"role_id"`
	Role_Name        string `json:"role_name"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
}

type Roles struct {
	Role_Id   int    `json:"role_id"`
	Role_Name string `json:"role_name"`
	Status    bool   `json:"status"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
