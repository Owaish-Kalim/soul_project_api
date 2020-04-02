package routes

import (
	"fmt"
	"net/http"
	"soul_api/config"
	"encoding/json"
	"database/sql"
	"time" 
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"soul_api/middleware"
)

type Team struct {
	TeamId int `json:"teamid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json: "password", Db:"password"`
	Address string `json:"address"`	
	Token 	string `json:"token"`
	MobileNo string	`json:"mobileno"`
	Status string `json:"status"`
	Joining_Date  time.Time 
	CreatedAt time.Time
}

func CreateTeam(w http.ResponseWriter, r *http.Request) (Team, ErrorMsg) {
	r.ParseForm()
	team := Team{}
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil { panic(err)}

	if team.FirstName == "" || team.LastName == "" || team.Email == "" || team.Password == "" || team.Address == "" || team.MobileNo == "" || team.Status == "" {
		// http.Error(w, http.StatusText(400), http.StatusBadRequest) 
		w.WriteHeader(http.StatusBadRequest) 
		return team, ErrorMsg{Message: "Fields cannot be empty."}
	}

	team.Joining_Date = time.Now().Local()
	team.CreatedAt = time.Now().Local()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return team,  ErrorMsg{Message: "Internal Server Error."}
	} 

	sqlStatement := `
	INSERT INTO slh_teams ("FirstName","LastName","Email", "Address", "JoiningDate", "CreatedAt", "Password", "MobileNo", "Status")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING ("TeamId")`

	team.TeamId = 0 
	err = config.Db.QueryRow(sqlStatement, team.FirstName, team.LastName, team.Email,team.Address, team.Joining_Date, team.CreatedAt, string(hashedPassword), team.MobileNo, team.Status).Scan(&team.TeamId)
	if err != nil {
		// return team, err
		return team,  ErrorMsg{Message: "Internal Server Error."}
	} 

	team.Password = ""
	return team, ErrorMsg{Message: ""}
}


func LoginTeam(w http.ResponseWriter, r *http.Request) (Team, error) {
	fmt.Println(1)
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	var client=Team{};

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		panic(err)
	}
	fmt.Println(2)
	sqlStatement := `SELECT ("FirstName"), ("LastName"), ("Email"), ("Password"), ("Address"), ("MobileNo"), ("Status") FROM slh_teams WHERE ("Email")=$1;`
	fmt.Println(3)
	
	team := Team{} ;
	row := config.Db.QueryRow(sqlStatement, client.Email)
	err = row.Scan(&team.FirstName, &team.LastName, &team.Email, &team.Password, &team.Address, &team.MobileNo, &team.Status) 
	fmt.Println(4)
	switch err {
	case sql.ErrNoRows:
		return team, sql.ErrNoRows
	case nil: 
		
		// hsPwd,bErr:= bcrypt.GenerateFromPassword([]byte(client.Password), 8)
		// fmt.Println(hsPwd)
		// if bErr != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	} 
		
		eror := bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(client.Password))
		if eror != nil {
			return team, eror
		} 

		expirationTime := time.Now().Add(15 * time.Minute)
		claims := &Claims{
			Username: team.Email, 
			StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			},
		}
		fmt.Println("Owaish")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return team, err
		}

		ps := &team
		ps.Token = tokenString
		
		sqlStatement := `UPDATE slh_teams SET "Token"=$1 WHERE "Email"=$2`

		_, err = config.Db.Exec(sqlStatement, tokenString, team.Email)
		if err != nil {
			return team, err
		}
		team.Password = ""
		return team, nil

	default:
	panic(err)
	}
} 


func UpdateTeam(w http.ResponseWriter, r *http.Request) (Team, error) { 

	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	var team = Team{};
	
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}


	if team.FirstName == "" || team.LastName == "" || team.Email == "" || team.Password == "" || team.Address == "" || team.MobileNo == "" || team.Status == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)  
		return team, err
	}

	userEmail := context.Get(r,  middleware.Decoded)
	
	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Email" = $3, "Password" = $4, "Address" = $5, "MobileNo" = $6, "Status" = $7 WHERE ("Email") = $8`

	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Email, team.Password, team.Address, team.MobileNo, team.Status, userEmail) 
	if err != nil {
  	panic(err)
	} 
	team.Password = "" 
	return team,nil
}
