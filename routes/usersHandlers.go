package routes

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

	http.StatusText(200)

	usr, err := CreateUser(w, r)
	fmt.Println(usr)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	userJson, err := json.Marshal(usr)
	if err != nil {
		panic(err)
	}
	fmt.Println(userJson)

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usr)
}
