package customers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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

func AddCustomer(w http.ResponseWriter, r *http.Request) (Customer, ErrorMsg) {
	fmt.Println(1)
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

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(customer.Customer_Mobile_No) == false {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Invalid Mobile No"
		return Customer{}, res
	}

	fmt.Println(1)
	check := Shared.MobileVerify(customer.Customer_Mobile_No)
	fmt.Println(1)

	if check != "correct" {
		res.Message = check
		return Customer{}, res
	}

	customer.Status = "Active"
	curr_time := time.Now()
	customer.CreatedAt = curr_time.Format("02-01-2006 3:4:5 PM")
	customer.Last_Access_Time = curr_time.Format("02-01-2006 3:4:5 PM")
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
			fmt.Println(32)
			w.WriteHeader(http.StatusInternalServerError)
			return customer, ErrorMsg{Message: "Internal Server Error"}
		}

		sqlStatements := `SELECT ("Customer_Id"), ("Customer_Souls_Id"), ("CreatedAt") FROM slh_customers WHERE ("Customer_Mobile_No")=$1`
		row := config.Db.QueryRow(sqlStatements, customer.Customer_Mobile_No)
		err = row.Scan(&customer.Customer_Id, &customer.Customer_Souls_Id, &customer.CreatedAt)
		if err != nil {
			fmt.Println(32)
			w.WriteHeader(http.StatusInternalServerError)
			return customer, ErrorMsg{Message: "Internal Server Error"}
		}
		fmt.Println(32)
		return customer, ErrorMsg{}

	}

	customer.Customer_Souls_Id = curr_time.Format("20060102") + strconv.Itoa(customer.Customer_Id)
	sqlStatement = `UPDATE slh_customers SET "Customer_Souls_Id" = $1 WHERE "Customer_Id" =  $2`
	_, err = config.Db.Exec(sqlStatement, customer.Customer_Souls_Id, customer.Customer_Id)
	if err != nil {

		return Customer{}, ErrorMsg{Message: "Internal Server Error."}
	}
	return customer, ErrorMsg{}
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

	var res Shared.Result

	curr_time := time.Now()
	customer.Last_Access_Time = curr_time.Format("02-01-2006 3:4:5 PM")

	sqlStatement := ` UPDATE slh_customers SET "Customer_Name" = $1, "Customer_Email" = $2, "Customer_Address" = $3, "Pincode" = $4, "Customer_Gender" = $5,
	"Status" = $6, "Registered_Source" = $7, "Last_Access_Time" = $8 WHERE ("Customer_Mobile_No") = $9`

	res, err = config.Db.Exec(sqlStatement, customer.Customer_Name, customer.Customer_Email, customer.Customer_Address, customer.Pincode,
		customer.Customer_Gender, customer.Status, customer.Registered_Source, customer.Last_Access_Time, customer.Customer_Mobile_No)
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
	fmt.Print("Owaish")
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []Customer{}, ErrorMessage{Message: "Internal Server Error"}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []Customer{}, ErrorMessage{Message: "Internal Server Error"}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}

	q.Customer_Souls_Id = r.Form.Get("customer_souls_id")
	q.Customer_Name = r.Form.Get("customer_name")
	q.Customer_Mobile_No = r.Form.Get("customer_mobile_no")
	q.Customer_Email = r.Form.Get("customer_email")
	q.Customer_Gender = r.Form.Get("customer_gender")
	q.Pincode = r.Form.Get("pincode")
	q.Status = r.Form.Get("status")
	// q.CreatedAt = r.Form  // Created

	fmt.Println(q)
	offset := q.Limit * q.Page
	// fmt.Print("ASHS")
	var customers []Customer
	sqlStatement := `SELECT ("Customer_Souls_Id"),("Customer_Name"),("Customer_Mobile_No"),("Status"), ("Customer_Id"), ("Customer_Gender"), ("Pincode"), 
	("Customer_Email"), ("Customer_Address"), ("Registered_Source") FROM slh_customers 
	WHERE ("Customer_Souls_Id") LIKE  ''||$1||'%' 
	AND ("Customer_Name") LIKE ''|| $2 ||'%' 
	AND ("Customer_Gender") LIKE ''|| $3 ||'%' 
	AND ("Customer_Email") LIKE ''|| $4 ||'%' 
	AND ("Customer_Mobile_No") LIKE ''|| $5 ||'%' 
	AND ("Pincode") LIKE ''|| $6 ||'%' 
	AND ("Status") LIKE ''|| $7 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $8 OFFSET $9`
	rows, err := config.Db.Query(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Customer_Gender, q.Customer_Email, q.Customer_Mobile_No,
		q.Pincode, q.Status, q.Limit, offset)
	// fmt.Println(rows)
	if err != nil {
		fmt.Print("asfafs")
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: "Internal Server Error."}
	}
	// fmt.Print("ASHS")
	// fmt.Println(len(rows))
	for rows.Next() {
		var customer = Customer{}
		rows.Scan(&customer.Customer_Souls_Id, &customer.Customer_Name, &customer.Customer_Mobile_No, &customer.Status,
			&customer.Customer_Id, &customer.Customer_Gender, &customer.Pincode, &customer.Customer_Email, &customer.Customer_Address,
			&customer.Registered_Source)
		customers = append(customers, customer)
		// cnt = cnt + 1
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_customers 
	WHERE ("Customer_Souls_Id") LIKE  ''||$1||'%' 
	AND ("Customer_Name") LIKE ''|| $2 ||'%' 
	AND ("Customer_Gender") LIKE ''|| $3 ||'%' 
	AND ("Customer_Email") LIKE ''|| $4 ||'%' 
	AND ("Customer_Mobile_No") LIKE ''|| $5 ||'%' 
	AND ("Pincode") LIKE ''|| $6 ||'%' 
	AND ("Status") LIKE ''|| $7 ||'%' `
	cntRow := config.Db.QueryRow(sqlStatement, q.Customer_Souls_Id, q.Customer_Name, q.Customer_Gender, q.Customer_Email, q.Customer_Mobile_No,
		q.Pincode, q.Status)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		// fmt.Println(232)
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
