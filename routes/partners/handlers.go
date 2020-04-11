package partners

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ppopop")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	partner, err := CreatePartner(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(partner)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	partner, err := UpdatePartner(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(partner)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	partner, err := ListPartner(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(partner)
}
