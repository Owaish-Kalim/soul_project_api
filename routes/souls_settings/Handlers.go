package souls_settings

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	temp, err := CreateSettings(w, r)
	fmt.Println(temp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	temp, err := UpdateSettings(w, r)
	// fmt.Println(temp)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(temp)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	temp, err := ListSettings(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(temp)
}
