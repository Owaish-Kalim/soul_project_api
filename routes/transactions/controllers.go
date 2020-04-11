package transactions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
	"strconv"
	// "time"
)

func BuildResp(response *Resp, customer CustomerTran) Resp {
	response.Merchant_Transaction_Id = customer.Merchant_Transaction_Id
	response.Payment_Gateway_Id = customer.Payment_Gateway_Id
	response.Payment_Gateway_Mode = customer.Payment_Gateway_Mode
	response.Bank_Type = customer.Bank_Type
	response.Transaction_Mode = customer.Transaction_Mode
	return *response
}

func CustomerTransaction(w http.ResponseWriter, r *http.Request) (CustomerTran, ErorMesg) {
	r.ParseForm()

	customer := CustomerTran{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}

	var res ErorMesg
	res.Message = ""
	CheckEmptyTran(customer, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return customer, res
	}

	sqlStatements := `SELECT ("Customer_Order_Id"), ("Customer_Id"), ("Customer_Souls_Id"), ("Number_Of_Therapists_Required"), ("Therapist_Gender"), 
	("Massage_For"),("Slot_Time"), ("Slot_Date"), ("Customer_Address"), ("Pincode"), ("Latitude"), ("Longitude"), 
	("Massage_Duration"), ("CreatedAt"), ("Customer_Name"), ("Total_Order_Amount") FROM slh_customers_pending_orders WHERE ("Merchant_Transaction_Id")=$1;`
	row := config.Db.QueryRow(sqlStatements, customer.Merchant_Transaction_Id)
	err = row.Scan(&customer.Customer_Order_Id, &customer.Customer_Id, &customer.Customer_Souls_Id, &customer.Number_Of_Therapist, &customer.Therapist_Gender,
		&customer.Massage_For, &customer.Slot_Time, &customer.Slot_Date, &customer.Customer_Address, &customer.Pincode, &customer.Latitude, &customer.Longitude,
		&customer.Massage_Duration, &customer.CreatedAt, &customer.Customer_Name, &customer.Total_Order_Amount)
	fmt.Println(customer.Customer_Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return CustomerTran{}, ErorMesg{Message: "Unauthorised User"}
	}
	fmt.Println(12345678)
	sqlStatement := `
	INSERT INTO slh_transactions ("Customer_Order_Id","Customer_Id", "Customer_Souls_Id", "Number_Of_Therapist_Required", "Therapist_Gender", 
	"Massage_For","Slot_Time", "Slot_Date", "Customer_Address", "Pincode", "Latitude", "Longitude", "Merchant_Transaction_Id", 
	"Massage_Duration", "CreatedAt", "Customer_Name", "Total_Order_Amount", "Payment_Gateway_Id", "Payment_Gateway_Mode", "Transaction_Mode", "Bank_Type") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)`

	_, err = config.Db.Exec(sqlStatement, customer.Customer_Order_Id, customer.Customer_Id, customer.Customer_Souls_Id, customer.Number_Of_Therapist, customer.Therapist_Gender,
		customer.Massage_For, customer.Slot_Time, customer.Slot_Date, customer.Customer_Address, customer.Pincode, customer.Latitude, customer.Longitude,
		customer.Merchant_Transaction_Id, customer.Massage_Duration, customer.CreatedAt, customer.Customer_Name, customer.Total_Order_Amount,
		customer.Payment_Gateway_Id, customer.Payment_Gateway_Mode, customer.Transaction_Mode, customer.Bank_Type)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//panic(err)
		return customer, ErorMesg{Message: "Internal Server Error"}
	}

	return customer, ErorMesg{}
}

func ViewCustomerTransaction(w http.ResponseWriter, r *http.Request) (CustomerTran, ErrorMessage) {
	r.ParseForm()

	customer := CustomerTran{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}

	sqlStatements := `SELECT ("Customer_Order_Id"), ("Customer_Id"), ("Customer_Souls_Id"), ("Number_Of_Therapist_Required"), ("Therapist_Gender"), 
	("Massage_For"),("Slot_Time"), ("Slot_Date"), ("Customer_Address"), ("Pincode"), ("Latitude"), ("Longitude"), 
	("Massage_Duration"), ("CreatedAt"), ("Customer_Name"), ("Total_Order_Amount"), ("Payment_Gateway_Id"), ("Payment_Gateway_Mode"), ("Transaction_Mode"),
	("Bank_Type") FROM slh_transactions WHERE ("Merchant_Transaction_Id")=$1;`
	row := config.Db.QueryRow(sqlStatements, customer.Merchant_Transaction_Id)
	err = row.Scan(&customer.Customer_Order_Id, &customer.Customer_Id, &customer.Customer_Souls_Id, &customer.Number_Of_Therapist, &customer.Therapist_Gender,
		&customer.Massage_For, &customer.Slot_Time, &customer.Slot_Date, &customer.Customer_Address, &customer.Pincode, &customer.Latitude, &customer.Longitude,
		&customer.Massage_Duration, &customer.CreatedAt, &customer.Customer_Name, &customer.Total_Order_Amount, &customer.Payment_Gateway_Id,
		&customer.Payment_Gateway_Mode, &customer.Transaction_Mode, &customer.Bank_Type)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return CustomerTran{}, ErrorMessage{Message: "Internal Server Error"}
	}
	return customer, ErrorMessage{}
}

func ListCustomerTransaction(w http.ResponseWriter, r *http.Request) ([]CustomerTran, ErrorMessage) {
	r.ParseForm()
	var response []CustomerTran
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			return response, ErrorMessage{Message: "parseerr"}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			return response, ErrorMessage{Message: "parseerr"}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}

	mobile := r.Form.Get("mobile_no")
	if mobile != "" {
		if err := Shared.ParseInt(r.Form.Get("mobile_no"), &q.Customer_Mobile_No); err != nil {
			return response, ErrorMessage{Message: "parseerr"}
		}
		q.Customer_Mobile_No = q.Customer_Mobile_No
	} else {
		q.Customer_Mobile_No = 0
	}

	q.Customer_Souls_Id = r.Form.Get("customer_souls_id")
	q.Customer_Name = r.Form.Get("customer_name")
	q.Customer_Order_Id = r.Form.Get("customer_id")
	q.Customer_Address = r.Form.Get("customer_address")
	q.Massage_Duration = r.Form.Get("massage_duration")
	q.Massage_For = r.Form.Get("massage_for")
	q.Merchant_Transaction_Id = r.Form.Get("merchant_transaction_id")
	q.Payment_Gateway_Id = r.Form.Get("payment_gateway_id")
	q.Payment_Gateway_Mode = r.Form.Get("payment_gateway_mode")
	q.Transaction_Mode = r.Form.Get("transaction_mode")
	q.Bank_Type = r.Form.Get("bank_type")

	// q.Customer_Mobile_No = r.Form.Get("mobile_no")
	// q.Status = r.Form.Get("status")
	// q. = r.Form.Get("customer_souls_id")
	// q.Customer_Name = r.Form.Get("name")
	// q.Customer_Order_Id = r.Form.Get("order_id")
	// q.Customer_Mobile_No = r.Form.Get("mobile_no")
	// q.Status = r.Form.Get("status")
	// fmt.Println(q)
	offset := q.Limit * q.Page
	var customers []CustomerTran
	sqlStatement := `SELECT ("Customer_Order_Id"), ("Customer_Souls_Id"),("Customer_Name"), ("Customer_Id"),("Pincode"),("Customer_Address"),
	("Massage_Duration") ,("Number_Of_Therapist_Required"),("Massage_For"), ("Therapist_Gender"), ("Merchant_Transaction_Id"),
	("Total_Order_Amount"), ("Latitude"),("Longitude"), ("CreatedAt"), ("Slot_Date"), ("Slot_Time"), ("Payment_Gateway_Id"), ("Payment_Gateway_Mode"), 
	("Transaction_Mode"), ("Bank_Type")  FROM slh_transactions
	WHERE ("Customer_Souls_Id") LIKE  ''||$1||'%' 
	AND ("Customer_Name") LIKE ''|| $2 ||'%' 
	AND ("Customer_Address") LIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") LIKE ''|| $4 ||'%' 
	AND ("Massage_For") LIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") LIKE ''|| $6 ||'%' 
	AND ("Payment_Gateway_Id") LIKE ''|| $7 ||'%' 
	AND ("Payment_Gateway_Mode") LIKE ''|| $8 ||'%' 
	AND ("Transaction_Mode") LIKE ''|| $9 ||'%' 
	AND ("Bank_Type") LIKE ''|| $10 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $11 OFFSET $12`

	rows, err := config.Db.Query(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Customer_Address, q.Massage_Duration, q.Massage_For,
		q.Merchant_Transaction_Id, q.Payment_Gateway_Id, q.Payment_Gateway_Mode, q.Transaction_Mode, q.Bank_Type, q.Limit, offset)

	if err != nil {
		fmt.Print("asfafs")
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: "Internal Server Error."}
	}

	for rows.Next() {
		var customer = CustomerTran{}
		rows.Scan(&customer.Customer_Order_Id, &customer.Customer_Souls_Id, &customer.Customer_Name,
			&customer.Customer_Id, &customer.Pincode, &customer.Customer_Address, &customer.Massage_Duration, &customer.Number_Of_Therapist,
			&customer.Massage_For, &customer.Therapist_Gender, &customer.Merchant_Transaction_Id, &customer.Total_Order_Amount, &customer.Latitude,
			&customer.Longitude, &customer.CreatedAt, &customer.Slot_Date, &customer.Slot_Time, &customer.Payment_Gateway_Id,
			&customer.Payment_Gateway_Mode, &customer.Transaction_Mode, &customer.Bank_Type)
		customers = append(customers, customer)
		// cnt = cnt + 1
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_transactions WHERE 
	("Customer_Souls_Id") LIKE  ''||$1||'%' 
	AND ("Customer_Name") LIKE ''|| $2 ||'%' 
	AND ("Customer_Address") LIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") LIKE ''|| $4 ||'%' 
	AND ("Massage_For") LIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") LIKE ''|| $6 ||'%'
	AND ("Payment_Gateway_Id") LIKE ''|| $7 ||'%' 
	AND ("Payment_Gateway_Mode") LIKE ''|| $8 ||'%' 
	AND ("Transaction_Mode") LIKE ''|| $9 ||'%' 
	AND ("Bank_Type") LIKE ''|| $10 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Customer_Address, q.Massage_Duration,
		q.Massage_For, q.Merchant_Transaction_Id, q.Payment_Gateway_Id, q.Payment_Gateway_Mode, q.Transaction_Mode, q.Bank_Type)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		fmt.Println(232)
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: "Internal Server Error."}
	}

	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt / q.Limit
	if cnt%q.Limit != 0 {
		totalPages = totalPages + 1
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	fmt.Println(cnt)
	return customers, ErrorMessage{Message: ""}
}

func UpdateCustomerTransaction(w http.ResponseWriter, r *http.Request) (Resp, ErorMesg) {
	fmt.Println(12211)
	r.ParseForm()
	customer := CustomerTran{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}
	response := Resp{}

	var res Shared.Result

	sqlStatement := ` UPDATE slh_transactions SET "Payment_Gateway_Id" = $1, "Bank_Type" = $2, "Payment_Gateway_Mode" = $3, "Transaction_Mode" = $4
	WHERE ("Merchant_Transaction_Id") = $5`

	res, err = config.Db.Exec(sqlStatement, customer.Payment_Gateway_Id, customer.Bank_Type, customer.Payment_Gateway_Mode, customer.Transaction_Mode,
		customer.Merchant_Transaction_Id)
	if err != nil {
		fmt.Println(22)
		w.WriteHeader(http.StatusInternalServerError)
		return Resp{}, ErorMesg{Message: "Internal Server Error"}
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		w.WriteHeader(http.StatusNotFound)
		return Resp{}, ErorMesg{Message: "Unauthorised User"}
	}
	BuildResp(&response, customer)
	return response, ErorMesg{}
}
