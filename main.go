package main

import (
	"fmt"
	"log"
	"net/http"
	"soul_api/middleware"
	"soul_api/routes/customers"
	"soul_api/routes/team"
	"soul_api/routes/users"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)

	r.HandleFunc("/api/users", users.Create).Methods("POST")
	r.HandleFunc("/api/users/show", middleware.ValidateTokenMiddleware(users.List)).Methods("GET")
	r.HandleFunc("/api/users/update", middleware.ValidateTokenMiddleware(users.Update)).Methods("PUT")
	r.HandleFunc("/api/users/delete", middleware.ValidateTokenMiddleware(users.Delete)).Methods("DELETE")
	r.HandleFunc("/api/users/login", users.Login).Methods("POST")

	r.HandleFunc("/team/login", team.Login).Methods("POST")
	r.HandleFunc("/team/list", middleware.ValidateTokenMiddleware(team.List)).Methods("GET")
	r.HandleFunc("/team/add-member", team.Create).Methods("POST")
	r.HandleFunc("/team/update-member", middleware.ValidateTokenMiddleware(team.Update)).Methods("PUT")
	r.HandleFunc("/team/update-team-member", middleware.ValidateTokenMiddleware(team.UpdateMember)).Methods("POST")
	r.HandleFunc("/team/view-member", middleware.ValidateTokenMiddleware(team.View)).Methods("GET")
	// r.HandleFunc(`/team/edit-team-member/{:id}`, middleware.ValidateTokenMiddleware(team.View)).Methods("POST")
	r.HandleFunc("/team/update-status", middleware.ValidateTokenMiddleware(team.UpdateStatus)).Methods("POST")
	r.HandleFunc("/api/customers", customers.CrtCustomers).Methods("POST")

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Println("Server Started")
}
