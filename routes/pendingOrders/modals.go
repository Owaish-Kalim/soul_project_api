package pendingOrders

import (
	"time"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type Tran struct {
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
}

type CustomerOrder struct {
	Customer_Name           string `json:"customer_name"`
	Customer_Order_Id       int    `json:"order_id"`
	Customer_Id             int    `json:"customer_id"`
	Customer_Souls_Id       string `json:"customer_souls_id"`
	Pincode                 int    `json:"pincode"`
	Customer_Address        string `json:"customer_address"`
	Number_Of_Therapist     int    `json:"number_of_therapist"`
	Therapist_Gender        string `json:"therapist_gender"`
	Massage_Duration        string `json:"massage_duration"`
	Massage_For             string `json:"massage_for"`
	Slot_Time               time.Time
	Slot_Date               time.Time
	Latitude                string `json:"latitude"`
	Longitude               string `json:"longitude"`
	Is_Order_Confirmed      bool   `json:"is_order_confirmed"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	CreatedAt               time.Time
	Total_Order_Amount      int `json:"total_order_amount"`
}

type ErorMsg struct {
	Message             string `json:"message"`
	Customer_Id         string `json:"customer_id"`
	Number_Of_Therapist string `json:"number_of_therapist"`
	Therapist_Gender    string `json:"therapist_gender"`
	Massage_Duration    string `json:"massage_duration"`
	Massage_For         string `json:"massage_for"`
}

type query struct {
	Limit                   int
	Page                    int
	Customer_Souls_Id       string
	Customer_Name           string
	Customer_Order_Id       string
	Customer_Mobile_No      int
	Status                  bool
	Customer_Gender         string
	Customer_Email          string
	Registered_Source       string
	Customer_Address        string
	Massage_Duration        string
	Massage_For             string
	Merchant_Transaction_Id string
	Total_Order_Amount      int
	Payment_Gateway_Id      string
	Payment_Gateway_Mode    string
	Transaction_Mode        string
	Bank_Type               string
}

func CheckEmpty(customer CustomerOrder, res *ErorMsg) {

	if customer.Number_Of_Therapist == 0 {
		res.Number_Of_Therapist = "Number of therapist cannot be empty."
		res.Message = "Error"
	}

	if customer.Customer_Id == 0 {
		res.Customer_Id = "Customer_Id cannot be empty."
		res.Message = "Error"
	}

	if customer.Therapist_Gender == "" {
		res.Therapist_Gender = "Therapist_Gender cannot be empty."
		res.Message = "Error"
	}

	if customer.Massage_Duration == "" {
		res.Massage_Duration = "Massage_Duration cannot be empty."
		res.Message = "Error"
	}

	if customer.Massage_For == "" {
		res.Massage_For = "Massage_For cannot be empty."
		res.Message = "Error"
	}

}
