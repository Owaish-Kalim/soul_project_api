package customer_with_partner

import (
	"time"
)

type CustomerPartner struct {
	Slot_Date    time.Time
	Slot_Time    time.Time
	CreatedAt    time.Time
	Customer_Souls_Id       string `json:"customer_souls_id"`
	Customer_Name           string `json:"customer_name"`
	Customer_Id             int    `json:"customer_id"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	Status                  string `json:"status"`
	Commission_Amount       int    `json:"commission_amount"`
	Created_By              string `json:"created_by"`
	Updated_By              string `json:"updated_by"`
}

