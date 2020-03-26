package main


import (
	"net/http"
	"soul_api/routes"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST","PUT","DELETE"})
	origins := handlers.AllowedMethods([]string{"*"})

	r.HandleFunc("/api/user/me", routes.Show).Methods("GET")
	r.HandleFunc("/api/users", routes.Create).Methods("POST")
	// r.Run()
	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8930", handlers.CORS(headers, methods, origins)(r)))
	fmt.Println("Server Started")
}