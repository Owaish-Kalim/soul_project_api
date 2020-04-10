package customers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Customers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	customer, err := AddCustomer(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}

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

func Transaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	customer, err := CustomerTransaction(w, r)
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
	customer, err := ViewCustomer(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func ViewBooking(w http.ResponseWriter, r *http.Request) {
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

func ListBooking(w http.ResponseWriter, r *http.Request) {
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

func ViewTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	customer, err := ViewCustomerTransaction(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func ListTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	customer, err := ListCustomerTransaction(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	customer, err := UpdateCustomer(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	customer, err := UpdateCustomerTransaction(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	customer, err := ListCustomer(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}
