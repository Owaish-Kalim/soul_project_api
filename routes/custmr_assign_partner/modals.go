package custmr_assign_partner

import (
	Shared "soul_api/routes"
	"time"
)

type Partner struct {
	Souls_Id string
	Dis      float64
}

type CustomerPartner struct {
	Slot_Date               string    `json:"slot_date"`
	Slot_Time               string    `json:"slot_time"`
	CreatedAt               time.Time //string
	Customer_Souls_Id       string    `json:"customer_souls_id"`
	Customer_Name           string    `json:"customer_name"`
	Customer_Gender         string    `json:"customer_gender"`
	Customer_Address        string    `json:"customer_address"`
	Customer_Pincode        string    `json:"pincode"`
	Merchant_Transaction_Id string    `json:"merchant_transaction_id"`
	Partner_Souls_Id        string    `json:"partner_souls_id"`
	Partner_Name            string    `json:"partner_name"`
	Partner_Mobile_No       string    `json:"partner_mobile_no"`
	Commission_Type         string    `json:"commission_type"`
	Id                      int       `json:"id"`
	Status                  string    `json:"status"`
	Commission_Amount       string    `json:"commission_amount"`
	Created_By              string    `json:"created_by"`
	Updated_By              string    `json:"updated_by"`
}

type Temp struct {
	Cust_Lat      string `json:"cust_lat"`
	Cust_Long     string `json:"cust_long"`
	Part_Lat      string `json:"part_lat"`
	Part_Long     string `json:"part_long"`
	Part_Souls_Id string `json:"part_souls_id"`
}

func CheckEmpty(customerpartner CustomerPartner, res *Shared.ErrorMesg) {

	if customerpartner.Commission_Amount == 0 {
		res.Commission_Amount = "Commission Amount cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Created_By == "" {
		res.Created_By = "Created By   cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Updated_By == "" {
		res.Updated_By = "Updated By  cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Customer_Souls_Id == "" {
		res.Customer_Souls_Id = "Customer Souls Id cannot be empty."
		res.Message = "Error"
	}

	if customerpartner.Status == "" {
		res.Status = "Status cannot be empty."
		res.Message = "Error"
	}

}
