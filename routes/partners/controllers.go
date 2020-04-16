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
	var res ErrPartner
	// response := Response{}
	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		return Partner{}, res
		// panic(err)
	}

	fmt.Println(12)
	res.Message = ""
	CheckEmpty(partner, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return partner, res
	}

	fmt.Println(12)

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(partner.Partner_MobileNo) == false {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = err.Error()
		return Partner{}, res
	}
	fmt.Println(12)

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
		// panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = err.Error()
		return partner, res
	}

	partner.Partner_Souls_Id = curr_time.Format("20060102") + strconv.Itoa(partner.Partner_Id)
	sqlStatement = `UPDATE slh_partners SET "Partner_Souls_Id" = $1 WHERE "Partner_Id" =  $2`
	_, err = config.Db.Exec(sqlStatement, partner.Partner_Souls_Id, partner.Partner_Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return Partner{}, ErrPartner{Message: err.Error()}
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
	"Pincode"=$5,"Latitude"=$6, "Longitude"=$7, "Per_Visit_Price_Commission"=$8,"Commission_Type"=$9,"Last_Updated_By"=$10, 
	"Onboard_Date"=$11,"CreatedBy"=$12 WHERE ("Partner_Email") = $13`

	result, err = config.Db.Exec(sqlStatement, partner.Partner_Name, partner.Partner_MobileNo, partner.Partner_Gender, partner.Partner_Address, partner.Pincode,
		partner.Latitude, partner.Longitude, partner.Rate, partner.Commission_Type, partner.UpdatedBy, partner.Onboard_Date, partner.CreatedBy, partner.Partner_Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return Partner{}, ErrPartner{Message: err.Error()}
		// panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count == 0 {
		// panic(err)
		w.WriteHeader(http.StatusNotFound)
		return Partner{}, ErrPartner{Message: err.Error()}
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
			return []Partner{}, ErrorMessage{Message: err.Error()}
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return []Partner{}, ErrorMessage{Message: err.Error()}
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}

	q.Partner_Name = r.Form.Get("partner_name")
	q.Partner_Email = r.Form.Get("partner_email")
	q.Partner_Gender = r.Form.Get("partner_gender")
	q.Partner_Souls_Id = r.Form.Get("partner_souls_id")
	q.Partner_MobileNo = r.Form.Get("partner_mobileno")
	q.Pincode = r.Form.Get("pincode")
	q.Rate = r.Form.Get("rate")
	q.Commission_Type = r.Form.Get("commission_type")
	q.UpdatedAt = r.Form.Get("updated_at")
	q.CreatedAt = r.Form.Get("created_at")

	offset := q.Limit * q.Page
	var partners []Partner

	sqlStatement := `SELECT ("Partner_Id"), ("Partner_Souls_Id"), ("Partner_Name"),("Partner_Email"), ("Partner_Mobile_No"),("Partner_Address"),("Pincode"),
	("Latitude"), ("Longitude"),("Per_Visit_Price_Commission"), ("Commission_Type"), ("Onboard_Date"),
	("UpdatedAt"), ("CreatedAt"),("Last_Updated_By"), ("CreatedBy"), ("Partner_Gender") FROM slh_partners
	WHERE ("Partner_Name") ILIKE  ''||$1||'%' 
	AND ("Partner_Email") ILIKE ''|| $2 ||'%' 
	AND ("Partner_Gender") ILIKE ''|| $3 ||'%' 
	AND ("Partner_Souls_Id") ILIKE ''|| $4 ||'%' 
	AND ("Partner_Mobile_No") ILIKE ''|| $5 ||'%' 
	AND ("Pincode") ILIKE ''|| $6 ||'%' 
	AND ("Per_Visit_Price_Commission") ILIKE ''|| $7 ||'%' 
	AND ("Commission_Type") ILIKE ''|| $8 ||'%' 
	AND ("UpdatedAt") ILIKE ''|| $9 ||'%'
	AND ("CreatedAt") ILIKE ''|| $10 ||'%' 
	ORDER BY ("CreatedAt") DESC LIMIT $11 OFFSET $12`

	rows, err := config.Db.Query(sqlStatement, q.Partner_Name, q.Partner_Email, q.Partner_Gender, q.Partner_Souls_Id, q.Partner_MobileNo, q.Pincode,
		q.Rate, q.Commission_Type, q.UpdatedAt, q.CreatedAt, q.Limit, offset)

	if err != nil {
		// fmt.Print("asfafs")
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return partners, ErrorMessage{Message: err.Error()}
	}

	for rows.Next() {
		partner := Partner{}
		rows.Scan(&partner.Partner_Id, &partner.Partner_Souls_Id, &partner.Partner_Name, &partner.Partner_Email,
			&partner.Partner_MobileNo, &partner.Partner_Address, &partner.Pincode, &partner.Latitude, &partner.Longitude,
			&partner.Rate, &partner.Commission_Type, &partner.Onboard_Date, &partner.UpdatedAt, &partner.CreatedAt,
			&partner.UpdatedBy, &partner.CreatedBy, &partner.Partner_Gender)
		partners = append(partners, partner)
		// cnt = cnt + 1
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_partners
	WHERE ("Partner_Name") ILIKE  ''||$1||'%' 
	AND ("Partner_Email") ILIKE ''|| $2 ||'%' 
	AND ("Partner_Gender") ILIKE ''|| $3 ||'%' 
	AND ("Partner_Souls_Id") ILIKE ''|| $4 ||'%' 
	AND ("Partner_Mobile_No") ILIKE ''|| $5 ||'%' 
	AND ("Pincode") ILIKE ''|| $6 ||'%' 
	AND ("Per_Visit_Price_Commission") ILIKE ''|| $7 ||'%' 
	AND ("Commission_Type") ILIKE ''|| $8 ||'%' 
	AND ("UpdatedAt") ILIKE ''|| $9 ||'%'
	AND ("CreatedAt") ILIKE ''|| $10 ||'%'`
	cntRow := config.Db.QueryRow(sqlStatement, q.Partner_Name, q.Partner_Email, q.Partner_Gender, q.Partner_Souls_Id, q.Partner_MobileNo, q.Pincode,
		q.Rate, q.Commission_Type, q.UpdatedAt, q.CreatedAt)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		// fmt.Println(232)
		// fmt.Println(4747)
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return partners, ErrorMessage{Message: err.Error()}
	}

	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt / q.Limit
	if cnt%q.Limit != 0 {
		totalPages = totalPages + 1
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	// fmt.Println(cnt)
	return partners, ErrorMessage{Message: ""}
}
