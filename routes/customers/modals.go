package customers

import (
	"time"
)

type Tran struct {
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Customer struct {
	Customer_Souls_Id  string `json:"customer_souls_id"`
	Customer_Id        int    `json:"customer_id"`
	Customer_Name      string `json:"customer_name"`
	Customer_Mobile_No int    `json:"customer_mobile_no"`
	Customer_Gender    string `json:"customer_gender"`
	Pincode            int    `json:"pincode"`
	Customer_Email     string `json:"customer_email"`
	Customer_Address   string `json:"customer_address"`
	Status             bool   `json:"status"`
	Last_Access_Time   time.Time
	Registered_Source  string `json:"registered_source" `
	CreatedAt          time.Time
}

type CustomerUpd struct {
	Customer_Id        int    `json:"customer_id"`
	Customer_Name      string `json:"customer_name"`
	Customer_Mobile_No int    `json:"customer_mobile_no"`
	Customer_Gender    string `json:"customer_gender"`
	Pincode            int    `json:"pincode"`
	Customer_Email     string `json:"customer_email"`
	Customer_Address   string `json:"customer_address"`
	Status             bool   `json:"status"`
	Registered_Source  string `json:"registered_source" `
}

type ErrorMsg struct {
	Message          string `json:"message"`
	Customer_Name    string `json:"name"`
	Mobile_No        string `json:"mobile"`
	Customer_Gender  string `json:"gender"`
	Pincode          string `json:"pincode"`
	Customer_Email   string `json:"email"`
	Customer_Address string `json:"address"`
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

type ErorMesg struct {
	Message                 string `json:"message"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	Payment_Gateway_Id      string `json:"payment_gateway_id"`
	Payment_Gateway_Mode    string `json:"payment_gateway_mode"`
	Transaction_Mode        string `json:"transaction_mode"`
	Bank_Type               string `json:"bank_type"`
}

type Resp struct {
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
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

func CheckEmptyList(customer Customer, res *ErrorMsg) {

	if customer.Customer_Name == "" {
		res.Customer_Name = "Customer_Name cannot be empty."
		res.Message = "Error"
	}

	if customer.Customer_Address == "" {
		res.Customer_Address = "Customer_Address cannot be empty."
		res.Message = "Error"
	}

	if customer.Customer_Email == "" {
		res.Customer_Email = "Customer_Email cannot be empty."
		res.Message = "Error"
	}

	if customer.Customer_Mobile_No == 0 {
		res.Mobile_No = "Mobile_No cannot be empty."
		res.Message = "Error"
	}

	if customer.Customer_Gender == "" {
		res.Customer_Gender = "Customer_Gender cannot be empty."
		res.Message = "Error"
	}

	if customer.Pincode == 0 {
		res.Pincode = "Pincode cannot be empty."
		res.Message = "Error"
	}

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
