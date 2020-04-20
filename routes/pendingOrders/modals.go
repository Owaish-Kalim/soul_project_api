package pendingOrders

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
	Is_Order_Confirmed      string `json:"is_order_confirmed"`
	Merchant_Transaction_Id string `json:"merchant_transaction_id"`
	CreatedAt               string `json:"created_at"`
	Total_Order_Amount      string `json:"total_order_amount"`
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
	Is_Order_Confirmed      string
	Pincode                 string
	Massage_Duration        string
	Merchant_Transaction_Id string
	Total_Order_Amount      string
	Slot_Date               string
	Slot_Time               string
	CreatedAt               string
}

func CheckEmpty(customer CustomerOrder, res *ErorMsg) {

	if customer.Number_Of_Therapist == "" {
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
