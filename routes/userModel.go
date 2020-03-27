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
}

var err error

func CreateUser(w http.ResponseWriter, r *http.Request) (User, error) {

	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	user.CreatedAt = time.Now().Local()

	if user.Name == "" || user.Password == "" || user.Email == "" {
		http.Error(w, http.StatusText(400), 400)
		return user, nil
	}

	sqlStatement := `
	INSERT INTO users (name,email,password)
	VALUES ($1, $2, $3)
	RETURNING id`

	id := 0
	err = config.DB.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
	panic(err)
	}

	fmt.Println("New record ID is:", id)

	return user, err
} 

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	

	sqlStatement := `
	UPDATE users
	SET name = $2, email = $3, password = &4
	WHERE id = $1;`

	_, err = config.DB.Exec(sqlStatement, 1, "name", "email", "password") 
	if err != nil {
  	panic(err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	sqlStatement := `
	DELETE FROM users
	WHERE id = $1;`
	_, err = config.DB.Exec(sqlStatement, 1)
	if err != nil {
  	panic(err)
	}
} 

