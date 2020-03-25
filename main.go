package main


import (
	"net/http"
	"soul_api/routes"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
)


// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
//     (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//     (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func main() {
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST","PUT","DELETE"})
	origins := handlers.AllowedMethods([]string{"*"})

	r.HandleFunc("/api/user/me", routes.Show).Methods("GET")
	r.HandleFunc("/api/users", routes.Create).Methods("POST")
	// r.Run()
	log.Fatal(http.ListenAndServe(":8930", handlers.CORS(headers, methods, origins)(r)))
}