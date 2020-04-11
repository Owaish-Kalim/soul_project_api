package pendingOrders

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
	"strconv"
	"time"
)

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

func ListCustomerBooking(w http.ResponseWriter, r *http.Request) ([]CustomerOrder, ErrorMessage) {
	r.ParseForm()
	var response []CustomerOrder
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
	// q.Customer_Mobile_No = r.Form.Get("mobile_no")
	// q.Status = r.Form.Get("status")
	// q. = r.Form.Get("customer_souls_id")
	// q.Customer_Name = r.Form.Get("name")
	// q.Customer_Order_Id = r.Form.Get("order_id")
	// q.Customer_Mobile_No = r.Form.Get("mobile_no")
	// q.Status = r.Form.Get("status")
	q.Customer_Email = r.Form.Get("customer_email")
	q.Customer_Gender = r.Form.Get("customer_gender")

	// fmt.Println(q)
	offset := q.Limit * q.Page
	// fmt.Print("ASHS")
	var customers []CustomerOrder
	sqlStatement := `SELECT ("Customer_Order_Id"), ("Customer_Souls_Id"),("Customer_Name"), ("Customer_Id"),("Pincode"),("Customer_Address"),
	("Massage_Duration") ,("Number_Of_Therapists_Required"),("Massage_For"), ("Therapist_Gender"), ("Merchant_Transaction_Id"),
	("Total_Order_Amount"), ("Latitude"),("Longitude"), ("Is_Order_Confirmed"), ("CreatedAt"), ("Slot_Date"), ("Slot_Time") FROM slh_customers_pending_orders
	WHERE ("Customer_Souls_Id") LIKE  ''||$1||'%' 
	AND ("Customer_Name") LIKE ''|| $2 ||'%' 
	AND ("Customer_Address") LIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") LIKE ''|| $4 ||'%' 
	AND ("Massage_For") LIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") LIKE ''|| $6 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $7 OFFSET $8`

	rows, err := config.Db.Query(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Customer_Address, q.Massage_Duration, q.Massage_For,
		q.Merchant_Transaction_Id, q.Limit, offset)

	if err != nil {
		fmt.Print("asfafs")
		panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: "Internal Server Error."}
	}
	// fmt.Print("ASHS")
	// fmt.Println(len(rows))
	for rows.Next() {
		var customer = CustomerOrder{}
		rows.Scan(&customer.Customer_Order_Id, &customer.Customer_Souls_Id, &customer.Customer_Name,
			&customer.Customer_Id, &customer.Pincode, &customer.Customer_Address, &customer.Massage_Duration, &customer.Number_Of_Therapist,
			&customer.Massage_For, &customer.Therapist_Gender, &customer.Merchant_Transaction_Id, &customer.Total_Order_Amount, &customer.Latitude,
			&customer.Longitude, &customer.Is_Order_Confirmed, &customer.CreatedAt, &customer.Slot_Date, &customer.Slot_Time)
		customers = append(customers, customer)
		// cnt = cnt + 1
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_customers_pending_orders WHERE 
	("Customer_Souls_Id") LIKE  ''||$1||'%' 
	AND ("Customer_Name") LIKE ''|| $2 ||'%' 
	AND ("Customer_Address") LIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") LIKE ''|| $4 ||'%' 
	AND ("Massage_For") LIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") LIKE ''|| $6 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Customer_Address, q.Massage_Duration,
		q.Massage_For, q.Merchant_Transaction_Id)
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