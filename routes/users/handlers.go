package users

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/context"
	"soul_api/middleware"
	"soul_api/config"
	// "soul_api/routes"
)

func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usr, err := CreateUser(w, r)
	fmt.Println(usr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	 json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usr)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { 
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("HERE")
	usr, err := LoginUser(w, r)
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


 var code string

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method != "GET" { 
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	
 	userEmail := context.Get(r, middleware.Decoded) 
	fmt.Println(userEmail)

}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	
	usr, err := UpdateUser(w, r)
	fmt.Println(usr)
	if err != nil {
		fmt.Println("ow")
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	 json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usr)
}


func Delete(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	if r.Method != "DELETE" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	
	userEmail := context.Get(r,  middleware.Decoded)

	sqlStatement := ` DELETE FROM users WHERE "Email" = $1;`
	_, err := config.Db.Exec(sqlStatement, userEmail) 
	if err != nil {
  	panic(err)
	}

	// usr, err := DeleteUser(w, r)
	// fmt.Println(usr)
	// if err != nil {
	// 	fmt.Println("ow")
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	//  json.NewEncoder(w).Encode(err)
	// }

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(usr)
}
