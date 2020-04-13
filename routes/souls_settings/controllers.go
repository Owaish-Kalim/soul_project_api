package souls_settings

import (
	"encoding/json"
	"net/http"
	"soul_api/config"
)

func CreateTemp(w http.ResponseWriter, r *http.Request) (Temp, error) {

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	temp := Temp{}

	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		panic(err)
	}

	sqlStatement := `
	INSERT INTO slh_communication_templates ("Comm.Template_Type","Trigger_Time","Trigger_For", "SMS_Content","Email_Content", "Subject")
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING "Comm.Template_Id"`

	temp.Templ_id = 0
	err = config.Db.QueryRow(sqlStatement, temp.Templ_type, temp.Trigger_time, temp.Trigger_for, temp.SMS_content, temp.Email_content, temp.Subject).Scan(&temp.Templ_id)
	if err != nil {
		return temp, err
	}

	return temp, err
}
