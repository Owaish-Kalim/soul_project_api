package main


import (
	"net/http"
	"soul_api/routes"
	"github.com/gorilla/mux"
    "github.com/rs/cors"
	"log"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"encoding/json"
)


type ErrorMsg struct {
	Message string `json:"message"`
}


func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("asfaf")
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		
		fmt.Println("asfkhagf")
		
		fmt.Println(req.Method)
		authorizationHeader := req.Header.Get("Authorization")
		fmt.Println("LOL")
		fmt.Println(authorizationHeader)
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			fmt.Println("LOL")
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
				fmt.Println("LOL")
				if token.Valid {

					var user routes.User
					mapstructure.Decode(token.Claims, &user)
					// req.user = user
					fmt.Println(user)


					// vars := mux.Vars(req)
					// name := vars["userId"]
					// if name != user.Username {
					// 	json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
					// 	return
					// }

					context.Set(req, "decoded", token.Claims)
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



func main() {
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)
	r.HandleFunc("/api/users", routes.Create).Methods("POST")
	r.HandleFunc("/api/users/show", validateTokenMiddleware(routes.List)).Methods("GET")

	r.HandleFunc("/api/users/login", routes.Login).Methods("POST") 
	r.HandleFunc("/api/teams", routes.CrtTeam).Methods("POST")
	r.HandleFunc("/api/customers", routes.CrtCustomers).Methods("POST")
	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Println("Server Started")
}

