package routes

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func CrtCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hellozzz")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	customer, err := CreateCustomers(w, r)
	// fmt.Println(team)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
