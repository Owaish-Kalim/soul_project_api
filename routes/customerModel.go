package routes

import (
	"fmt"
	"net/http"
	"soul_api/config"
	"encoding/json"
	"time" 
)

type Customer struct {
	Customer_Soul_Id string `json:"customersoulid"` 
	Customer_Name string `json:"name"`
	Mobile_No string `json:"mobile"`
	Customer_Gender string `json:"gender"`
	Pincode string `json:"pincode"`
	Customer_Email string `json:"email"`
	Customer_Address string `json:"address"` 
	CreatedAt time.Time
}

func CreateCustomers(w http.ResponseWriter, r *http.Request) (Customer, error) {
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	customer := Customer{}

	err := json.NewDecoder(r.Body).Decode(&customer) 
	if err != nil {
		panic(err)
	}

	customer.CreatedAt = time.Now().Local()

	sqlStatement := `
	INSERT INTO slh_customers ("Customer_Soul_Id","Customer_Name","Customer_Email", "Customer_Address", "Pincode", "Customer_Gender","Mobile_No", "CreatedAt")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING "Customer_Id"`
 
	Customer_Id := 1
	err = config.Db.QueryRow(sqlStatement, customer.Customer_Soul_Id, customer.Customer_Name, customer.Customer_Email, customer.Customer_Address, customer.Pincode , customer.Customer_Gender, customer.Mobile_No, customer.CreatedAt).Scan(&Customer_Id)
	if err != nil {
		fmt.Println(err)
		return customer, err
	}
	return customer, err
}

