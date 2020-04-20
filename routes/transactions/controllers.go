package transactions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"soul_api/config"
	"soul_api/email"
	Shared "soul_api/routes"
	"strconv"

	// "time"
	"math"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var Updated = false

var SocketResponseData SocketResponse

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {

	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func Assign(MTId string) CustomerPartner {

	cust_part := CustomerPartner{}

	//Select Partner
	temp := Temp{}
	var Number_Of_Therapists string
	sqlStatement := `SELECT ("Latitude"), ("Longitude"), ("Number_Of_Therapist_Required") FROM slh_transactions  WHERE ("Merchant_Transaction_Id")=$1;`
	row := config.Db.QueryRow(sqlStatement, MTId)
	err := row.Scan(&temp.Cust_Lat, &temp.Cust_Long, &Number_Of_Therapists)
	if err != nil {
		fmt.Print("HOLSS")
		panic(err)
	}
	Lat1, _ := strconv.ParseFloat(temp.Cust_Lat, 8)
	Long1, _ := strconv.ParseFloat(temp.Cust_Long, 8)
	Number_Of_Therapist, _ := strconv.Atoi(Number_Of_Therapists)

	sqlStatement = `SELECT ("Latitude"), ("Longitude"), ("Partner_Souls_Id") FROM slh_partners  WHERE 1=1 ;`
	rows, err := config.Db.Query(sqlStatement)
	if err != nil {
		fmt.Print("asfafFFs")
		panic(err)
	}

	// MinDist := math.MaxFloat64
	var partners []Partner

	for rows.Next() {
		temp := Temp{}
		rows.Scan(&temp.Part_Lat, &temp.Part_Long, &temp.Part_Souls_Id)
		Lat2, _ := strconv.ParseFloat(temp.Part_Lat, 8)
		Long2, _ := strconv.ParseFloat(temp.Part_Long, 8)

		dist := Distance(Lat1, Long1, Lat2, Long2)

		partner := Partner{}
		partner.Souls_Id = temp.Part_Souls_Id
		partner.Dis = dist
		partners = append(partners, partner)
	}

	sort.SliceStable(partners, func(i, j int) bool {
		return partners[i].Dis < partners[j].Dis
	})

	for i := 0; i < len(partners); i++ {

		sqlStatement = `SELECT ("Partner_Name"), ("Partner_Email") FROM slh_partners  WHERE ("Partner_Souls_Id")=$1; `
		row = config.Db.QueryRow(sqlStatement, partners[i].Souls_Id)
		err = row.Scan(&cust_part.Partner_Name, &cust_part.Partner_Email)
		if err != nil {
			fmt.Print("LOLA")
			panic(err)
		}

		email.SendEmail(cust_part.Partner_Name, cust_part.Partner_Email)
		Number_Of_Therapist--

		// if Number_Of_Therapist == 0 {
		// 	 func check(&Num_Of_Therapist)
		// }

		// if Number_Of_Therapist == 0 {
		// 	 return
		// }

	}

	// sqlStatement = `SELECT ("Partner_Name"), ("Partner_Mobile_No") FROM slh_partners  WHERE ("Partner_Souls_Id")=$1; `
	// row = config.Db.QueryRow(sqlStatement, cust_part.Partner_Souls_Id)
	// err = row.Scan(&cust_part.Partner_Name, &cust_part.Partner_Mobile_No)
	// if err != nil {
	// 	fmt.Print("KOLA")
	// 	panic(err)
	// }

	cust_part.Slot_Date = "14-04-2020"
	cust_part.Slot_Time = "1 PM"
	cust_part.Status = "Pending"

	return CustomerPartner{}

}

func Socket(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	var conn, _ = upgrader.Upgrade(w, r, nil)
	fmt.Println("SOCKET START")
	go func(conn *websocket.Conn) {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
			}
		}
	}(conn)
	fmt.Println("SOCKET MID")
	go func(conn *websocket.Conn) {
		ch := time.Tick(5 * time.Second)
		fmt.Println("DOESIT")
		for range ch {
			if Updated != false {
				conn.WriteJSON(SocketResponseData)
				Updated = false
			}

		}
	}(conn)

	fmt.Println("SOCKET END")

	// userEmail := context.Get(r, middleware.Decoded)

	// fmt.Println(userEmail)

}

func CustomerTransaction(w http.ResponseWriter, r *http.Request) (CustomerTran, ErorMesg) {

	r.ParseForm()

	customer := CustomerTran{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Println("HERE")
		panic(err)
	}

	var res ErorMesg
	res.Message = ""
	CheckEmptyTran(customer, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusPreconditionFailed)
		return customer, res
	}

	fmt.Println(12)

	sqlStatements := `SELECT ("Customer_Order_Id"), ("Customer_Id"), ("Customer_Souls_Id"), ("Number_Of_Therapists_Required"), ("Therapist_Gender"), 
	("Massage_For"),("Slot_Time"), ("Slot_Date"), ("Customer_Address"), ("Pincode"), ("Latitude"), ("Longitude"), 
	("Massage_Duration"), ("CreatedAt"), ("Customer_Name"), ("Total_Order_Amount") FROM slh_customers_pending_orders WHERE ("Merchant_Transaction_Id")=$1;`
	row := config.Db.QueryRow(sqlStatements, customer.Merchant_Transaction_Id)
	err = row.Scan(&customer.Customer_Order_Id, &customer.Customer_Id, &customer.Customer_Souls_Id, &customer.Number_Of_Therapist, &customer.Therapist_Gender,
		&customer.Massage_For, &customer.Slot_Time, &customer.Slot_Date, &customer.Customer_Address, &customer.Pincode, &customer.Latitude, &customer.Longitude,
		&customer.Massage_Duration, &customer.CreatedAt, &customer.Customer_Name, &customer.Total_Order_Amount)
	fmt.Println(customer.Customer_Id)
	if err != nil {
		fmt.Println(12)
		w.WriteHeader(http.StatusNotFound)
		return CustomerTran{}, ErorMesg{Message: err.Error()}
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
		fmt.Println(12)
		return customer, ErorMesg{Message: err.Error()}
	}

	// NOTIFICATION

	Updated = true
	SocketResponseData.Customer_Name = customer.Customer_Name
	SocketResponseData.Customer_Souls_Id = customer.Customer_Souls_Id
	SocketResponseData.Merchant_Transaction_Id = customer.Merchant_Transaction_Id

	//Customer_Assign_Partner
	Assign(customer.Merchant_Transaction_Id)

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
		return CustomerTran{}, ErrorMessage{Message: err.Error()}
	}
	return customer, ErrorMessage{}
}

func ListCustomerTransaction(w http.ResponseWriter, r *http.Request) ([]CustomerTran, ErrorMessage) {
	// fmt.Println(1266)
	r.ParseForm()
	// var response []CustomerTran
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []CustomerTran{}, ErrorMessage{Message: err.Error()}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []CustomerTran{}, ErrorMessage{Message: err.Error()}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}

	fmt.Println(12)

	q.Customer_Souls_Id = r.Form.Get("customer_souls_id")
	q.Customer_Name = r.Form.Get("customer_name")
	q.Total_Order_Amount = r.Form.Get("total_order_amount")
	q.Pincode = r.Form.Get("pincode")
	q.Massage_Duration = r.Form.Get("massage_duration")
	q.Merchant_Transaction_Id = r.Form.Get("merchant_transaction_id")
	q.Payment_Gateway_Id = r.Form.Get("payment_gateway_id")
	q.Payment_Gateway_Mode = r.Form.Get("payment_gateway_mode")
	q.Transaction_Mode = r.Form.Get("transaction_mode")
	q.Bank_Type = r.Form.Get("bank_type")
	q.Slot_Date = r.Form.Get("slot_date")
	q.Slot_Time = r.Form.Get("slot_time")
	q.CreatedAt = r.Form.Get("created_at")

	offset := q.Limit * q.Page
	var customers []CustomerTran
	sqlStatement := `SELECT ("Customer_Order_Id"), ("Customer_Souls_Id"),("Customer_Name"), ("Customer_Id"),("Pincode"),("Customer_Address"),
	("Massage_Duration") ,("Number_Of_Therapist_Required"),("Massage_For"), ("Therapist_Gender"), ("Merchant_Transaction_Id"),
	("Total_Order_Amount"), ("Latitude"),("Longitude"), ("CreatedAt"), ("Slot_Date"), ("Slot_Time"), ("Payment_Gateway_Id"), ("Payment_Gateway_Mode"), 
	("Transaction_Mode"), ("Bank_Type")  FROM slh_transactions
	WHERE ("Customer_Souls_Id") ILIKE  ''||$1||'%' 
	AND ("Customer_Name") ILIKE ''|| $2 ||'%' 
	AND ("Total_Order_Amount") ILIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") ILIKE ''|| $4 ||'%' 
	AND ("Pincode") ILIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") ILIKE ''|| $6 ||'%' 
	AND ("Payment_Gateway_Id") ILIKE ''|| $7 ||'%' 
	AND ("Payment_Gateway_Mode") ILIKE ''|| $8 ||'%' 
	AND ("Transaction_Mode") ILIKE ''|| $9 ||'%' 
	AND ("Bank_Type") ILIKE ''|| $10 ||'%' 
	AND ("Slot_Date") ILIKE ''|| $11||'%' 
	AND ("Slot_Time") ILIKE ''|| $12 ||'%' 
	AND ("CreatedAt") ILIKE ''|| $13 ||'%'  
	ORDER BY ("CreatedAt") DESC LIMIT $14 OFFSET $15`

	rows, err := config.Db.Query(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Total_Order_Amount, q.Massage_Duration, q.Pincode,
		q.Merchant_Transaction_Id, q.Payment_Gateway_Id, q.Payment_Gateway_Mode, q.Transaction_Mode, q.Bank_Type, q.Slot_Date,
		q.Slot_Time, q.CreatedAt, q.Limit, offset)

	fmt.Println(12)

	if err != nil {
		// fmt.Print("asfafs")
		// fmt.Println(12)
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: err.Error()}
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
	("Customer_Souls_Id") ILIKE  ''||$1||'%' 
	AND ("Customer_Name") ILIKE ''|| $2 ||'%' 
	AND ("Total_Order_Amount") ILIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") ILIKE ''|| $4 ||'%' 
	AND ("Pincode") ILIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") ILIKE ''|| $6 ||'%' 
	AND ("Payment_Gateway_Id") ILIKE ''|| $7 ||'%' 
	AND ("Payment_Gateway_Mode") ILIKE ''|| $8 ||'%' 
	AND ("Transaction_Mode") ILIKE ''|| $9 ||'%' 
	AND ("Bank_Type") ILIKE ''|| $10 ||'%'
	AND ("Slot_Date") ILIKE ''|| $11 ||'%' 
	AND ("Slot_Time") ILIKE ''|| $12 ||'%' 
	AND ("CreatedAt") ILIKE ''|| $13 ||'%'   `
	cntRow := config.Db.QueryRow(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Total_Order_Amount, q.Massage_Duration, q.Pincode,
		q.Merchant_Transaction_Id, q.Payment_Gateway_Id, q.Payment_Gateway_Mode, q.Transaction_Mode, q.Bank_Type, q.Slot_Date,
		q.Slot_Time, q.CreatedAt)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		// fmt.Println(232)
		// fmt.Println(2398)
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: err.Error()}
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

func UpdateCustomerTransaction(w http.ResponseWriter, r *http.Request) (CustomerTran, ErorMesg) {
	// fmt.Println(12211)
	r.ParseForm()
	customer := CustomerTran{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		panic(err)
	}
	// response := Resp{}

	var res Shared.Result

	sqlStatement := ` UPDATE slh_transactions SET 
	"Number_Of_Therapist_Required" = $1, "Therapist_Gender" = $2, "Massage_For" = $3, "Slot_Time" = $4, "Slot_Date" = $5, "Customer_Address" = $6, 
	"Pincode" = $7, "Latitude" = $8,"Longitude" = $9, "Massage_Duration" = $10 WHERE ("Merchant_Transaction_Id") = $11`

	res, err = config.Db.Exec(sqlStatement,
		customer.Number_Of_Therapist, customer.Therapist_Gender, customer.Massage_For, customer.Slot_Time, customer.Slot_Date,
		customer.Customer_Address, customer.Pincode, customer.Latitude, customer.Longitude,
		customer.Massage_Duration, customer.Merchant_Transaction_Id)
	if err != nil {
		// fmt.Println(22)
		w.WriteHeader(http.StatusInternalServerError)
		return CustomerTran{}, ErorMesg{Message: err.Error()}
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		w.WriteHeader(http.StatusNotFound)
		return CustomerTran{}, ErorMesg{Message: err.Error()}
	}
	// BuildResp(&response, customer)
	return customer, ErorMesg{}
}

func ListAssign(w http.ResponseWriter, r *http.Request) ([]CustomerTran, ErrorMessage) {
	// fmt.Println(1266)
	return []CustomerTran{}, ErrorMessage{}
}
