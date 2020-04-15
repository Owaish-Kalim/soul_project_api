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

func CreateTeam(w http.ResponseWriter, r *http.Request) (Response, Shared.ErrorMsg) {

	userEmail := context.Get(r, middleware.Decoded)

	var role string
	var res Shared.ErrorMsg
	r.ParseForm()
	response := Response{}

	sqlStatement := `SELECT ("Role") FROM slh_teams WHERE ("Email" = $1);`
	row := config.Db.QueryRow(sqlStatement, userEmail)
	err := row.Scan(&role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "Internal Server Error."
		return Response{}, res
	}

	if role != "Admin" {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Unauthorised User"
		return Response{}, res
	}

	team := Team{}
	err = json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	res.Message = ""
	CheckEmpty(team, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return response, res
	}

	team_role := TeamRole{}
	sqlStatement = `SELECT ("Role_Id") FROM slh_roles WHERE ("Role_Name" = $1);`
	row = config.Db.QueryRow(sqlStatement, team.Role)
	err = row.Scan(&team_role.Team_Has_Role_Id)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Role_Name does not exist"
		return response, res
	}

	curr_time := time.Now()
	team.CreatedAt = curr_time.Format("02-01-2006 3:4:5 PM")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "Internal Server Error."
		return response, res
	}

	sqlStatement = `
	INSERT INTO slh_teams ("FirstName","LastName","Email", "Address", "JoiningDate", "CreatedAt", "Password", "MobileNo", "Status", "Role")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING ("TeamId")`

	team.TeamId = 0
	err = config.Db.QueryRow(sqlStatement, team.FirstName, team.LastName, team.Email, team.Address, team.Joining_Date, team.CreatedAt, string(hashedPassword), team.MobileNo, team.Status, team.Role).Scan(&team.TeamId)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Email already registered"
		return response, res
	}
	BuildResponse(&response, team)
	res.Message = ""

	team_role.UpdatedAt = curr_time.Format("02-01-2006 3:4:5 PM")
	sqlStatement = `SELECT ("Role_Id") FROM slh_roles WHERE ("Role_Name" = $1);`
	row = config.Db.QueryRow(sqlStatement, team.Role)
	err = row.Scan(&team_role.Team_Has_Role_Id)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Role_Name does not exist"
		return response, res
	}

	sqlStatement = `
	INSERT INTO slh_team_has_role ("Team_Id", "FirstName", "LastName", "Team_Has_Role_Id","CreatedAt", "Status", "UpdatedAt")  
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = config.Db.Exec(sqlStatement, team.TeamId, team.FirstName, team.LastName, team_role.Team_Has_Role_Id, team.CreatedAt, team.Status, team_role.UpdatedAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return response, Shared.ErrorMsg{Message: "Internal Server Error."}
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

	sqlStatement := `SELECT ("TeamId"), ("FirstName"), ("LastName"), ("Email"), ("Password"), ("Address"), ("MobileNo"), ("Status"), 
	("Role") FROM slh_teams WHERE ("Email")=$1;`
	team := Team{}
	row := config.Db.QueryRow(sqlStatement, client.Email)
	err = row.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Password, &team.Address, &team.MobileNo, &team.Status,
		&team.Role)

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

	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Address" = $3, "MobileNo" = $4, "Status" = $5, "Gender" = $6 
	WHERE ("Email") = $7`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Address, team.MobileNo, team.Status, team.Gender, userEmail)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return response, Shared.ErrorMsg{Message: "Internal Server Error."}
	}
	BuildUpdateResponse(&response, team)
	return response, Shared.ErrorMsg{Message: ""}
}

func UpdateMemberDetails(w http.ResponseWriter, r *http.Request) (UpdateResponse, Shared.ErrorMsg) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	userEmail := context.Get(r, middleware.Decoded)

	var role string
	var res Shared.ErrorMsg

	sqlStatement := `SELECT ("Role") FROM slh_teams WHERE ("Email" = $1);`
	row := config.Db.QueryRow(sqlStatement, userEmail)
	err := row.Scan(&role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "Internal Server Error."
		return UpdateResponse{}, res
	}

	if role != "Admin" {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Unauthorised User"
		return UpdateResponse{}, res
	}

	var team = Team{}
	var response = UpdateResponse{}
	err = json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	res.Message = ""
	CheckEmptyUp(team, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return response, res
	}

	userEmail = team.Email
	fmt.Println(userEmail)
	sqlStatement = ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Address" = $3, "MobileNo" = $4, "Status" = $5, "Gender" = $6 
	WHERE ("Email") = $7`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Address, team.MobileNo, team.Status, team.Gender, userEmail)
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
	sqlStatement := `SELECT ("TeamId"),("FirstName"),("LastName"),("Email"),("Address"),("MobileNo"), ("Status"),("Role"),("JoiningDate"), ("Gender")
	FROM slh_teams WHERE ("Email")=$1;`
	row := config.Db.QueryRow(sqlStatement, userEmail)

	row.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Address, &team.MobileNo, &team.Status, &team.Role, &team.Joining_Date,
		&team.Gender)

	return team, Shared.ErrorMsg{Message: ""}
}

func ListTeam(w http.ResponseWriter, r *http.Request) ([]Response, Shared.ErrorMessage) {

	userEmail := context.Get(r, middleware.Decoded)

	var role string
	var res Shared.ErrorMessage
	r.ParseForm()

	sqlStatement := `SELECT ("Role") FROM slh_teams WHERE ("Email" = $1);`
	row := config.Db.QueryRow(sqlStatement, userEmail)
	err := row.Scan(&role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "Internal Server Error"
		return []Response{}, res
	}

	if role != "Admin" {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Unauthorised User"
		return []Response{}, res
	}

	var response []Response
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			return response, Shared.ErrorMessage{Message: "parseerr"}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			return response, Shared.ErrorMessage{Message: "parseerr"}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}
	teamid := r.Form.Get("teamid")
	if teamid != "" {
		if err := Shared.ParseInt(r.Form.Get("teamid"), &q.TeamId); err != nil {
			return response, Shared.ErrorMessage{Message: "parseerr"}
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
	sqlStatement = `SELECT ("TeamId"),("FirstName"),("LastName"),("Email"),("Address"),("MobileNo"), ("Status"),("Role"),("JoiningDate"),
	("Gender") FROM slh_teams 
	WHERE ("TeamId")::text ilike  ''|| $1 ||'%'
	OR ("FirstName") ILIKE ''|| $2 ||'%' 
	AND ("LastName") ILIKE '' || $3 || '%' 
	AND ("Email") ILIKE '' ||$4|| '%'  
	AND ("MobileNo") ILIKE '' ||$5|| '%' 
	AND ("Status") ILIKE ''|| $6 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $7 OFFSET $8`
	rows, err := config.Db.Query(sqlStatement, q.TeamId, q.FirstName, q.LastName, q.Email, q.MobileNo, q.Status, q.Limit, offset)
	// fmt.Println(rows)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return teams, Shared.ErrorMessage{Message: "Internal Server Error."}
	}

	// fmt.Println(len(rows))
	for rows.Next() {
		var team = Response{}
		// fmt.Println(100)
		rows.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Address, &team.MobileNo, &team.Status, &team.Role, &team.Joining_Date,
			&team.Gender)
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
		return teams, Shared.ErrorMessage{Message: "Internal Server Error."}
	}

	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt / q.Limit
	if cnt%q.Limit != 0 {
		totalPages = totalPages + 1
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	return teams, Shared.ErrorMessage{Message: ""}
}

func UpdateMemberPassword(w http.ResponseWriter, r *http.Request) (Team, Shared.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")

	userEmail := context.Get(r, middleware.Decoded)

	var role string
	var res Shared.ErrorMessage
	r.ParseForm()

	sqlStatement := `SELECT ("Role") FROM slh_teams WHERE ("Email" = $1);`
	row := config.Db.QueryRow(sqlStatement, userEmail)
	err := row.Scan(&role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "Internal Server Error"
		return Team{}, res
	}

	if role != "Admin" {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Unauthorised User"
		return Team{}, res
	}

	var team = Team{}
	err = json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	if team.Password == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return team, Shared.ErrorMessage{Message: "BLANK FIELDS"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)
	userEmail = team.Email

	sqlStatement = ` UPDATE slh_teams SET "Password" = $1  WHERE ("Email") = $2`

	_, err = config.Db.Exec(sqlStatement, string(hashedPassword), userEmail)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	return team, Shared.ErrorMessage{Message: "Password Changed"}
}

func TeamLogout(w http.ResponseWriter, r *http.Request) Shared.ErrorMessage {
	r.ParseForm()

	userEmail := context.Get(r, middleware.Decoded)
	fmt.Println(userEmail)
	sqlStatement := `UPDATE slh_teams SET "Token"=$1 WHERE "Email"=$2`
	_, err := config.Db.Exec(sqlStatement, "qw", userEmail)
	if err != nil {
		return Shared.ErrorMessage{Message: "Internal Server Error."}
	}
	return Shared.ErrorMessage{Message: "Successfully Logout"}
}
