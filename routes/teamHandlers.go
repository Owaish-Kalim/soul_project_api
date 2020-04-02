package routes

import (
	"fmt"
	"encoding/json"
	"net/http"

)

func TeamCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Hello Teamsss")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	team, err := CreateTeam(w, r)
	fmt.Println(team)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}


func TeamLogin(w http.ResponseWriter, r *http.Request) { 
	fmt.Println(r.Method)
	if r.Method != "POST" { 
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("HERE")
	usr, err := LoginTeam(w, r)
	fmt.Println("LOL")
//	fmt.Println(usr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	 json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usr)
}


func TeamUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(1)
	usr, err := UpdateTeam(w, r)
	fmt.Println(2)
	if err != nil {
		fmt.Println("ow")
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	 json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usr)
}

