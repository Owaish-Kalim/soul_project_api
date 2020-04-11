package transactions

import (
	"time"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type CustomerTran struct {
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
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	CreatedAt               time.Time
	Total_Order_Amount      int    `json:"total_order_amount"`
	Payment_Gateway_Id      string `json:"payment_gateway_id"`
	Payment_Gateway_Mode    string `json:"payment_gateway_mode"`
	Transaction_Mode        string `json:"transaction_mode"`
	Bank_Type               string `json:"bank_type"`
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

type Resp struct {
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	Payment_Gateway_Id      string `json:"payment_gateway_id"`
	Payment_Gateway_Mode    string `json:"payment_gateway_mode"`
	Transaction_Mode        string `json:"transaction_mode"`
	Bank_Type               string `json:"bank_type"`
}

type ErorMesg struct {
	Message                 string `json:"message"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	Payment_Gateway_Id      string `json:"payment_gateway_id"`
	Payment_Gateway_Mode    string `json:"payment_gateway_mode"`
	Transaction_Mode        string `json:"transaction_mode"`
	Bank_Type               string `json:"bank_type"`
}

func CheckEmptyTran(customer CustomerTran, res *ErorMesg) {

	// if customer.Merchant_Transaction_Id == "" {
	// 	res.Merchant_Transaction_Id = "Merchant_Transaction_Id cannot be empty."
	// 	res.Message = "Error"
	// }

	if customer.Payment_Gateway_Id == "" {
		res.Payment_Gateway_Id = "Payment_Gateway_Id cannot be empty."
		res.Message = "Error"
	}

	if customer.Payment_Gateway_Mode == "" {
		res.Payment_Gateway_Mode = "Payment_Gateway_Mode cannot be empty."
		res.Message = "Error"
	}

	if customer.Transaction_Mode == "" {
		res.Transaction_Mode = "Transaction_Mode cannot be empty."
		res.Message = "Error"
	}

	if customer.Bank_Type == "" {
		res.Bank_Type = "Bank_Type cannot be empty."
		res.Message = "Error"
	}

}
