package main

import (
	"fmt"
	"log"
	"net/http"
	"soul_api/middleware"
	"soul_api/routes/customers"
	"soul_api/routes/partners"
	"soul_api/routes/pendingOrders"
	"soul_api/routes/team"
	"soul_api/routes/teamHasRole"

	"soul_api/routes/transactions"

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
	r.HandleFunc("/team/update-member/password", middleware.ValidateTokenMiddleware(team.UpdatePassword)).Methods("PUT")
	r.HandleFunc("/team/update-team-member", middleware.ValidateTokenMiddleware(team.UpdateMember)).Methods("PUT")
	r.HandleFunc("/team/view-member", middleware.ValidateTokenMiddleware(team.View)).Methods("GET")
	r.HandleFunc("/team/update-status", middleware.ValidateTokenMiddleware(team.UpdateStatus)).Methods("POST")
	r.HandleFunc("/team/logout", middleware.ValidateTokenMiddleware(team.Logout)).Methods("GET")

	r.HandleFunc("/team/role", middleware.ValidateTokenMiddleware(teamHasRole.Role)).Methods("POST")
	r.HandleFunc("/team/has-role/list", middleware.ValidateTokenMiddleware(teamHasRole.HasRole)).Methods("GET")
	r.HandleFunc("/team/has-role/update", middleware.ValidateTokenMiddleware(teamHasRole.HasRoleUpdate)).Methods("PUT")

	r.HandleFunc("/customers/registration", customers.Create).Methods("POST")
	r.HandleFunc("/customers/update", middleware.ValidateTokenMiddleware(customers.Update)).Methods("PUT")
	r.HandleFunc("/customers/view", middleware.ValidateTokenMiddleware(customers.View)).Methods("GET")
	r.HandleFunc("/customers/list", middleware.ValidateTokenMiddleware(customers.List)).Methods("GET")

	r.HandleFunc("/customers/booking", pendingOrders.Create).Methods("POST")
	r.HandleFunc("/customers/booking/view", middleware.ValidateTokenMiddleware(pendingOrders.View)).Methods("GET")
	r.HandleFunc("/customers/booking/list", middleware.ValidateTokenMiddleware(pendingOrders.List)).Methods("GET")

	r.HandleFunc("/customers/transaction", transactions.Create).Methods("POST")
	r.HandleFunc("/customers/transaction/view", middleware.ValidateTokenMiddleware(transactions.View)).Methods("GET")
	r.HandleFunc("/customers/transaction/list", middleware.ValidateTokenMiddleware(transactions.List)).Methods("GET")
	r.HandleFunc("/customers/transaction/update", middleware.ValidateTokenMiddleware(transactions.Update)).Methods("PUT")

	r.HandleFunc("/partner/register", middleware.ValidateTokenMiddleware(partners.Create)).Methods("POST")
	r.HandleFunc("/partner/update", middleware.ValidateTokenMiddleware(partners.Update)).Methods("PUT")
	r.HandleFunc("/partner/List", middleware.ValidateTokenMiddleware(partners.List)).Methods("GET")

	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Println("Server Started")
}
