package team

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	"soul_api/middleware"
	Shared "soul_api/routes"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	// "github.com/gemcook/pagination-go"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

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

func CreateTeam(w http.ResponseWriter, r *http.Request) (Response, Shared.ErrorMsg) {
	r.ParseForm()
	team := Team{}
	response := Response{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	var res Shared.ErrorMsg
	res.Message = ""
	CheckEmpty(team, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return response, res
	}

	team.Joining_Date = time.Now().Local()
	team.CreatedAt = time.Now().Local()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "Internal Server Error."
		return response, res
	}

	sqlStatement := `
	INSERT INTO slh_teams ("FirstName","LastName","Email", "Address", "JoiningDate", "CreatedAt", "Password", "MobileNo", "Status", "Role")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING ("TeamId")`

	team.TeamId = 0
	err = config.Db.QueryRow(sqlStatement, team.FirstName, team.LastName, team.Email, team.Address, team.Joining_Date, team.CreatedAt, string(hashedPassword), team.MobileNo, team.Status, team.Role).Scan(&team.TeamId)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Email already registered"
		return response, res
	}
	BuildResponse(&response, team)
	res.Message = ""

	team_role := TeamRole{}
	team_role.UpdatedAt = time.Now().Local()

	sqlStatement = `SELECT ("Role_Id") FROM slh_roles WHERE ("Role_Name" = $1);`
	row := config.Db.QueryRow(sqlStatement, team.Role)
	err = row.Scan(&team_role.Team_Has_Role_Id)
	if err != nil {
		fmt.Println(1)
		panic(err)
	}

	sqlStatement = `
	INSERT INTO slh_team_has_role ("Team_Id","Team_Has_Role_Id","CreatedAt", "Status", "UpdatedAt")  
	VALUES ($1, $2, $3, $4, $5)`

	_, err = config.Db.Exec(sqlStatement, team.TeamId, team_role.Team_Has_Role_Id, team.CreatedAt, team.Status, team_role.UpdatedAt)
	if err != nil {
		fmt.Println(2)
		panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Email already registered"
		return response, res
	}

	return response, res
}

func LoginTeam(w http.ResponseWriter, r *http.Request) (LoginResponse, Shared.ErrorMsg) {
	r.ParseForm()
	var client = Team{}
	var response = LoginResponse{}
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		panic(err)
	}

	if client.Email == "" ||
		client.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return response, Shared.ErrorMsg{Message: "Fields cannot be empty."}
	}

	sqlStatement := `SELECT ("TeamId"), ("FirstName"), ("LastName"), ("Email"), ("Password"), ("Address"), ("MobileNo"), ("Status") FROM slh_teams WHERE ("Email")=$1;`
	team := Team{}
	row := config.Db.QueryRow(sqlStatement, client.Email)
	err = row.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Password, &team.Address, &team.MobileNo, &team.Status)

	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		return response, Shared.ErrorMsg{Message: "Email or Password Incorrect."}
	case nil:
		eror := bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(client.Password))
		if eror != nil {
			w.WriteHeader(http.StatusForbidden)
			return response, Shared.ErrorMsg{Message: "Email or Password Incorrect."}
		}

		expirationTime := time.Now().Add(15000 * time.Minute)
		claims := &Shared.Claims{
			Username: team.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(Shared.JwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return response, Shared.ErrorMsg{Message: "Internal Server Error."}
		}
		sqlStatement := `UPDATE slh_teams SET "Token"=$1 WHERE "Email"=$2`
		_, err = config.Db.Exec(sqlStatement, tokenString, team.Email)
		if err != nil {
			return response, Shared.ErrorMsg{Message: "Internal Server Error."}
		}
		team.Token = tokenString
		BuildLoginResponse(&response, team)
		return response, Shared.ErrorMsg{Message: ""}
	default:
		panic(err)
	}
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) (UpdateResponse, Shared.ErrorMsg) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var team = Team{}
	var response = UpdateResponse{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	var res Shared.ErrorMsg
	res.Message = ""
	CheckEmpty(team, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return response, res
	}

	userEmail := context.Get(r, middleware.Decoded)

	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Email" = $3, "Password" = $4, "Address" = $5, "MobileNo" = $6, "Status" = $7 WHERE ("Email") = $8`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Email, team.Password, team.Address, team.MobileNo, team.Status, userEmail)
	if err != nil {
		return response, Shared.ErrorMsg{Message: "Email already registered"}
		//panic(err)
	}
	BuildUpdateResponse(&response, team)
	return response, Shared.ErrorMsg{Message: ""}
}

func UpdateMemberDetails(w http.ResponseWriter, r *http.Request) (UpdateResponse, Shared.ErrorMsg) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var team = Team{}
	var response = UpdateResponse{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	var res Shared.ErrorMsg
	res.Message = ""
	CheckEmptyUp(team, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return response, res
	}

	userEmail := team.Email
	fmt.Println(userEmail)
	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Address" = $3, "MobileNo" = $4, "Status" = $5 WHERE ("Email") = $6`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Address, team.MobileNo, team.Status, userEmail)
	if err != nil {
		panic(err)
	}
	BuildUpdateResponse(&response, team)
	return response, Shared.ErrorMsg{Message: ""}
}

func UpdateTeamStatus(w http.ResponseWriter, r *http.Request) (StatusResponse, Shared.ErrorMsg) {

	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	var member = StatusResponse{}

	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		panic(err)
	}

	if member.Email == "" || member.Status == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return member, Shared.ErrorMsg{Message: "BLANK FIELDS"}
	}

	userEmail := context.Get(r, middleware.Decoded)
	sqlStatement := ` UPDATE slh_teams SET "Status" = $1 WHERE ("Email") = $2`
	_, err = config.Db.Exec(sqlStatement, member.Status, userEmail)
	if err != nil {
		return member, Shared.ErrorMsg{Message: "Internal Server Error."}
	}
	return member, Shared.ErrorMsg{Message: ""}
}

func ViewTeam(w http.ResponseWriter, r *http.Request) (Response, Shared.ErrorMsg) {
	r.ParseForm()
	var team = Response{}
	userEmail := context.Get(r, middleware.Decoded)
	fmt.Println(userEmail)
	sqlStatement := `SELECT ("TeamId"),("FirstName"),("LastName"),("Email"),("Address"),("MobileNo"), ("Status"),("JoiningDate") FROM slh_teams WHERE ("Email")=$1;`
	row := config.Db.QueryRow(sqlStatement, userEmail)

	row.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Address, &team.MobileNo, &team.Status, &team.Joining_Date)
	fmt.Println(team)
	return team, Shared.ErrorMsg{Message: ""}
}

func ViewTeamMember(w http.ResponseWriter, r *http.Request) (Response, Shared.ErrorMsg) {
	r.ParseForm()
	var team = Response{}
	userEmail := context.Get(r, middleware.Decoded)
	fmt.Println(userEmail)
	sqlStatement := `SELECT ("TeamId"),("FirstName"),("LastName"),("Email"),("Address"),("MobileNo"), ("Status"),("JoiningDate") FROM slh_teams WHERE ("Email")=$1;`
	row := config.Db.QueryRow(sqlStatement, userEmail)

	row.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Address, &team.MobileNo, &team.Status, &team.Joining_Date)
	fmt.Println(team)
	return team, Shared.ErrorMsg{Message: ""}
}

func ListTeam(w http.ResponseWriter, r *http.Request) ([]Response, Shared.ErrorMsg) {
	r.ParseForm()
	var response []Response
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}
	teamid := r.Form.Get("teamid")
	if teamid != "" {
		if err := Shared.ParseInt(r.Form.Get("teamid"), &q.TeamId); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
	}
	q.FirstName = r.Form.Get("firstname")
	q.LastName = r.Form.Get("lastname")
	q.Email = r.Form.Get("email")
	q.MobileNo = r.Form.Get("mobileno")
	q.Status = r.Form.Get("status")

	fmt.Println(q)
	offset := q.Limit * q.Page

	var teams []Response
	sqlStatement := `SELECT ("TeamId"),("FirstName"),("LastName"),("Email"),("Address"),("MobileNo"), ("Status"),("JoiningDate") FROM slh_teams 
	WHERE ("TeamId")=$1 OR ("FirstName") LIKE ''|| $2 ||'%' AND ("LastName") LIKE '' || $3 || '%' AND ("Email") LIKE '' ||$4|| '%' AND 
	("MobileNo") LIKE '' ||$5|| '%' AND ("Status") LIKE ''|| $6 ||'%' ORDER BY ("CreatedAt") DESC LIMIT $7 OFFSET $8`
	rows, err := config.Db.Query(sqlStatement, q.TeamId, q.FirstName, q.LastName, q.Email, q.MobileNo, q.Status, q.Limit, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return teams, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	// fmt.Println(len(rows))
	for rows.Next() {
		var team = Response{}
		rows.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Address, &team.MobileNo, &team.Status, &team.Joining_Date)
		teams = append(teams, team)
		// cnt = cnt + 1
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_teams WHERE ("TeamId")=$1 OR ("FirstName") LIKE ''|| $2 ||'%' AND ("LastName") LIKE '' || $3 || '%' AND 
	("Email") LIKE '' ||$4|| '%' AND ("MobileNo") LIKE '' ||$5|| '%' AND ("Status") LIKE ''|| $6 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.TeamId, q.FirstName, q.LastName, q.Email, q.MobileNo, q.Status)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return teams, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt / q.Limit
	if cnt%q.Limit != 0 {
		totalPages = totalPages + 1
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	fmt.Println(cnt)
	return teams, Shared.ErrorMsg{Message: ""}
}

func UpdateMemberPassword(w http.ResponseWriter, r *http.Request) (Team, Shared.ErrorMsg) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var team = Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	if team.Password == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return team, Shared.ErrorMsg{Message: "BLANK FIELDS"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)
	userEmail := team.Email

	sqlStatement := ` UPDATE slh_teams SET "Password" = $1  WHERE ("Email") = $2`

	_, err = config.Db.Exec(sqlStatement, string(hashedPassword), userEmail)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	return team, Shared.ErrorMsg{Message: "Password Changed"}
}

func TeamLogout(w http.ResponseWriter, r *http.Request) Shared.ErrorMsg {
	r.ParseForm()
	//var team = Response{}
	userEmail := context.Get(r, middleware.Decoded)
	fmt.Println(userEmail)
	sqlStatement := `UPDATE slh_teams SET "Token"=$1 WHERE "Email"=$2`
	_, err := config.Db.Exec(sqlStatement, "qw", userEmail)
	if err != nil {
		return Shared.ErrorMsg{Message: "Internal Server Error."}
	}
	return Shared.ErrorMsg{Message: "Successfully Logout"}
}

func TeamHasRole(w http.ResponseWriter, r *http.Request) ([]TeamRole, Shared.ErrorMsg) {
	r.ParseForm()
	var response []TeamRole
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}
	teamid := r.Form.Get("teamid")
	if teamid != "" {
		if err := Shared.ParseInt(r.Form.Get("teamid"), &q.TeamId); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		}
	}
	q.Status = r.Form.Get("status")

	fmt.Println(q)
	offset := q.Limit * q.Page

	var teamRoles []TeamRole
	fmt.Println(12)
	sqlStatement := `SELECT ("Team_Has_Role_Id"),("Team_Id"),("Status"),("CreatedAt"),("UpdatedAt") FROM slh_team_has_role 
	WHERE ("Status") LIKE ''|| $1 ||'%' ORDER BY ("CreatedAt") DESC LIMIT $2 OFFSET $3`
	rows, err := config.Db.Query(sqlStatement, q.Status, q.Limit, offset)
	fmt.Println(13)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return response, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	for rows.Next() {
		var team = TeamRole{}
		rows.Scan(&team.Team_Has_Role_Id, &team.TeamId, &team.Status, &team.CreatedAt, &team.UpdatedAt)
		teamRoles = append(teamRoles, team)
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_team_has_role WHERE ("Status") LIKE ''|| $1 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.Status)
	fmt.Println(14)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return teamRoles, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt / q.Limit
	if cnt%q.Limit != 0 {
		totalPages = totalPages + 1
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	fmt.Println(cnt)
	fmt.Println(15)
	return teamRoles, Shared.ErrorMsg{Message: ""}
}

func Role_Team(w http.ResponseWriter, r *http.Request) (Roles, ErrorMessage) {
	r.ParseForm()
	roles := Roles{}
	err := json.NewDecoder(r.Body).Decode(&roles)
	if err != nil {
		panic(err)
	}

	roles.Status = true

	sqlStatement := `
	INSERT INTO slh_roles ("Role_Name", "Role_Status") 
	VALUES ($1, $2)
	RETURNING ("Role_Id")`

	roles.Role_Id = 0
	err = config.Db.QueryRow(sqlStatement, roles.Role_Name, roles.Status).Scan(&roles.Role_Id)
	if err != nil {
		return Roles{}, ErrorMessage{Message: "Role_Name already Registered"}
	}

	return roles, ErrorMessage{}
}

func TeamHasRoleUpdate(w http.ResponseWriter, r *http.Request) (RoleUp, ErrorMessage) {
	fmt.Println(4)
	r.ParseForm()
	roles := RoleUp{}
	err := json.NewDecoder(r.Body).Decode(&roles)
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT ("Role_Id") FROM slh_roles WHERE ("Role_Name")=$1;`
	row := config.Db.QueryRow(sqlStatement, roles.Role_Name)
	err = row.Scan(&roles.Role_Id)
	if err != nil {
		fmt.Println(1)
		panic(err)
	}

	sqlStatement = ` UPDATE slh_team_has_role SET "Team_Has_Role_Id" = $1  WHERE ("Team_Id") = $2`

	_, err = config.Db.Exec(sqlStatement, roles.Role_Id, roles.TeamId)
	if err != nil {
		fmt.Println(2)
		panic(err)
	}

	sqlStatement = ` UPDATE slh_teams SET "Role" = $1  WHERE ("TeamId") = $2`

	_, err = config.Db.Exec(sqlStatement, roles.Role_Name, roles.TeamId)
	if err != nil {
		fmt.Println(3)
		panic(err)
	}

	return roles, ErrorMessage{}
}
