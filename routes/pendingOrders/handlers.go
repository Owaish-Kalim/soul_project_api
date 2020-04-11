package pendingOrders

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	customer, err := CustomerBooking(w, r)
	// fmt.Println(team)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}

	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func View(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	customer, err := ViewCustomerBooking(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1234)
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	customer, err := ListCustomerBooking(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}
