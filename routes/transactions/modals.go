package transactions

type ErrorMessage struct {
	Message string `json:"message"`
}

type Partner struct {
	Souls_Id string
	Dis      float64
}

type CustomerTran struct {
	Customer_Name           string `json:"customer_name"`
	Customer_Order_Id       int    `json:"order_id"`
	Customer_Id             int    `json:"customer_id"`
	Customer_Souls_Id       string `json:"customer_souls_id"`
	Pincode                 string `json:"pincode"`
	Customer_Address        string `json:"customer_address"`
	Number_Of_Therapist     string `json:"number_of_therapist"`
	Therapist_Gender        string `json:"therapist_gender"`
	Massage_Duration        string `json:"massage_duration"`
	Massage_For             string `json:"massage_for"`
	Slot_Time               string `json:"slot_time"`
	Slot_Date               string `json:"slot_date"`
	Latitude                string `json:"latitude"`
	Longitude               string `json:"longitude"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	CreatedAt               string `json:"created_at"`
	Total_Order_Amount      string `json:"total_order_amount"`
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
	Massage_Duration        string
	Merchant_Transaction_Id string
	Total_Order_Amount      string
	Payment_Gateway_Id      string
	Payment_Gateway_Mode    string
	Transaction_Mode        string
	Bank_Type               string
	Pincode                 string
}

type ErorMesg struct {
	Message                 string `json:"message"`
	Pincode                 string `json:"pincode"`
	Customer_Address        string `json:"customer_address"`
	Number_Of_Therapist     string `json:"number_of_therapist"`
	Therapist_Gender        string `json:"therapist_gender"`
	Massage_Duration        string `json:"massage_duration"`
	Massage_For             string `json:"massage_for"`
	Slot_Time               string `json:"slot_time"`
	Slot_Date               string `json:"slot_date"`
	Latitude                string `json:"latitude"`
	Longitude               string `json:"longitude"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	Total_Order_Amount      string `json:"total_order_amount"`
	Payment_Gateway_Id      string `json:"payment_gateway_id"`
	Payment_Gateway_Mode    string `json:"payment_gateway_mode"`
	Transaction_Mode        string `json:"transaction_mode"`
	Bank_Type               string `json:"bank_type"`
}

type CustomerPartner struct {
	Slot_Date               string `json:"slot_date"`
	Slot_Time               string `json:"slot_time"`
	CreatedAt               string `json:"created_at"`
	Customer_Souls_Id       string `json:"customer_souls_id"`
	Customer_Name           string `json:"customer_name"`
	Customer_Gender         string `json:"customer_gender"`
	Customer_Address        string `json:"customer_address"`
	Customer_Pincode        string `json:"pincode"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	Partner_Souls_Id        string `json:"partner_souls_id"`
	Partner_Name            string `json:"partner_name"`
	Partner_Email           string `json:"partner_email"`
	Partner_Mobile_No       string `json:"partner_mobile_no"`
	Commission_Type         string `json:"commission_type"`
	Id                      int    `json:"id"`
	Status                  string `json:"status"`
	Commission_Amount       string `json:"commission_amount"`
	Created_By              string `json:"created_by"`
	Updated_By              string `json:"updated_by"`
}

type Temp struct {
	Cust_Lat      string `json:"cust_lat"`
	Cust_Long     string `json:"cust_long"`
	Part_Lat      string `json:"part_lat"`
	Part_Long     string `json:"part_long"`
	Part_Souls_Id string `json:"part_souls_id"`
}

func CheckEmptyTran(customer CustomerTran, res *ErorMesg) {

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
