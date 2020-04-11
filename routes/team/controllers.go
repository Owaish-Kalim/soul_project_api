package team

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	// "html"
	// "net/url"
	// "net"
	"time"

	"soul_api/config"
	"soul_api/middleware"
	"soul_api/routes"

	"github.com/dgrijalva/jwt-go"
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
	// response.Role = team.Role
	response.Joining_Date = team.Joining_Date
	return *response
}

func CheckEmpty(team Team) Shared.ErrorMsg {
	if team.FirstName == "" {
		return Shared.ErrorMsg{Message: "FirstName cannot be empty."}
	}

	if team.LastName == "" {
		return Shared.ErrorMsg{Message: "LastName cannot be empty."}
	}

	if team.Email == "" {
		return Shared.ErrorMsg{Message: "Email cannot be empty."}
	}

	if team.Password == "" {
		return Shared.ErrorMsg{Message: "Password cannot be empty."}
	}

	if team.Address == "" {
		return Shared.ErrorMsg{Message: "Address cannot be empty."}
	}

	if team.MobileNo == "" {
		return Shared.ErrorMsg{Message: "MobileNo cannot be empty."}
	}

	if team.Status == "" {
		return Shared.ErrorMsg{Message: "Status cannot be empty."}
	}

	return Shared.ErrorMsg{Message: "nil"}
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

func CreateTeam(w http.ResponseWriter, r *http.Request) (Response, Shared.ErrorMsg) {
	r.ParseForm()
	team := Team{}
	response := Response{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}


	var res Shared.ErrorMsg
	res = CheckEmpty(team)
	if res.Message != "nil" {
		w.WriteHeader(http.StatusBadRequest)
		return response, res
	}

	team.Joining_Date = time.Now().Local()
	team.CreatedAt = time.Now().Local()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return response, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	sqlStatement := `
	INSERT INTO slh_teams ("FirstName","LastName","Email", "Address", "JoiningDate", "CreatedAt", "Password", "MobileNo", "Status")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING ("TeamId")`

	team.TeamId = 0
	err = config.Db.QueryRow(sqlStatement, team.FirstName, team.LastName, team.Email, team.Address, team.Joining_Date, team.CreatedAt, string(hashedPassword), team.MobileNo, team.Status).Scan(&team.TeamId)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		panic(err)
		return response, Shared.ErrorMsg{Message: "Email already registered"}
	}
	BuildResponse(&response, team)
	return response, Shared.ErrorMsg{Message: ""}
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
			return response, Shared.ErrorMsg{Message: "Bad Credentials"}
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
	// RETURN ARRAY REQUIRED & NOT REQUIRED

	if team.FirstName == "" || team.LastName == "" || team.Email == ""  || team.Password == "" || team.Address == "" || team.MobileNo == "" || team.Status == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return response, Shared.ErrorMsg{Message: "BLANK FIELDS"}
	}

	userEmail := context.Get(r, middleware.Decoded)

	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Email" = $3, "Password" = $4, "Address" = $5, "MobileNo" = $6, "Status" = $7, WHERE ("Email") = $8`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Email, team.Password, team.Address, team.MobileNo, team.Status, userEmail)
	if err != nil {
		panic(err)
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
	// RETURN ARRAY REQUIRED & NOT REQUIRED

	if team.FirstName == "" || team.LastName == "" || team.Email == "" || team.Password == "" || team.Address == "" || team.MobileNo == "" || team.Status == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return response, Shared.ErrorMsg{Message: "BLANK FIELDS"}
	}

	userEmail := team.Email

	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Email" = $3, "Password" = $4, "Address" = $5, "MobileNo" = $6, "Status" = $7 WHERE ("Email") = $8`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Email, team.Password, team.Address, team.MobileNo, team.Status, userEmail)
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
	limit := r.Form.Get("limit");
	if limit!="" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
	} 
	} else {
		q.Limit = 10;
	}
	page := r.Form.Get("page");
	if page!="" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			return response, Shared.ErrorMsg{Message: "parseerr"}
		} 
			q.Page = q.Page-1;
	} else {
		q.Page = 0;
	}
	teamid := r.Form.Get("teamid");
	if teamid!="" {
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
	sqlStatement := `SELECT ("TeamId"),("FirstName"),("LastName"),("Email"),("Address"),("MobileNo"), ("Status"),("JoiningDate") FROM slh_teams WHERE ("TeamId")=$1 OR ("FirstName") LIKE ''|| $2 ||'%' AND ("LastName") LIKE '' || $3 || '%' AND ("Email") LIKE '' ||$4|| '%' AND ("MobileNo") LIKE '' ||$5|| '%' AND ("Status") LIKE ''|| $6 ||'%' ORDER BY ("CreatedAt") DESC LIMIT $7 OFFSET $8`
	rows, err := config.Db.Query(sqlStatement,q.TeamId, q.FirstName, q.LastName, q.Email,q.MobileNo,q.Status,q.Limit,offset )

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
		return teams, Shared.ErrorMsg{Message: "Internal Server Error."}
	}
	cnt:=0

	// fmt.Println(len(rows))

	for rows.Next() {
		var team = Response{}
		rows.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Address, &team.MobileNo, &team.Status, &team.Joining_Date)
		teams = append(teams, team)
		cnt = cnt+1
	}
	
	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt/q.Limit 
	if cnt%q.Limit!=0 {
		totalPages =totalPages+1;
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	fmt.Println(cnt)
	return teams, Shared.ErrorMsg{Message: ""}
}
