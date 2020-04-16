package users

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
	"soul_api/routes"
)


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
	err = row.Scan(&user.Name, &user.Email, &user.Password)

	switch err {
	case sql.ErrNoRows:
		return user, sql.ErrNoRows
	case nil: 
		
		// hsPwd,bErr:= bcrypt.GenerateFromPassword([]byte(client.Password), 8)
		// fmt.Println(hsPwd)
		// if bErr != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	} 
		
		eror := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(client.Password))
		if eror != nil {
			return user, eror
		} 

		expirationTime := time.Now().Add(15 * time.Minute)
		claims := &Shared.Claims{
			Username: user.Email,
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
			return user, err
		}

		ps := &user
		ps.Token = tokenString
		
		sqlStatement := `UPDATE users SET "Token"=$1 WHERE "Email"=$2`

		_, err = config.Db.Exec(sqlStatement, tokenString, user.Email)
		if err != nil {
			return user, err
		}

		return user, nil

	default:
	panic(err)
	}
} 



func UpdateUser(w http.ResponseWriter, r *http.Request) (User,error) { 

	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	var user = User{};
	
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	userEmail := context.Get(r,  middleware.Decoded)
	
	sqlStatement := ` UPDATE users SET "Name" = $1, "Email" = $2, "Password" = $3 WHERE ("Email") = $4`

	_, err = config.Db.Exec(sqlStatement, user.Name, user.Email, user.Password, userEmail) 
	if err != nil {
  	panic(err)
	}
	return user,nil
}

// func DeleteUser(w http.ResponseWriter, r *http.Request) (User, error) {
	
// 	// userEmail := context.Get(r,  middleware.Decoded)

// 	// sqlStatement := ` DELETE FROM users WHERE Email = $1;`
// 	// _, err := config.Db.Exec(sqlStatement, userEmail) 
// 	// if err != nil {
//   	// panic(err)
// 	// }
// } 


