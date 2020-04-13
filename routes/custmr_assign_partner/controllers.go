package custmr_assign_partner

import (
	"encoding/json"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
	"time"
)

func CheckEmpty(customerpartner CustomerPartner, res *Shared.ErrorMesg) {

	if customerpartner.Commission_Amount == 0 {
		res.Commission_Amount = "Commission Amount cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Created_By == "" {
		res.Created_By = "Created By   cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Updated_By == "" {
		res.Updated_By = "Updated By  cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Customer_Souls_Id == "" {
		res.Customer_Souls_Id = "Customer Souls Id cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Status == "" {
		res.Status = "Status cannot be empty."
		res.Message = "Error"
	}

}

func CreateCustomerPartner(w http.ResponseWriter, r *http.Request) (CustomerPartner, Shared.ErrorMesg) {
	r.ParseForm()
	customerpartner := CustomerPartner{}
	//response := Response{}
	err := json.NewDecoder(r.Body).Decode(&customerpartner)
	if err != nil {
		panic(err)
	}

	var res Shared.ErrorMesg
	res.Message = ""
	CheckEmpty(customerpartner, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return customerpartner, res
	}

	customerpartner.Slot_Date = time.Now().Local()
	customerpartner.CreatedAt = time.Now().Local()
	customerpartner.Slot_Time = time.Now().Local()

	sqlStatements := `SELECT ("Customer_Name"),("Customer_Id"),("Merchant_Transaction_Id") 
	FROM slh_customers_pending_orders WHERE ("Customer_Souls_Id")=$1;`
	row := config.Db.QueryRow(sqlStatements, customerpartner.Merchant_Transaction_Id)
	err = row.Scan(&customerpartner.Customer_Name, &customerpartner.Customer_Id, &customerpartner.Merchant_Transaction_Id)
	//fmt.Println(customer.Customer_Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res.Message = "Unauthorized User"
		return customerpartner, res
	}

	sqlStatement := `
	INSERT INTO slh_assign_customer_with_partner ("Customer_Souls_Id ","Customer_Name", "Customer_Id",
	"Merchant_Transaction_Id","Status","Commission_Amount", 
	"Created_By","Updated_By","CreatedAt", "Slot_Date", "Slot_Time")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	//team.TeamId = 0
	_, err = config.Db.Exec(sqlStatement, customerpartner.Customer_Souls_Id, customerpartner.Customer_Name,
		customerpartner.Customer_Id, customerpartner.Merchant_Transaction_Id, customerpartner.Status,
		customerpartner.Commission_Amount, customerpartner.Created_By, customerpartner.Updated_By,
		customerpartner.CreatedAt, customerpartner.Slot_Date, customerpartner.Slot_Time)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Internal Server Error"
		return customerpartner, res
	}
	//BuildResponse(&response, team)
	res.Message = ""
	return customerpartner, res
}

// func UpdateCustomerPartner(w http.ResponseWriter, r *http.Request) (CustomerPartner, Shared.ErrorMsg) {
// 	w.Header().Set("Content-Type", "application/json")
// 	r.ParseForm()
// 	var customerpartner = CustomerPartner{}
// 	err := json.NewDecoder(r.Body).Decode(&customerpartner)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var res Shared.ErrorMsg
// 	res.Message = ""

// 	CheckEmpty(customerpartner, &res)
// 	if res.Message != "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return customerpartner, res
// 	}

// 	userEmail := context.Get(r, middleware.Decoded)

// 	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Email" = $3, "Password" = $4, "Address" = $5, "MobileNo" = $6, "Status" = $7 WHERE ("Email") = $8`

// 	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Email, team.Password, team.Address, team.MobileNo, team.Status, userEmail)
// 	if err != nil {
// 		return response, Shared.ErrorMsg{Message: "Email already registered"}
// 		//panic(err)
// 	}
// 	BuildUpdateResponse(&response, team)
// 	return response, Shared.ErrorMsg{Message: ""}
// }
