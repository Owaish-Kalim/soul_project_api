package main


import (
	"net/http"
	"soul_api/routes"
	"github.com/gorilla/mux"
    "github.com/rs/cors"
	"log"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)
	r.HandleFunc("/api/users", routes.Create).Methods("POST")
	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Println("Server Started")
}

