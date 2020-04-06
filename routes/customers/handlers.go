package customers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Customers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Hellozzz")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	customer, err := AddCustomer(w, r)
	// fmt.Println(team)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Print("AAAA")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func Booking(w http.ResponseWriter, r *http.Request) {
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
