package partner

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func Create(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("In")
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
	fmt.Println("out")
	json.NewEncoder(w).Encode(partner)
}