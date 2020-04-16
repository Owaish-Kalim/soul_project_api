package comm_template

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
	"strconv"
)

func CreateTemp(w http.ResponseWriter, r *http.Request) (Temp, Shared.ErrorMessage) {

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	temp := Temp{}

	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		return temp, Shared.ErrorMessage{Message: err.Error()}
		panic(err)
	}

	sqlStatement := `
	INSERT INTO slh_communication_templates ("Comm.Template_Type","Trigger_Time","Trigger_For", "SMS_Content","Email_Content", "Subject", "Status")
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING "Comm.Template_Id"`
	// fmt.Println(time.Now())

	var target string
	if temp.Trigger_for == "customer" {
		target = "CUSTOMER"
	}
	if temp.Trigger_for == "partner" {
		target = "PARTNER"
	}
	fmt.Println(target)

	SendResponseAfter(10000000000, target, temp.Email_content, temp.SMS_content)

	temp.Templ_id = 0
	err = config.Db.QueryRow(sqlStatement, temp.Templ_type, temp.Trigger_time, temp.Trigger_for, temp.SMS_content, temp.Email_content, temp.Subject,
		temp.Status).Scan(&temp.Templ_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return temp, Shared.ErrorMessage{Message: err.Error()}
	}

	return temp, Shared.ErrorMessage{Message: ""}
}

func ListCom(w http.ResponseWriter, r *http.Request) ([]Temp, Shared.ErrorMessage) {
	r.ParseForm()
	var res Shared.ErrorMessage
	q := &query{}
	limit := r.Form.Get("limit")
	if limit != "" {
		if err := Shared.ParseInt(r.Form.Get("limit"), &q.Limit); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Message = err.Error()
			return []Temp{}, res
		}
	} else {
		q.Limit = 10
	}
	page := r.Form.Get("page")
	if page != "" {
		if err := Shared.ParseInt(r.Form.Get("page"), &q.Page); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Message = err.Error()
			return []Temp{}, res
		}
		q.Page = q.Page - 1
	} else {
		q.Page = 0
	}
	templ_id := r.Form.Get("templ_id")
	if templ_id != "" {
		if err := Shared.ParseInt(r.Form.Get("templ_id"), &q.Templ_id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res.Message = err.Error()
			return []Temp{}, res
		}
	}

	q.Templ_type = r.Form.Get("templ_type")
	q.Trigger_time = r.Form.Get("trigger_time")
	q.Trigger_for = r.Form.Get("trigger_for")
	q.SMS_content = r.Form.Get("sms_content")
	q.Subject = r.Form.Get("subject")
	q.Email_content = r.Form.Get("email_content")
	q.Status = r.Form.Get("status")

	// fmt.Println(q)
	offset := q.Limit * q.Page

	var temps []Temp
	sqlStatement := `SELECT ("Comm.Template_Id"),("Comm.Template_Type"),("Trigger_Time"),("Trigger_For"), ("SMS_Content"),("Email_Content"), 
	("Subject"),("Status") FROM slh_communication_templates 
	WHERE ("Comm.Template_Id")::text ilike  ''|| $1 ||'%'
	OR ("Comm.Template_Type") ILIKE ''|| $2 ||'%' 
	AND ("Trigger_For") ILIKE '' || $3 || '%' 
	AND ("SMS_Content") ILIKE '' ||$4|| '%'  
	AND ("Email_Content") ILIKE '' ||$5|| '%' 
	AND ("Subject") ILIKE ''|| $6 ||'%' 
	AND ("Trigger_Time") ILIKE ''|| $7 ||'%' 
	AND ("Status") ILIKE ''|| $8 ||'%' 
	ORDER BY ("Comm.Template_Id") LIMIT $9 OFFSET $10`
	rows, err := config.Db.Query(sqlStatement, q.Templ_id, q.Templ_type, q.Trigger_for, q.SMS_content, q.Email_content, q.Subject, q.Trigger_time,
		q.Status, q.Limit, offset)
	fmt.Println(12)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return temps, Shared.ErrorMessage{Message: err.Error()}
	}
	fmt.Println(12)

	for rows.Next() {
		var temp = Temp{}
		rows.Scan(&temp.Templ_id, &temp.Templ_type, &temp.Trigger_time, &temp.Trigger_for, &temp.SMS_content, &temp.Email_content, &temp.Subject, &temp.Status)
		temps = append(temps, temp)
	}

	sqlStatement = `SELECT COUNT(*) FROM slh_communication_templates  
	WHERE ("Comm.Template_Id")::text ilike  ''|| $1 ||'%'
	OR ("Comm.Template_Type") ILIKE ''|| $2 ||'%' 
	AND ("Trigger_For") ILIKE '' || $3 || '%' 
	AND ("SMS_Content") ILIKE '' ||$4|| '%'  
	AND ("Email_Content") ILIKE '' ||$5|| '%' 
	AND ("Subject") ILIKE ''|| $6 ||'%' 
	AND ("Trigger_Time") ILIKE ''|| $7 ||'%' 
	AND ("Status") ILIKE ''|| $8 ||'%' `
	cntRow := config.Db.QueryRow(sqlStatement, q.Templ_id, q.Templ_type, q.Trigger_for, q.SMS_content, q.Email_content, q.Subject, q.Trigger_time,
		q.Status)
	cnt := 0
	err = cntRow.Scan(&cnt)
	if err != nil {
		// panic(err)
		fmt.Println(12)
		w.WriteHeader(http.StatusInternalServerError)
		return temps, Shared.ErrorMessage{Message: err.Error()}
	}
	fmt.Println(12)
	w.Header().Set("Total-Count", strconv.Itoa(cnt))
	totalPages := cnt / q.Limit
	if cnt%q.Limit != 0 {
		totalPages = totalPages + 1
	}

	w.Header().Set("Total-Pages", strconv.Itoa(totalPages))

	return temps, Shared.ErrorMessage{Message: ""}
}

func UpdateComm(w http.ResponseWriter, r *http.Request) (Temp, Shared.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	var temp = Temp{}
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		panic(err)
	}

	sqlStatement := ` UPDATE slh_communication_templates SET "Status" = $1,"Comm.Template_Type" = $2,"Trigger_Time" = $3,"Trigger_For" = $4, 
	"SMS_Content" = $5, "Email_Content" = $6, "Subject" = $7	WHERE "Comm.Template_Id" = $8`

	_, err = config.Db.Exec(sqlStatement, temp.Status, temp.Templ_type, temp.Trigger_time, temp.Trigger_for, temp.SMS_content, temp.Email_content,
		temp.Subject, temp.Templ_id)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return temp, Shared.ErrorMessage{Message: err.Error()}
	}

	return temp, Shared.ErrorMessage{Message: ""}
}
