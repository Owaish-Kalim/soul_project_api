package partners

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

func CreatePartner(w http.ResponseWriter, r *http.Request) (Partner, ErrPartner) {
	fmt.Println("owaas")
	r.ParseForm()
	partner := Partner{}
	// response := Response{}
	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		panic(err)
	}

	var res ErrPartner
	res.Message = ""
	CheckEmpty(partner, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return partner, res
	}

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(partner.Partner_MobileNo) == false {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Invalid Mobile No"
		return Partner{}, res
	}

	curr_time := time.Now()
	partner.UpdatedAt = curr_time.Format("02-01-2006 3:4:5 PM")
	partner.CreatedAt = curr_time.Format("02-01-2006 3:4:5 PM")

	sqlStatement := `
	INSERT INTO slh_partners ("Partner_Name", "Partner_Email", "Partner_Mobile_No", "Partner_Address", "Pincode", "Latitude", "Longitude", 
	"Per_Visit_Price_Commission", "Commission_Type", "Onboard_Date", "UpdatedAt", "CreatedAt", "Last_Updated_By","Partner_Gender", "CreatedBy")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) 
	RETURNING ("Partner_Id")`

	partner.Partner_Id = 0
	err = config.Db.QueryRow(sqlStatement, partner.Partner_Name, partner.Partner_Email, partner.Partner_MobileNo, partner.Partner_Address, partner.Pincode,
		partner.Latitude, partner.Longitude, partner.Rate, partner.Commission_Type, partner.Onboard_Date, partner.UpdatedAt, partner.CreatedAt,
		partner.UpdatedBy, partner.Partner_Gender, partner.CreatedBy).Scan(&partner.Partner_Id)
	if err != nil {
		fmt.Println("owaishhh")
		panic(err)
		// w.WriteHeader(http.StatusPreconditionFailed)
		// res.Message = "Email already registered"
		// return partner, res
	}

	partner.Partner_Souls_Id = curr_time.Format("20060102") + strconv.Itoa(partner.Partner_Id)
	sqlStatement = `UPDATE slh_partners SET "Partner_Souls_Id" = $1 WHERE "Partner_Id" =  $2`
	_, err = config.Db.Exec(sqlStatement, partner.Partner_Souls_Id, partner.Partner_Id)
	if err != nil {

		return Partner{}, ErrPartner{Message: "Internal Server Error."}
	}

	res.Message = ""
	return partner, res
}

func UpdatePartner(w http.ResponseWriter, r *http.Request) (Partner, ErrPartner) {
	fmt.Println("91919")
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var partner = Partner{}
	// var response = UpResponse{}
	err := json.NewDecoder(r.Body).Decode(&partner)
	fmt.Println(partner.Partner_Id)
	if err != nil {
		panic(err)
	}

	var res ErrPartner
	res.Message = ""
	CheckEmpty(partner, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return Partner{}, res
	}

	curr_time := time.Now()
	partner.UpdatedAt = curr_time.Format("02-01-2006 3:4:5 PM")
	fmt.Println(partner.Partner_Id)

	var result Shared.Result

	sqlStatement := `UPDATE slh_partners SET "Partner_Name" = $1, "Partner_Mobile_No" = $2,"Partner_Gender"=$3, "Partner_Address" = $4,
	"Pincode"=$5,"Latitude"=$6, "Longitude"=$7, "Per_Visit_Price_Commission"=$8,"Commission_Type"=$9,"Last_Updated_By"=$10 WHERE ("Partner_Id") = $11`

	result, err = config.Db.Exec(sqlStatement, partner.Partner_Name, partner.Partner_MobileNo, partner.Partner_Gender, partner.Partner_Address, partner.Pincode,
		partner.Latitude, partner.Longitude, partner.Rate, partner.Commission_Type, partner.UpdatedBy, partner.Partner_Id)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		// panic(err)
		w.WriteHeader(http.StatusNotFound)
		return Partner{}, ErrPartner{Message: "Unauthorised User"}
	}

	return partner, ErrPartner{}
}

func ListPartner(w http.ResponseWriter, r *http.Request) ([]Partner, ErrorMessage) {
	r.ParseForm()
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []Partner{}, ErrorMessage{Message: "Internal Server Error."}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []Partner{}, ErrorMessage{Message: "Internal Server Error."}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}

	q.Partner_Name = r.Form.Get("partner_name")
	q.Partner_Email = r.Form.Get("partner_email")
	q.Partner_Gender = r.Form.Get("partner_gender")

	offset := q.Limit * q.Page
	var customers []Partner

	sqlStatement := `SELECT ("Partner_Id"), ("Partner_Name"),("Partner_Email"), ("Partner_Mobile_No"),("Partner_Address"),("Pincode"),("Latitude") ,
	("Longitude"),("Per_Visit_Price_Commission"), ("Commission_Type"), ("Onboard_Date"),
	("UpdatedAt"), ("CreatedAt"),("Last_Updated_By"), ("Partner_Gender") FROM slh_partners
	WHERE ("Partner_Name") LIKE  ''||$1||'%' 
	AND ("Partner_Email") LIKE ''|| $2 ||'%' 
	AND ("Partner_Gender") LIKE ''|| $3 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $4 OFFSET $5`

	rows, err := config.Db.Query(sqlStatement, q.Partner_Name, q.Partner_Email, q.Partner_Gender, q.Limit, offset)

	if err != nil {
		// fmt.Print("asfafs")
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return customers, ErrorMessage{Message: "Internal Server Error."}
	}

	for rows.Next() {
		customer := Partner{}
		rows.Scan(&customer.Partner_Id, &customer.Partner_Name, &customer.Partner_Email,
			&customer.Partner_MobileNo, &customer.Partner_Address, &customer.Pincode, &customer.Latitude, &customer.Longitude,
			&customer.Rate, &customer.Commission_Type, &customer.Onboard_Date, &customer.UpdatedAt, &customer.CreatedAt,
			&customer.UpdatedBy, &customer.Partner_Gender)
		customers = append(customers, customer)
		// cnt = cnt + 1
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_partners
	WHERE ("Partner_Name") LIKE  ''||$1||'%' 
	AND ("Partner_Email") LIKE ''|| $2 ||'%' 
	AND ("Partner_Gender") LIKE ''|| $3 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.Partner_Name, q.Partner_Email, q.Partner_Gender)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		// fmt.Println(232)
		// fmt.Println(4747)
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
