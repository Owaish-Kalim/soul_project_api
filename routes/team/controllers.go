package team

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	"soul_api/middleware"
	Shared "soul_api/routes"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"golang.org/x/crypto/bcrypt"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) (Team, Shared.ErrorMsg) {
	r.ParseForm()
	team := Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	if team.FirstName == "" ||
		team.LastName == "" ||
		team.Email == "" ||
		team.Password == "" ||
		team.Address == "" ||
		team.MobileNo == "" ||
		team.Status == "" {
		w.WriteHeader(http.StatusBadRequest)
		return team, Shared.ErrorMsg{Message: "Fields cannot be empty."}
	}

	team.Joining_Date = time.Now().Local()
	team.CreatedAt = time.Now().Local()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return team, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	sqlStatement := `
	INSERT INTO slh_teams ("FirstName","LastName","Email", "Address", "JoiningDate", "CreatedAt", "Password", "MobileNo", "Status")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING ("TeamId")`

	team.TeamId = 0
	err = config.Db.QueryRow(sqlStatement, team.FirstName, team.LastName, team.Email, team.Address, team.Joining_Date, team.CreatedAt, string(hashedPassword), team.MobileNo, team.Status).Scan(&team.TeamId)
	if err != nil {
		return team, Shared.ErrorMsg{Message: "Internal Server Error."}
	}

	team.Password = ""
	return team, Shared.ErrorMsg{Message: ""}
}

func ListTeam(w http.ResponseWriter, r *http.Request) ([]Response, Shared.ErrorMsg) {
	r.ParseForm()
	var teams []Response
	fmt.Println("asfgds")
	sqlStatement := `SELECT ("TeamId"),("FirstName"),("LastName"),("Email"),("Address"),("MobileNo"), ("Status"),("JoiningDate") FROM slh_teams WHERE 1=1;`
	fmt.Println(sqlStatement)
	rows, err := config.Db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var team = Response{}
		rows.Scan(&team.TeamId, &team.FirstName, &team.LastName, &team.Email, &team.Address, &team.MobileNo, &team.Status, &team.Joining_Date)
		teams = append(teams, team)
	}
	return teams, Shared.ErrorMsg{Message: ""}
}

func LoginTeam(w http.ResponseWriter, r *http.Request) (Team, Shared.ErrorMsg) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var client = Team{}

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT ("FirstName"), ("LastName"), ("Email"), ("Password"), ("Address"), ("MobileNo"), ("Status") FROM slh_teams WHERE ("Email")=$1;`
	team := Team{}
	row := config.Db.QueryRow(sqlStatement, client.Email)
	err = row.Scan(&team.FirstName, &team.LastName, &team.Email, &team.Password, &team.Address, &team.MobileNo, &team.Status)

	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		return team, Shared.ErrorMsg{Message: "Email or Password Incorrect."}
	case nil:
		eror := bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(client.Password))
		if eror != nil {
			w.WriteHeader(http.StatusForbidden)
			return team, Shared.ErrorMsg{Message: "Bad Credentials"}
		}

		expirationTime := time.Now().Add(50 * time.Minute)
		claims := &Shared.Claims{
			Username: team.Email,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		fmt.Println("Owaish")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(Shared.JwtKey)

		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return team, Shared.ErrorMsg{Message: "Internal Server Error."}
		}

		ps := &team
		ps.Token = tokenString

		sqlStatement := `UPDATE slh_teams SET "Token"=$1 WHERE "Email"=$2`

		_, err = config.Db.Exec(sqlStatement, tokenString, team.Email)
		if err != nil {
			return team, Shared.ErrorMsg{Message: "Internal Server Error."}
		}
		team.Password = ""
		return team, Shared.ErrorMsg{Message: ""}
	default:
		panic(err)
	}
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) (Team, Shared.ErrorMsg) {

	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	var team = Team{}

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}

	if team.FirstName == "" || team.LastName == "" || team.Email == "" || team.Password == "" || team.Address == "" || team.MobileNo == "" || team.Status == "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return team, Shared.ErrorMsg{Message: "BLANK FIELDS"}
	}

	userEmail := context.Get(r, middleware.Decoded)

	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Email" = $3, "Password" = $4, "Address" = $5, "MobileNo" = $6, "Status" = $7 WHERE ("Email") = $8`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Email, team.Password, team.Address, team.MobileNo, team.Status, userEmail)
	if err != nil {
		panic(err)
	}
	team.Password = ""
	return team, Shared.ErrorMsg{Message: ""}
}
