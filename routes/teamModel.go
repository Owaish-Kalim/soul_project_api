package routes

import (
	"fmt"
	"net/http"
	"soul_api/config"
	"encoding/json"
	"time" 
	"golang.org/x/crypto/bcrypt"
)

type Team struct {
	Team_Id int 
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
	Email     string `json:"email"`
	Password string `json: "password", Db:"password"`
	Address string `json:"address"`
	
	Joining_Date  time.Time 
	CreatedAt time.Time

	// UpdatedAt time.Time
}

func CreateTeam(w http.ResponseWriter, r *http.Request) (Team, error) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	team := Team{}

	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		panic(err)
	}
	team.Joining_Date = time.Now().Local()
	team.CreatedAt = time.Now().Local()
	// team.UpdatedAt = time.Now().Local()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(team.Password), 8)
	
	// fmt.Println(string(hashedPassword))
	if err != nil {
	// fmt.Println("kalim") 
	w.WriteHeader(http.StatusInternalServerError)
	} 

	sqlStatement := `
	INSERT INTO slh_teams ("FirstName","LastName","Email", "Address", "JoiningDate", "CreatedAt", "Password")
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING ("TeamId")`

	Team_Id := 0
	err = config.Db.QueryRow(sqlStatement, team.First_Name, team.Last_Name, team.Email,team.Address, team.Joining_Date, team.CreatedAt, string(hashedPassword)).Scan(&Team_Id)
	if err != nil {
		fmt.Println(err)
		return team, err
	}
	return team, err
}

