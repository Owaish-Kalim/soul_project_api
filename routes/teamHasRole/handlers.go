package teamHasRole

import (
	"encoding/json"
	"net/http"
)

func HasRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	team, err := TeamHasRole(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(team)
}

func HasRoleUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	team, err := TeamHasRoleUpdate(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(team)
}

func Role(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	team, err := Role_Team(w, r)
	if err.Message != "" {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(team)
}
