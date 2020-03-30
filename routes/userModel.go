package routes

import (
	"fmt"
	"net/http"
	"soul_api/config"
	"encoding/json"
	"database/sql"
	"time" 
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id 	int
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// var err error

func CreateUser(w http.ResponseWriter, r *http.Request) (User, error) {

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	sqlStatement := `
	INSERT INTO users ("Name","Email","Password", "CreatedAt","UpdatedAt")
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	id := 0
	err = config.Db.QueryRow(sqlStatement, user.Name, user.Email, string(hashedPassword), user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return user, err
	}

	return user, err
} 


func LoginUser(w http.ResponseWriter, r *http.Request) (User, error) {

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	var client=User{};

	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		panic(err)
	}

	sqlStatement := `SELECT ("Name"), ("Email"), ("Password") FROM users WHERE ("Email")=$1;`
	
	var user User;
	row := config.Db.QueryRow(sqlStatement, client.Email)
	err = row.Scan(&user.Name, &user.Email,
	&user.Password)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("BAD CREDENTIALS!")
		return user, sql.ErrNoRows
	case nil: 
		fmt.Println("GHUSGAYA")
		fmt.Println(client.Password)
		hsPwd,bErr:= bcrypt.GenerateFromPassword([]byte(client.Password), 8)
		if bErr != nil {
			// fmt.Println("kalim") 
			w.WriteHeader(http.StatusInternalServerError)
			} 
		fmt.Println(string(hsPwd))
		fmt.Println(user.Password)
		eror := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(client.Password))
		if eror != nil {
			return user, eror
		}
		return user, nil
	default:
	panic(err)
	}
} 



func UpdateUser(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `
	UPDATE users
	SET name = $2, email = $3, password = &4
	WHERE Email = $1;`

	_, err := config.Db.Exec(sqlStatement, 1, "name", "email", "password") 
	if err != nil {
		fmt.Println("ERROR: ")
		fmt.Println(err)
  	panic(err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	sqlStatement := `
	DELETE FROM users
	WHERE Email = $1;`
	_, err := config.Db.Exec(sqlStatement, 1)
	if err != nil {
  	panic(err)
	}
} 

