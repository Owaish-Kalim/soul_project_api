package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
//     (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//     (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func Show(w http.ResponseWriter, r *http.Request) {
	// setupResponse(&w, r)
	// if (*r).Method == "OPTIONS" {
	// 	return
	// }

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("sadasf")

	usr, err := ShowUser()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}

func Create(w http.ResponseWriter, r *http.Request) {
	// setupResponse(&w, r)

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	http.StatusText(200)

	usr, err := CreateUser(w, r)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	json.NewEncoder(w).Encode(usr)
}
