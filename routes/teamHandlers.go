package routes

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func CrtTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello Teamsss")
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	team, err := CreateTeam(w, r)
	fmt.Println(team)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}
