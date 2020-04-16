package comm_template

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
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

	sqlStatement := `SELECT ("Comm.Template_Id"),("Comm.Template_Type"),("Trigger_Time"),("Trigger_For"), ("SMS_Content"),("Email_Content"), 
	("Subject"),("Status") FROM slh_communication_templates WHERE 1=1 ;`
	rows, err := config.Db.Query(sqlStatement)
	// fmt.Println(rows)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return []Temp{}, Shared.ErrorMessage{Message: err.Error()}
	}

	// fmt.Println(len(rows))
	var temps []Temp
	for rows.Next() {
		var temp = Temp{}
		// fmt.Println(100)
		rows.Scan(&temp.Templ_id, &temp.Templ_type, &temp.Trigger_time, &temp.Trigger_for, &temp.SMS_content, &temp.Email_content, &temp.Subject, &temp.Status)
		temps = append(temps, temp)
		// cnt = cnt + 1
	}

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
