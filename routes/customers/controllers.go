package customers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
	"strconv"
	"time"
)

func BuildRespons(response *CustomerUpd, customer Customer) CustomerUpd {
	response.Customer_Id = customer.Customer_Id
	response.Customer_Name = customer.Customer_Name
	response.Customer_Gender = customer.Customer_Gender
	response.Customer_Email = customer.Customer_Email
	response.Customer_Address = customer.Customer_Address
	response.Customer_Mobile_No = customer.Customer_Mobile_No
	response.Status = customer.Status
	response.Registered_Source = customer.Registered_Source
	response.Pincode = customer.Pincode
	return *response
}

func BuildResp(response *Resp, customer CustomerTran) Resp {
	response.Merchant_Transaction_Id = customer.Merchant_Transaction_Id
	response.Payment_Gateway_Id = customer.Payment_Gateway_Id
	response.Payment_Gateway_Mode = customer.Payment_Gateway_Mode
	response.Bank_Type = customer.Bank_Type
	response.Transaction_Mode = customer.Transaction_Mode
	return *response
}

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
		w.WriteHeader(http.StatusPreconditionFailed)
		return customer, res
	}

	customer.Status = true
	customer.CreatedAt = time.Now().Local()
	customer.Last_Access_Time = time.Now().Local()
	customer.Registered_Source = "website"

	sqlStatement := `
	INSERT INTO slh_customers ("Customer_Name", "Customer_Email", "Customer_Address", "Pincode", "Customer_Gender","Customer_Mobile_No", 
	"CreatedAt", "Last_Access_Time", "Status", "Registered_Source")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING "Customer_Id"`

	customer.Customer_Id = 0
	err = config.Db.QueryRow(sqlStatement, customer.Customer_Name, customer.Customer_Email, customer.Customer_Address, customer.Pincode,
		customer.Customer_Gender, customer.Customer_Mobile_No, customer.CreatedAt, customer.Last_Access_Time, customer.Status,
		customer.Registered_Source).Scan(&customer.Customer_Id)
	if err != nil {

		sqlStatement := ` UPDATE slh_customers SET "Customer_Name" = $1, "Customer_Email" = $2, "Customer_Address" = $3, "Pincode" = $4,
		 "Customer_Gender" = $5, "Last_Access_Time" = $6 WHERE ("Customer_Mobile_No") = $7`

		_, err = config.Db.Exec(sqlStatement, customer.Customer_Name, customer.Customer_Email, customer.Customer_Address, customer.Pincode,
			customer.Customer_Gender, customer.Last_Access_Time, customer.Customer_Mobile_No)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return customer, ErrorMsg{Message: "Internal Server Error"}
		}

		sqlStatements := `SELECT ("Customer_Id"), ("Customer_Souls_Id"), ("CreatedAt") FROM slh_customers WHERE ("Customer_Mobile_No")=$1`
		row := config.Db.QueryRow(sqlStatements, customer.Customer_Mobile_No)
		err = row.Scan(&customer.Customer_Id, &customer.Customer_Souls_Id, &customer.CreatedAt)
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
			return customer, ErrorMsg{Message: "Internal Server Error"}
		}
		fmt.Println(32)
		return customer, ErrorMsg{}

	}

	customer.Customer_Souls_Id = customer.CreatedAt.Format("20060102") + strconv.Itoa(customer.Customer_Id)
	sqlStatement = `UPDATE slh_customers SET "Customer_Souls_Id" = $1 WHERE "Customer_Id" =  $2`
	_, err = config.Db.Exec(sqlStatement, customer.Customer_Souls_Id, customer.Customer_Id)
	if err != nil {

		return Customer{}, ErrorMsg{Message: "Internal Server Error."}
	}
	return customer, ErrorMsg{}
}

func CustomerBooking(w http.ResponseWriter, r *http.Request) (Tran, ErorMsg) {
	r.ParseForm()

	tran := Tran{}
	customer := CustomerOrder{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}

	var res ErorMsg
	res.Message = ""
	CheckEmpty(customer, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return Tran{}, res
	}

	sqlStatements := `SELECT ("Customer_Name"), ("Customer_Souls_Id"), ("Pincode"), ("Customer_Address") FROM slh_customers WHERE ("Customer_Id")=$1;`
	row := config.Db.QueryRow(sqlStatements, customer.Customer_Id)
	err = row.Scan(&customer.Customer_Name, &customer.Customer_Souls_Id, &customer.Pincode, &customer.Customer_Address)
	fmt.Println(customer.Customer_Id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return Tran{}, ErorMsg{Message: "Unauthorised User"}
	}

	customer.CreatedAt = time.Now().Local()
	// Input or Provide ??
	customer.Slot_Date = time.Now().Local()
	customer.Slot_Time = time.Now().Local()

	// string??
	customer.Latitude = "qwerty"
	customer.Longitude = "asdfg"

	customer.Is_Order_Confirmed = false
	customer.Total_Order_Amount = 10000

	cur_time := time.Now().Unix()
	tranId := strconv.FormatInt(int64(cur_time), 10) + "-" + strconv.Itoa(customer.Customer_Id)
	customer.Merchant_Transaction_Id = tranId
	//	fmt.Println(customer.Customer_Transaction_Id)

	sqlStatement := `
	INSERT INTO slh_customers_pending_orders ("Customer_Id", "Customer_Souls_Id", "Number_Of_Therapists_Required", "Therapist_Gender", 
	"Massage_For","Slot_Time", "Slot_Date", "Customer_Address", "Pincode", "Latitude", "Longitude", "Is_Order_Confirmed", "Merchant_Transaction_Id", 
	"Massage_Duration", "CreatedAt", "Customer_Name", "Total_Order_Amount") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)  
	RETURNING ("Customer_Order_Id")`

	customer.Customer_Order_Id = 0
	err = config.Db.QueryRow(sqlStatement, customer.Customer_Id, customer.Customer_Souls_Id, customer.Number_Of_Therapist, customer.Therapist_Gender,
		customer.Massage_For, customer.Slot_Time, customer.Slot_Date, customer.Customer_Address, customer.Pincode, customer.Latitude, customer.Longitude,
		customer.Is_Order_Confirmed, customer.Merchant_Transaction_Id, customer.Massage_Duration, customer.CreatedAt, customer.Customer_Name,
		customer.Total_Order_Amount).Scan(&customer.Customer_Order_Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return Tran{}, ErorMsg{Message: "Internal Server Error"}
	}
	tran.Merchant_Transaction_Id = customer.Merchant_Transaction_Id
	return tran, ErorMsg{}
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

func ViewCustomer(w http.ResponseWriter, r *http.Request) (Customer, ErrorMessage) {
	r.ParseForm()
	customer := Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}
	sqlStatement := `SELECT ("Customer_Email"),("CreatedAt"),("Customer_Gender"),("Customer_Name"),("Customer_Address"), ("Registered_Source"),
	("Last_Access_Time"), ("Customer_Souls_Id"), ("Customer_Mobile_No"), ("Status"), ("Pincode") FROM slh_customers WHERE ("Customer_Id")=$1;`
	row := config.Db.QueryRow(sqlStatement, customer.Customer_Id)
	err = row.Scan(&customer.Customer_Email, &customer.CreatedAt, &customer.Customer_Gender, &customer.Customer_Name, &customer.Customer_Address,
		&customer.Registered_Source, &customer.Last_Access_Time, &customer.Customer_Souls_Id, &customer.Customer_Mobile_No, &customer.Status, &customer.Pincode)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return Customer{}, ErrorMessage{Message: "Internal Server Error"}
	}
	return customer, ErrorMessage{}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) (CustomerUpd, ErrorMessage) {
	r.ParseForm()
	customer := Customer{}
	response := CustomerUpd{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}

	var res Result

	sqlStatement := ` UPDATE slh_customers SET "Customer_Name" = $1, "Customer_Email" = $2, "Customer_Address" = $3, "Pincode" = $4, "Customer_Gender" = $5,
	"Status" = $6, "Registered_Source" = $7, "Customer_Mobile_No" = $8 WHERE ("Customer_Id") = $9`

	res, err = config.Db.Exec(sqlStatement, customer.Customer_Name, customer.Customer_Email, customer.Customer_Address, customer.Pincode,
		customer.Customer_Gender, customer.Status, customer.Registered_Source, customer.Customer_Mobile_No, customer.Customer_Id)
	if err != nil {
		fmt.Println(22)
		w.WriteHeader(http.StatusNotFound)
		return CustomerUpd{}, ErrorMessage{Message: "Mobile_Number already registered"}
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		w.WriteHeader(http.StatusNotFound)
		return CustomerUpd{}, ErrorMessage{Message: "Unauthorised User"}
	}
	BuildRespons(&response, customer)
	return response, ErrorMessage{}
}

func ListCustomer(w http.ResponseWriter, r *http.Request) ([]Customer, ErrorMessage) {
	r.ParseForm()
	// fmt.Print("ASHS")
	var response []Customer
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

	mobile := r.Form.Get("mobile")
	if mobile != "" {
		if err := Shared.ParseInt(r.Form.Get("mobile_no"), &q.Customer_Mobile_No); err != nil {
			return response, ErrorMessage{Message: "parseerr"}
		}
		q.Customer_Mobile_No = q.Customer_Mobile_No - 1
	} else {
		q.Customer_Mobile_No = 0
	}

	q.Customer_Souls_Id = r.Form.Get("customer_souls_id")
	q.Customer_Name = r.Form.Get("name")
	q.Customer_Order_Id = r.Form.Get("order_id")
	// q.Customer_Mobile_No = r.Form.Get("mobile_no")
	q.Status = r.Form.Get("status")

	fmt.Println(q)
	offset := q.Limit * q.Page
	// fmt.Print("ASHS")
	var customers []Customer
	sqlStatement := `SELECT ("Customer_Souls_Id"),("Customer_Name"),("Customer_Mobile_No"),("Status") FROM slh_customers WHERE ("Customer_Souls_Id") LIKE ''||$1||'%' AND ("Customer_Name") LIKE ''|| $2 ||'%' ORDER BY ("CreatedAt") DESC LIMIT $3 OFFSET $4`
	rows, err := config.Db.Query(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Limit, offset)

	if err != nil {
		fmt.Print("asfafs")
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: "Internal Server Error."}
	}
	// fmt.Print("ASHS")
	// fmt.Println(len(rows))
	for rows.Next() {
		var customer = Customer{}
		rows.Scan(&customer.Customer_Souls_Id, &customer.Customer_Name, &customer.Customer_Mobile_No, &customer.Status)
		customers = append(customers, customer)
		// cnt = cnt + 1
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_customers WHERE ("Customer_Souls_Id") LIKE ''||$1||'%' AND  ("Customer_Name") LIKE ''|| $2 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.Customer_Souls_Id, q.Customer_Name)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
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

func ViewCustomerBooking(w http.ResponseWriter, r *http.Request) (CustomerOrder, ErrorMessage) {

	r.ParseForm()
	customer := CustomerOrder{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}

	sqlStatement := `
	SELECT ("Customer_Id"), ("Customer_Souls_Id"), ("Number_Of_Therapists_Required"), ("Therapist_Gender"), 
	("Massage_For"),("Slot_Time"), ("Slot_Date"), ("Customer_Address"), ("Pincode"), ("Latitude"), ("Longitude"), ("Is_Order_Confirmed"), 
	("Merchant_Transaction_Id"), ("Massage_Duration"), ("CreatedAt"), ("Customer_Name"), ("Total_Order_Amount") FROM slh_customers_pending_orders 
	WHERE ("Customer_Order_Id") =$1 ;`

	row := config.Db.QueryRow(sqlStatement, customer.Customer_Order_Id)
	err = row.Scan(&customer.Customer_Id, &customer.Customer_Souls_Id, &customer.Number_Of_Therapist, &customer.Therapist_Gender, &customer.Massage_For,
		&customer.Slot_Time, &customer.Slot_Date, &customer.Customer_Address, &customer.Pincode,
		&customer.Latitude, &customer.Longitude, &customer.Is_Order_Confirmed, &customer.Merchant_Transaction_Id, &customer.Massage_Duration,
		&customer.CreatedAt, &customer.Customer_Name, &customer.Total_Order_Amount)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return CustomerOrder{}, ErrorMessage{Message: "Internal Server Error"}
	}
	return customer, ErrorMessage{}
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

func UpdateCustomerTransaction(w http.ResponseWriter, r *http.Request) (Resp, ErorMesg) {
	r.ParseForm()
	customer := CustomerTran{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}
	response := Resp{}

	var res Result

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
