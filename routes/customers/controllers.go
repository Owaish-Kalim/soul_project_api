package customers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	"strconv"
	"time"
)

func AddCustomer(w http.ResponseWriter, r *http.Request) (Customer, ErrorMsg) {
	r.ParseForm()
	customer := Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}

	var res ErrorMsg
	res.Message = ""
	CheckEmptyList(customer, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return customer, res
	}

	customer.CreatedAt = time.Now().Local()
	customer.Last_Access_Time = time.Now().Local()

	sqlStatement := `
	INSERT INTO slh_customers ("Customer_Name", "Customer_Email", "Customer_Address", "Pincode", "Customer_Gender","Mobile_No", "CreatedAt", "Last_Access_Time", "Status", "Registered_Source")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING "Customer_Id"`

	customer.Customer_Id = 0
	err = config.Db.QueryRow(sqlStatement, customer.Customer_Name, customer.Customer_Email, customer.Customer_Address, customer.Pincode, customer.Customer_Gender, customer.Mobile_No, customer.CreatedAt, customer.Last_Access_Time, customer.Status, customer.Registered_Source).Scan(&customer.Customer_Id)
	if err != nil {

		// sqlStatement := ` UPDATE slh_customers SET "Customer_Name" = $1, "Customer_Email" = $2, "Customer_Address" = $3, "Pincode" = $4, "Customer_Gender" = $5, "Last_Access_Time" = $6 WHERE ("Mobile_No") = $7`

		// _, err = config.Db.Exec(sqlStatement, customer.Customer_Name, customer.Customer_Email, customer.Customer_Address, customer.Pincode, customer.Customer_Gender, customer.Last_Access_Time, customer.Mobile_No)
		// if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return customer, ErrorMsg{Message: "Mobile No. already registered"}
	}
	// Print only req. values
	// return customer, ErrorMsg{Message: ""}
	// w.WriteHeader(http.StatusPreconditionFailed)
	// return Customer{}, ErrorMsg{Message: "Mobile_No. already registered"}

	customer.Customer_Soul_Id = customer.CreatedAt.Format("060102") + strconv.Itoa(customer.Customer_Id)
	sqlStatement = `UPDATE slh_customers SET "Customer_Soul_Id" = $1 WHERE "Customer_Id" =  $2`
	_, err = config.Db.Exec(sqlStatement, customer.Customer_Soul_Id, customer.Customer_Id)
	if err != nil {

		return Customer{}, ErrorMsg{Message: "Internal Server Error."}
	}
	return customer, ErrorMsg{}
}

func CustomerBooking(w http.ResponseWriter, r *http.Request) (CustomerOrder, ErorMsg) {
	r.ParseForm()

	customer := CustomerOrder{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}

	var res ErorMsg
	res.Message = ""
	CheckEmpty(customer, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return CustomerOrder{}, res
	}

	sqlStatements := `SELECT ("Customer_Id"), ("Customer_Soul_Id"), ("Pincode"), ("Customer_Address") FROM slh_customers WHERE ("Mobile_No")=$1;`
	row := config.Db.QueryRow(sqlStatements, customer.Mobile_No)
	err = row.Scan(&customer.Customer_Id, &customer.Customer_Soul_Id, &customer.Pincode, &customer.Address)
	fmt.Println(customer.Customer_Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return CustomerOrder{}, ErorMsg{Message: "Unauthorised User"}
	}

	customer.CreatedAt = time.Now().Local()
	customer.Slot_Date = time.Now().Local()
	customer.Slot_Time = time.Now().Local()

	cur_time := time.Now().Unix()
	tranId := strconv.FormatInt(int64(cur_time), 10) + strconv.Itoa(customer.Customer_Id)
	customer.Transaction_Id = tranId
	fmt.Println(customer.Transaction_Id)

	sqlStatement := `
	INSERT INTO slh_customers_pending_orders ("Customer_Id", "Customer_Soul_Id", "Number_Of_Therapists_Required", "Therapist_Gender", 
	"Massage_For","Slot_Time", "Slot_Date", "Address", "Pincode", "Latitude", "Longitude", "Is_Order_Confirmed", "Transaction_Id", 
	"Massage_Duration", "CreatedAt", "Mobile_No")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
	RETURNING ("Customer_Primary_Order_Id")`

	customer.Customer_Primary_Order_Id = 0
	err = config.Db.QueryRow(sqlStatement, customer.Customer_Id, customer.Customer_Soul_Id, customer.Num_Therapist, customer.Therapist_Gender, customer.Massage_For,
		customer.Slot_Time, customer.Slot_Date, customer.Address, customer.Pincode, customer.Latitude, customer.Longitude, customer.Is_Order_Confirmed,
		customer.Transaction_Id, customer.Massage_Duration, customer.CreatedAt, customer.Mobile_No).Scan(&customer.Customer_Primary_Order_Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return customer, ErorMsg{Message: "Internal Server Error"}
	}

	return customer, ErorMsg{}
}
