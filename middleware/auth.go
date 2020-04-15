package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"soul_api/config"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

type ErrorMsg struct {
	Message string `json:"message"`
}

var Decoded string

func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		//	fmt.Println(authorizationHeader)
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("my_secret_key"), nil
				})
				//	fmt.Println("asf")
				if error != nil {
					json.NewEncoder(w).Encode(ErrorMsg{Message: error.Error()})
					fmt.Println(ErrorMsg{Message: error.Error()})
					return
				}
				//	fmt.Println("HEHR")
				if token.Valid {
					var email string
					//	fmt.Println("asdas")
					sqlStatement := `SELECT ("Email") FROM slh_teams WHERE ("Token")=$1;`
					//fmt.Println(1)
					row := config.Db.QueryRow(sqlStatement, bearerToken[1])
					//fmt.Println("HEasfkhaHR")
					err := row.Scan(&email)

					if err != nil {
						log.Fatal(err)
					}
					//	fmt.Println("HEasfkhaHR")
					mapstructure.Decode(token.Claims, &email)
					context.Set(req, Decoded, email)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(ErrorMsg{Message: err.Error()})
				}
			} else {
				json.NewEncoder(w).Encode(ErrorMsg{Message: err.Error()})
			}
		} else {
			json.NewEncoder(w).Encode(ErrorMsg{Message: err.Error()})
		}
	})
}
