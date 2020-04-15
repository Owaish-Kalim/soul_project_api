package souls_settings

import (
	"encoding/json"
	"net/http"
	"soul_api/config"
	Shared "soul_api/routes"
)

func CreateSettings(w http.ResponseWriter, r *http.Request) (Temp, Shared.ErrorMessage) {

	w.Header().Set("Content-Type", "application/json")
	var res Shared.ErrorMessage
	r.ParseForm()
	temp := Temp{}

	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		panic(err)
	}

	sqlStatement := `
	INSERT INTO slh_souls_settings ("Type","URL", "Description","HostName", "UserName", "Password")
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING "Souls_Setting_Id"`

	temp.Souls_Setting_Id = 0
	err = config.Db.QueryRow(sqlStatement, temp.Type, temp.URL, temp.Description, temp.HostName, temp.UserName, temp.Password).Scan(&temp.Souls_Setting_Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = err.Error()
		return temp, res
	}

	return temp, err
}

func UpdateSettings(w http.ResponseWriter, r *http.Request) (Temp, Shared.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()

	var temp = Temp{}
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		panic(err)
	}

	sqlStatement := ` UPDATE slh_souls_settings SET "Type" = $1,"URL" = $2, "Description" = $3,"HostName" = $4, "UserName" = $5, "Password" = $6
	WHERE "Souls_Setting_Id" = $7`

	_, err = config.Db.Exec(sqlStatement, temp.Type, temp.URL, temp.Description, temp.HostName, temp.UserName, temp.Password, temp.Souls_Setting_Id)
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusPreconditionFailed)
		return temp, Shared.ErrorMessage{Message: err.Error()}
	}

	return temp, Shared.ErrorMessage{Message: ""}
}

func ListSettings(w http.ResponseWriter, r *http.Request) ([]Temp, Shared.ErrorMessage) {
	r.ParseForm()

	sqlStatement := `SELECT ("Souls_Setting_Id"), ("Type"),("URL"), ("Description"),("HostName"), ("UserName"), ("Password") FROM slh_souls_settings WHERE 1=1 ;`
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
		rows.Scan(&temp.Souls_Setting_Id, &temp.Type, &temp.URL, &temp.Description, &temp.HostName, &temp.UserName, &temp.Password)
		temps = append(temps, temp)
		// cnt = cnt + 1
	}

	return temps, Shared.ErrorMessage{Message: ""}
}
