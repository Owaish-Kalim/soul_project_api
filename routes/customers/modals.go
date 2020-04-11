package customers

import (
	"time"
)

type ErrorMessage struct {
	Message string `json:"message"`
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
