package partners

import (
	"time"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type Partner struct {
	Partner_Id       int    `json:"partner_id"`
	Partner_Name     string `json:"partner_name"`
	Partner_Email    string `json:"partner_email"`
	Partner_MobileNo int    `json:"partner_mobileno"`
	Partner_Address  string `json:"partner_address"`
	Pincode          int    `json:"pincode"`
	Latitude         string `json:"latitude"`
	Longitude        string `"json:longitude"`
	Rate             int    `"rate"`
	Commission_Type  string `"json:"commission_type"`
	Onboard_Date     time.Time
	UpdatedAt        time.Time
	CreatedAt        time.Time
	CreatedBy        string `json:"created_by"`
	UpdatedBy        string `json:"updated_by"`
	Partner_Gender   string `json:"partner_gender"`
}

type ErrPartner struct {
	Message          string `json:message`
	Partner_Name     string `json:"partner_name"`
	Partner_Email    string `json:"partner_email"`
	Partner_MobileNo string `json:"partner_mobileno"`
	Partner_Address  string `json:"partner_address"`
	Pincode          string `json:"pincode"`
	Rate             string `"rate"`
	Commission_Type  string `"json:"commission_type"`
	Created_By       string `json:"created_by"`
	Updated_By       string `json:"updated_by"`
	Partner_Gender   string `json:"partner_gender"`
	Latitude         string `json:"latitude"`
	Longitude        string `"json:longitude"`
}

type UpResponse struct {
	Partner_Name     string `json:"partner_name"`
	Partner_Email    string `json:"partner_email"`
	Partner_MobileNo int    `json:"partner_mobileno"`
	Partner_Address  string `json:"partner_address"`
	Pincode          int    `json:"pincode"`
	Rate             int    `json:"rate"`
	Commission_Type  string `"json:"commission_type"`
	Updated_By       string `json:"updated_by"`
}

type query struct {
	Limit          int
	Page           int
	Partner_Name   string `json:"partner_name"`
	Partner_Email  string `json:"partner_email"`
	UpdatedBy      string `json:"updated_by"`
	Partner_Gender string `json:"partner_gender"`
}

func CheckEmpty(partner Partner, res *ErrPartner) {

	if partner.Partner_Name == "" {
		res.Partner_Name = "Partner_Name cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_Address == "" {
		res.Partner_Address = "Partner_Address cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_MobileNo == 0 {
		res.Partner_MobileNo = "Partner Mobile No cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_Email == "" {
		res.Partner_Email = "Partner_Emai cannot be empty."
		res.Message = "Error"
	}

	if partner.Pincode == 0 {
		res.Pincode = "Pincode cannot be empty."
		res.Message = "Error"
	}

	if partner.Rate == 0 {
		res.Rate = "Per Visit Price cannot be empty."
		res.Message = "Error"
	}

	if partner.Commission_Type == "" {
		res.Commission_Type = "Commission Type cannot be empty."
		res.Message = "Error"
	}

	if partner.Partner_Gender == "" {
		res.Partner_Gender = "Partner_Gender cannot be empty."
		res.Message = "Error"
	}

	if partner.CreatedBy == "" {
		res.Created_By = "Created By cannot be empty."
		res.Message = "Error"
	}

	if partner.UpdatedBy == "" {
		res.Updated_By = "Updated By cannot be empty."
		res.Message = "Error"
	}

	if partner.Latitude == "" {
		res.Latitude = "Latitude cannot be empty."
		res.Message = "Error"
	}

	if partner.Longitude == "" {
		res.Longitude = "Longitude cannot be empty."
		res.Message = "Error"
	}

}

// func CheckEmptyUp(partner Partner, res *Shared.ErrorMsg) {

// 	if partner.FirstName == "" {
// 		res.FirstName = "FirstName cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.LastName == "" {
// 		res.LastName = "LastName cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Partner_Address == "" {
// 		res.Address = "Address cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Partner_MobileNo == 0 {
// 		res.MobileNo = "MobileNo cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Partner_Alternate_MobileNo == 0 {
// 		res.MobileNo = "AlernateMobileNo cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Pincode == 0 {
// 		res.Status = "Pincode cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Latitude == "" {
// 		res.Status = "Latitude cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Longitude == "" {
// 		res.Status = "Longitude cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Per_Visit_Price_Commission == 0 {
// 		res.Status = "Per Visit Price cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Commission_Type == "" {
// 		res.Status = "Commission Type cannot be empty."
// 		res.Message = "Error"
// 	}

// 	if partner.Updated_By == "" {
// 		res.Status = "Updated By cannot be empty."
// 		res.Message = "Error"
// 	}

// }
