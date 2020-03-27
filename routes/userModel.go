package routes

import (
	"fmt"
	"net/http"
	"soul_api/config"
	"encoding/json"
	"time" 
)

type User struct {
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

	// if user.Name == "" || user.Password == "" || user.Email == "" {
	// 	http.Error(w, http.StatusText(400), 400)
	// 	return user, nil
	// }

	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()

	sqlStatement := `
	INSERT INTO users ("Name","Email","Password", "CreatedAt","UpdatedAt")
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	id := 0
	err = config.Db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&id)
	if err != nil {
		return user, err
	}

	return user, err
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

