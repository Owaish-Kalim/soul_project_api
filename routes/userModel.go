package routes

import (
	"fmt"
	"net/http"
	// "soul_api/config"
	"encoding/json"
	"time"
)

type User struct {
	name     string
	password string
	email    string
	CreatedAt time.Time
}

var err error

func CreateUser(w http.ResponseWriter, r *http.Request) (User, error) {

	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	user.CreatedAt = time.Now().Local()


	// user.name = r.FormValue("name")
	// user.email = r.FormValue("email")
	// user.password = r.FormValue("password")

	// if user.name == "" || user.password == "" || user.email == "" {
	// 	fmt.Println("BLANKFIELDS")
	// 	http.Error(w, http.StatusText(400), 400)
	// 	return user, nil
	// }
	fmt.Println(r.FormValue("name"))
	// _, err = config.DB.Exec(`Insert INTO users (name, email, password) VALUES ($1, $2, $3)`, user.name, user.email, user.password)

	// if err != nil {
	// 	http.Error(w, http.StatusText(500), 500)
	// 	return user, nil
	// }
	return user, err
}

func ShowUser() (User, error) {
	var usr User
	return usr, nil
}
