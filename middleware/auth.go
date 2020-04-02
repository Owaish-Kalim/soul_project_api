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
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
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
					fmt.Println("asdas") ; 
					sqlStatement := `SELECT ("Email") FROM slh_teams WHERE ("Token")=$1;`
					fmt.Println(1) ;
					row := config.Db.QueryRow(sqlStatement, bearerToken[1])
					err := row.Scan(&email) 
					if err != nil {
					log.Fatal(err)	 
					} 
					mapstructure.Decode(token.Claims, &email)
					context.Set(req,Decoded,email) 
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
