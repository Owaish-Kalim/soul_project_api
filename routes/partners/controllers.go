package partner

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"soul_api/config"
	//"soul_api/middleware"
	Shared "soul_api/routes"
	//"strconv"
	"time"

	//"github.com/dgrijalva/jwt-go"

	// "github.com/gemcook/pagination-go"
	//"github.com/gorilla/context"
	//"golang.org/x/crypto/bcrypt"
)

func BuildUpdateResponse(response *UpResponse, partner Partner) UpResponse {
	response.FirstName = partner.FirstName
	response.LastName = partner.LastName
	response.MiddleName=partner.MiddleName
	response.Partner_Email=partner.Partner_Email
	response.Partner_MobileNo=partner.Partner_MobileNo
	response.Partner_Alternate_MobileNo=partner.Partner_Alternate_MobileNo
	response.Partner_Address = partner.Partner_Address
	response.Pincode = partner.Pincode
	response.Latitude = partner.Latitude
	response.Longitude = partner.Longitude
	response.Per_Visit_Price_Commission=partner.Per_Visit_Price_Commission
	response.Commission_Type = partner.Commission_Type
	response.Onboard_Date=partner.Onboard_Date
	response.UpdatedAt=partner.UpdatedAt
	response.Updated_By=partner.Updated_By
	return *response
}

func CheckEmpty(partner Partner, res *Shared.ErrorMsg) {

	if partner.FirstName == "" {
		res.FirstName = "FirstName cannot be empty."
		res.Message = "Error"
	}

	if partner.LastName == "" {
		res.LastName = "LastName cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_Address == "" {
		res.Address = "Address cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_MobileNo == 0 {
		res.MobileNo = "MobileNo cannot be empty."
		res.Message = "Error"
	}
	
	if partner.Partner_Alternate_MobileNo == 0 {
		res.MobileNo = "AlernateMobileNo cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_Email == "" {
		res.Status = "Email cannot be empty."
		res.Message = "Error"
	}

    if partner.Pincode == 0{
		res.Status = "Pincode cannot be empty."
		res.Message = "Error"
	}

	if partner.Latitude == "" {
		res.Status = "Latitude cannot be empty."
		res.Message = "Error"
	}

	if partner.Longitude == "" {
		res.Status = "Longitude cannot be empty."
		res.Message = "Error"
	}

    if partner.Per_Visit_Price_Commission == 0 {
		res.Status = "Per Visit Price cannot be empty."
		res.Message = "Error"
	}
	
	if partner.Commission_Type == "" {
		res.Status = "Commission Type cannot be empty."
		res.Message = "Error"
	}
	
	// if partner.Onboard_Date == "" {
	// 	res.Status = "Onboard Date cannot be empty."
	// 	res.Message = "Error"
	// }

	if partner.Partner_Gender == "" {
		res.Status = "Gender cannot be empty."
		res.Message = "Error"
	}

	if partner.Created_By == "" {
		res.Status = "Created By cannot be empty."
		res.Message = "Error"
	}

	if partner.Updated_By == "" {
		res.Status = "Updated By cannot be empty."
		res.Message = "Error"
	}

}

func CheckEmptyUp(partner Partner, res *Shared.ErrorMsg) {

	if partner.FirstName == "" {
		res.FirstName = "FirstName cannot be empty."
		res.Message = "Error"
	}

	if partner.LastName == "" {
		res.LastName = "LastName cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_Address == "" {
		res.Address = "Address cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_MobileNo == 0 {
		res.MobileNo = "MobileNo cannot be empty."
		res.Message = "Error"
	}
	
	if partner.Partner_Alternate_MobileNo == 0 {
		res.MobileNo = "AlernateMobileNo cannot be empty."
		res.Message = "Error"
	}

    if partner.Pincode == 0{
		res.Status = "Pincode cannot be empty."
		res.Message = "Error"
	}

	if partner.Latitude == "" {
		res.Status = "Latitude cannot be empty."
		res.Message = "Error"
	}

	if partner.Longitude == "" {
		res.Status = "Longitude cannot be empty."
		res.Message = "Error"
	}

    if partner.Per_Visit_Price_Commission == 0 {
		res.Status = "Per Visit Price cannot be empty."
		res.Message = "Error"
	}
	
	if partner.Commission_Type == "" {
		res.Status = "Commission Type cannot be empty."
		res.Message = "Error"
	}

	if partner.Updated_By == "" {
		res.Status = "Updated By cannot be empty."
		res.Message = "Error"
	}

}

func CreatePartner(w http.ResponseWriter, r *http.Request) (Partner, Shared.ErrorMsg) {
	r.ParseForm()
	partner := Partner{}
	// response := Response{}
	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		panic(err)
	}

	var res Shared.ErrorMsg
	res.Message = ""
	CheckEmpty(partner, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return partner, res
	}
    //partner.Onboard_Date=time.Now().Local() 
	//partner.CreatedAt = time.Now().Local()
	//partner.UpdatedAt = time.Now().Local()

	currentTime :=time.Now()
	partner.UpdatedAt = currentTime.Format("02-01-2006 3:4:5 PM")
	partner.CreatedAt = currentTime.Format("02-01-2006 3:4:5 PM")
	partner.Onboard_Date = currentTime.Format("02-Jan-2006")

	sqlStatement := `
	INSERT INTO slh_partners ("FirstName","LastName","MiddleName", "Partner_Email", "Partner_MobileNo",
	 "Partner_Alternate_MobileNo", "Partner_Address", "Pincode", "Latitude", "Longitude", 
	 "Per_Visit_Price_Commission", "Commission_Type", "Onboard_Date", "UpdatedAt", "CreatedAt", "Created_By",
	  "Updated_By","Partner_Gender" )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	RETURNING ("PartnerId")`

	partner.PartnerId = 0
	err = config.Db.QueryRow(sqlStatement, partner.FirstName, partner.LastName, partner.MiddleName, partner.Partner_Email,
		 partner.Partner_MobileNo, partner.Partner_Alternate_MobileNo, partner.Partner_Address,
		 partner.Pincode, partner.Latitude, partner.Longitude, partner.Per_Visit_Price_Commission, partner.Commission_Type,
		 partner.Onboard_Date, partner.UpdatedAt, partner.CreatedAt, partner.Created_By, partner.Updated_By, 
		 partner.Partner_Gender).Scan(&partner.PartnerId)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		res.Message = "Email already registered"
		panic(err)
		return partner, res
	}
	// BuildResponse(&response, partner)
	res.Message = ""
	return partner, res
}

func UpdatePartner(w http.ResponseWriter, r *http.Request) (UpResponse, Shared.ErrorMsg) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	var partner = Partner{}
	var response = UpResponse{}
	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		panic(err)
	}

	var res Shared.ErrorMsg
	res.Message = ""
	CheckEmptyUp(partner, &res)
	if res.Message != "" {
		w.WriteHeader(http.StatusBadRequest)
		return response, res
	}
    fmt.Println(1)
	userEmail := partner.Partner_Email
	fmt.Println(userEmail)
	currentTime :=time.Now()
	partner.Onboard_Date = currentTime.Format("02-Jan-2006") 
	partner.UpdatedAt = currentTime.Format("02-01-2006 3:4:5 PM") 
	sqlStatement := ` UPDATE slh_partners SET "FirstName" = $1, "LastName" = $2, "MiddleName" =$3,
	 "Partner_MobileNo" = $4,"Partner_Alternate_MobileNo"=$5, "Partner_Address" = $6,"Pincode"=$7,"Latitude"=$8,
	 "Longitude"=$9, "Per_Visit_Price_Commission"=$10,"Commission_Type"=$11,"Updated_By"=$12 WHERE ("Partner_Email") = $13`

	_, err = config.Db.Exec(sqlStatement, partner.FirstName, partner.LastName, partner.MiddleName, 
		partner.Partner_MobileNo, partner.Partner_Alternate_MobileNo, partner.Partner_Address,partner.Pincode,
		partner.Latitude,partner.Longitude,partner.Per_Visit_Price_Commission,partner.Commission_Type,partner.Updated_By,
	    userEmail)
	if err != nil {
		panic(err)
	}
	BuildUpdateResponse(&response, partner)
	fmt.Println(response)
	return response, Shared.ErrorMsg{Message: ""}
}