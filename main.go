package main


import (
	"net/http"
	"soul_api/routes"
	"github.com/gorilla/mux"
    "github.com/rs/cors"
	"log"
	"fmt"
	"soul_api/middleware"
)

func main() {
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)
	r.HandleFunc("/api/users", routes.Create).Methods("POST")
	r.HandleFunc("/api/users/show", middleware.ValidateTokenMiddleware(routes.List)).Methods("GET")
	r.HandleFunc("/api/users/update", middleware.ValidateTokenMiddleware(routes.Update)).Methods("PUT")
	r.HandleFunc("/api/users/delete", middleware.ValidateTokenMiddleware(routes.Delete)).Methods("DELETE")
	r.HandleFunc("/api/users/login", routes.Login).Methods("POST") 
	r.HandleFunc("/team/add-member", routes.TeamCreate).Methods("POST")
	r.HandleFunc("/team/login", routes.TeamLogin).Methods("POST") 
	r.HandleFunc("/team/update-member", middleware.ValidateTokenMiddleware(routes.TeamUpdate)).Methods("PUT")
	r.HandleFunc("/api/customers", routes.CrtCustomers).Methods("POST")
	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Println("Server Started")
}

