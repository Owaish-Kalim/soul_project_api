package main


import (
	"net/http"
	"soul_api/routes"
	//"database/sql"
	"github.com/gorilla/mux"
    "github.com/rs/cors"
	"log"
	"fmt"
	// "github.com/mitchellh/mapstructure"
	// "strings"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/gorilla/context"
	// "encoding/json"
	// "soul_api/config"
	"soul_api/middleware"
)




func main() {
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)
	r.HandleFunc("/api/users", routes.Create).Methods("POST")
	r.HandleFunc("/api/users/show", middleware.ValidateTokenMiddleware(routes.List)).Methods("GET")

	r.HandleFunc("/api/users/login", routes.Login).Methods("POST") 
	r.HandleFunc("/api/teams", routes.CrtTeam).Methods("POST")
	r.HandleFunc("/api/customers", routes.CrtCustomers).Methods("POST")
	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Println("Server Started")
}

