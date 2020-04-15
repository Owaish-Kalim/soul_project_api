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
)

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
		fmt.Print("asfafs")
		panic(err)
	}
	Lat1, _ := strconv.ParseFloat(temp.Cust_Lat, 8)
	Long1, _ := strconv.ParseFloat(temp.Cust_Long, 8)
	Number_Of_Therapist, _ := strconv.Atoi(Number_Of_Therapists)

	sqlStatement = `SELECT ("Latitude"), ("Longitude"), ("Partner_Souls_Id") FROM slh_partners  WHERE 1=1 ;`
	rows, err := config.Db.Query(sqlStatement)
	if err != nil {
		fmt.Print("asfafs")
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
			fmt.Print("asfafs")
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

	sqlStatement = `SELECT ("Partner_Name"), ("Partner_Mobile_No") FROM slh_partners  WHERE ("Partner_Souls_Id")=$1; `
	row = config.Db.QueryRow(sqlStatement, cust_part.Partner_Souls_Id)
	err = row.Scan(&cust_part.Partner_Name, &cust_part.Partner_Mobile_No)
	if err != nil {
		fmt.Print("asfafs")
		panic(err)
	}

	cust_part.Slot_Date = "14-04-2020"
	cust_part.Slot_Time = "1 PM"
	cust_part.Status = "Pending"

	return CustomerPartner{}

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
	// fmt.Println(12345678)
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
		return CustomerTran{}, ErrorMessage{Message: "Internal Server Error"}
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
			return []CustomerTran{}, ErrorMessage{Message: "Internal Server Error"}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []CustomerTran{}, ErrorMessage{Message: "Internal Server Error"}
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

	offset := q.Limit * q.Page
	var customers []CustomerTran
	sqlStatement := `SELECT ("Customer_Order_Id"), ("Customer_Souls_Id"),("Customer_Name"), ("Customer_Id"),("Pincode"),("Customer_Address"),
	("Massage_Duration") ,("Number_Of_Therapist_Required"),("Massage_For"), ("Therapist_Gender"), ("Merchant_Transaction_Id"),
	("Total_Order_Amount"), ("Latitude"),("Longitude"), ("CreatedAt"), ("Slot_Date"), ("Slot_Time"), ("Payment_Gateway_Id"), ("Payment_Gateway_Mode"), 
	("Transaction_Mode"), ("Bank_Type")  FROM slh_transactions
	WHERE ("Customer_Souls_Id") LIKE  ''||$1||'%' 
	AND ("Customer_Name") LIKE ''|| $2 ||'%' 
	AND ("Total_Order_Amount") LIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") LIKE ''|| $4 ||'%' 
	AND ("Pincode") LIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") LIKE ''|| $6 ||'%' 
	AND ("Payment_Gateway_Id") LIKE ''|| $7 ||'%' 
	AND ("Payment_Gateway_Mode") LIKE ''|| $8 ||'%' 
	AND ("Transaction_Mode") LIKE ''|| $9 ||'%' 
	AND ("Bank_Type") LIKE ''|| $10 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $11 OFFSET $12`

	rows, err := config.Db.Query(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Total_Order_Amount, q.Massage_Duration, q.Pincode,
		q.Merchant_Transaction_Id, q.Payment_Gateway_Id, q.Payment_Gateway_Mode, q.Transaction_Mode, q.Bank_Type, q.Limit, offset)

	fmt.Println(12)

	if err != nil {
		// fmt.Print("asfafs")
		// fmt.Println(12)
		// panic(err)
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
	AND ("Total_Order_Amount") LIKE ''|| $3 ||'%' 
	AND ("Massage_Duration") LIKE ''|| $4 ||'%' 
	AND ("Pincode") LIKE ''|| $5 ||'%' 
	AND ("Merchant_Transaction_Id") LIKE ''|| $6 ||'%' 
	AND ("Payment_Gateway_Id") LIKE ''|| $7 ||'%' 
	AND ("Payment_Gateway_Mode") LIKE ''|| $8 ||'%' 
	AND ("Transaction_Mode") LIKE ''|| $9 ||'%' 
	AND ("Bank_Type") LIKE ''|| $10 ||'%' `
	cntRow := config.Db.QueryRow(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Total_Order_Amount, q.Massage_Duration, q.Pincode,
		q.Merchant_Transaction_Id, q.Payment_Gateway_Id, q.Payment_Gateway_Mode, q.Transaction_Mode, q.Bank_Type)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		// fmt.Println(232)
		// fmt.Println(2398)
		// panic(err)
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

	sqlStatement := ` UPDATE slh_transactions SET "Payment_Gateway_Id" = $1, "Bank_Type" = $2, "Payment_Gateway_Mode" = $3, "Transaction_Mode" = $4,
	"Number_Of_Therapist_Required" = $5, "Therapist_Gender" = $6, "Massage_For" = $7, "Slot_Time" = $8, "Slot_Date" = $9, "Customer_Address" = $10, 
	"Pincode" = $11, "Latitude" = $12,"Longitude" = $13, "Total_Order_Amount" = $14, "Massage_Duration" = $15 WHERE ("Merchant_Transaction_Id") = $16`

	res, err = config.Db.Exec(sqlStatement, customer.Payment_Gateway_Id, customer.Bank_Type, customer.Payment_Gateway_Mode, customer.Transaction_Mode,
		customer.Number_Of_Therapist, customer.Therapist_Gender, customer.Massage_For, customer.Slot_Time, customer.Slot_Date,
		customer.Customer_Address, customer.Pincode, customer.Latitude, customer.Longitude, customer.Total_Order_Amount,
		customer.Massage_Duration, customer.Merchant_Transaction_Id)
	if err != nil {
		// fmt.Println(22)
		w.WriteHeader(http.StatusInternalServerError)
		return CustomerTran{}, ErorMesg{Message: "Internal Server Error"}
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		w.WriteHeader(http.StatusNotFound)
		return CustomerTran{}, ErorMesg{Message: "Unauthorised User"}
	}
	// BuildResp(&response, customer)
	return customer, ErorMesg{}
}
