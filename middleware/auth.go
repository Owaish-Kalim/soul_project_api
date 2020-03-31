package middleware

import (
	"net/http"
	"log"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"encoding/json"
	"soul_api/config"
)

type ErrorMsg struct {
	Message string `json:"message"`
}

var Decoded string

func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	//fmt.Println("asfaf")
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		
		// fmt.Println("asfkhagf")
		
		// fmt.Println(req.Method)
		authorizationHeader := req.Header.Get("Authorization")
		//	fmt.Println("LOL")
		//fmt.Println(authorizationHeader)
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			//fmt.Println(bearerToken)
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("my_secret_key"), nil
					
				})
				
				if error != nil {
					json.NewEncoder(w).Encode(ErrorMsg{Message: error.Error()})
					return
				}

				if token.Valid {

					var email string ;
					fmt.Println(bearerToken[1])
					sqlStatement := `SELECT ("Email") FROM users WHERE ("Token")=$1;`
					row := config.Db.QueryRow(sqlStatement, bearerToken[1])
					err := row.Scan(&email) 
					fmt.Println(email)
					if err != nil {
					log.Fatal(err)	 
					} 
					// fmt.Println(email)
					// fmt.Println(email)
					// fmt.Println(req)
					//var user routes.User
					mapstructure.Decode(token.Claims, &email)
					// req.user = user
					// fmt.Println(user)
					// fmt.Println("Ayushi")

					// vars := mux.Vars(req)
					// name := vars["userId"]
					// if name != user.Username {
					// 	json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
					// 	return
					// } 	
									
					context.Set(req,Decoded, email)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(ErrorMsg{Message: "An authorization header is required"})
		}
	})
}
