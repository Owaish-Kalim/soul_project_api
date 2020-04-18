package team

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
	fmt.Println(role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		return Response{}, res
	}

	if role != "Admin" {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = err.Error()
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

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(team.MobileNo) == false {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Invalid Mobile No"
		return Response{}, res
	}

	team_role := TeamRole{}
	sqlStatement = `SELECT ("Role_Id") FROM slh_roles WHERE ("Role_Name" = $1);`

	row = config.Db.QueryRow(sqlStatement, team.Role)
	err = row.Scan(&team_role.Team_Has_Role_Id)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = err.Error()
		return response, res
	}

	curr_time := time.Now()
	team.CreatedAt = curr_time.Format("02-01-2006 3:4:5 PM")
	curr_time, err = time.Parse("2006-01-02", team.Joining_Date)

	team.Joining_Date = curr_time.Format("02 Apr 2020")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		return response, res
	}

	sqlStatement = `
	INSERT INTO slh_teams ("FirstName","LastName","Email", "Address", "JoiningDate", "CreatedAt", "Password", "MobileNo", "Status", "Role", "Gender")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING ("TeamId")`

	team.TeamId = 0
	err = config.Db.QueryRow(sqlStatement, team.FirstName, team.LastName, team.Email, team.Address, team.Joining_Date, team.CreatedAt, string(hashedPassword),
		team.MobileNo, team.Status, team.Role, team.Gender).Scan(&team.TeamId)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = err.Error()
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
		res.Message = err.Error()
		return response, res
	}

	sqlStatement = `
	INSERT INTO slh_team_has_role ("Team_Id", "FirstName", "LastName", "Team_Has_Role_Id","CreatedAt", "Status", "UpdatedAt")  
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = config.Db.Exec(sqlStatement, team.TeamId, team.FirstName, team.LastName, team_role.Team_Has_Role_Id, team.CreatedAt, team.Status, team_role.UpdatedAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return response, Shared.ErrorMsg{Message: err.Error()}
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
		return response, Shared.ErrorMsg{Message: err.Error()}
	}

	sqlStatement := `SELECT ("TeamId"), ("FirstName"), ("LastName"), ("Email"), ("Password"), ("Address"), ("MobileNo"), ("Status"), 
	("Role"), ("JoiningDate"), ("Gender") FROM slh_teams WHERE ("Email")=$1;`
	team := Team{}
	row := config.Db.QueryRow(sqlStatement, client.Email)
	err = row.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Password, &team.Address, &team.MobileNo, &team.Status,
		&team.Role, &team.Joining_Date, &team.Gender)

	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		return response, Shared.ErrorMsg{Message: err.Error()}
	case nil:
		eror := bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(client.Password))
		if eror != nil {
			w.WriteHeader(http.StatusForbidden)
			return response, Shared.ErrorMsg{Message: err.Error()}
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
			return response, Shared.ErrorMsg{Message: err.Error()}
		}
		sqlStatement := `UPDATE slh_teams SET "Token"=$1 WHERE "Email"=$2`
		_, err = config.Db.Exec(sqlStatement, tokenString, team.Email)
		if err != nil {
			return response, Shared.ErrorMsg{Message: err.Error()}
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
		return response, Shared.ErrorMsg{Message: err.Error()}
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
		res.Message = err.Error()
		return UpdateResponse{}, res
	}

	if role != "Admin" {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = err.Error()
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
		return member, Shared.ErrorMsg{Message: err.Error()}
	}

	userEmail := context.Get(r, middleware.Decoded)
	sqlStatement := ` UPDATE slh_teams SET "Status" = $1 WHERE ("Email") = $2`
	_, err = config.Db.Exec(sqlStatement, member.Status, userEmail)
	if err != nil {
		return member, Shared.ErrorMsg{Message: err.Error()}
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
	fmt.Println(role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		return []Response{}, res
	}

	if role != "Admin" {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Unauthorised User"
		return []Response{}, res
	}

	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Message = err.Error()
			return []Response{}, res
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Message = err.Error()
			return []Response{}, res
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}
	teamid := r.Form.Get("teamid")
	if teamid != "" {
		if err := Shared.ParseInt(r.Form.Get("teamid"), &q.TeamId); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Message = err.Error()
			return []Response{}, res
		}
	}

	q.FirstName = r.Form.Get("firstname")
	q.LastName = r.Form.Get("lastname")
	q.Email = r.Form.Get("email")
	q.MobileNo = r.Form.Get("mobileno")
	q.Status = r.Form.Get("status")
	q.Joining_Date = r.Form.Get("joining_date")

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
	AND ("JoiningDate") ILIKE ''|| $7 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $8 OFFSET $9`
	rows, err := config.Db.Query(sqlStatement, q.TeamId, q.FirstName, q.LastName, q.Email, q.MobileNo, q.Status, q.Joining_Date, q.Limit, offset)
	// fmt.Println(rows)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return teams, Shared.ErrorMessage{Message: err.Error()}
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

	sqlStatement = `SELECT COUNT(*) FROM slh_teams 
	WHERE ("TeamId")::text ilike  ''|| $1 ||'%'
	OR ("FirstName") ILIKE ''|| $2 ||'%' 
	AND ("LastName") ILIKE '' || $3 || '%' 
	AND ("Email") ILIKE '' ||$4|| '%'  
	AND ("MobileNo") ILIKE '' ||$5|| '%' 
	AND ("Status") ILIKE ''|| $6 ||'%' 
	AND ("JoiningDate") ILIKE ''|| $7 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.TeamId, q.FirstName, q.LastName, q.Email, q.MobileNo, q.Status, q.Joining_Date)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return teams, Shared.ErrorMessage{Message: err.Error()}
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
		res.Message = err.Error()
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
		return team, Shared.ErrorMessage{Message: err.Error()}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)
	userEmail = team.Email

	sqlStatement = ` UPDATE slh_teams SET "Password" = $1  WHERE ("Email") = $2`

	_, err = config.Db.Exec(sqlStatement, string(hashedPassword), userEmail)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	return team, Shared.ErrorMessage{Message: ""}
}

func TeamLogout(w http.ResponseWriter, r *http.Request) Shared.ErrorMessage {
	r.ParseForm()

	userEmail := context.Get(r, middleware.Decoded)
	fmt.Println(userEmail)
	sqlStatement := `UPDATE slh_teams SET "Token"=$1 WHERE "Email"=$2`
	_, err := config.Db.Exec(sqlStatement, "qw", userEmail)
	if err != nil {
		return Shared.ErrorMessage{Message: err.Error()}
	}
	return Shared.ErrorMessage{Message: "Successfully Logout"}
}

func UploadImage(w http.ResponseWriter, req *http.Request) (ImgResp, Shared.ErrorMessage) {

	resp := ImgResp{}
	req.ParseMultipartForm(2 * 1024 * 1024)

	fmt.Println(req.FormValue("email"))
	file, handler, err := req.FormFile("myfile")
	if err != nil {
		panic(err)
		return resp, Shared.ErrorMessage{Message: err.Error()}
	}
	defer file.Close()
	fmt.Println("File info")
	fmt.Println("File Name: ", handler.Filename)
	// fmt.Println("File Size: ",handler.Size())
	fmt.Println("File Type: ", handler.Header.Get("Content-Type"))

	fmt.Println(req.FormValue("email"))

	if handler.Header.Get("Content-Type") != "image/jpeg" {
		fmt.Println("Upload jpeg image")
		return resp, Shared.ErrorMessage{Message: "Upload jpg image"}
	}

	fmt.Println(handler.Header)
	tempFile, err2 := ioutil.TempFile("uploads", "upload-*.jpg")

	if err2 != nil {
		return resp, Shared.ErrorMessage{Message: err.Error()}
	}
	fmt.Println(3)
	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		return resp, Shared.ErrorMessage{Message: err.Error()}
	}
	tempFile.Write(fileBytes)
	fmt.Println(4)
	resp.Member_Image = handler.Filename
	fmt.Println(resp.Member_Image)
	userEmail := req.FormValue("email")
	fmt.Println(userEmail)
	sqlStatement := ` UPDATE slh_teams SET "Member_Image" = $1 WHERE ("Email") = $2`

	_, err = config.Db.Exec(sqlStatement, resp.Member_Image, userEmail)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	return resp, Shared.ErrorMessage{Message: ""}

}
