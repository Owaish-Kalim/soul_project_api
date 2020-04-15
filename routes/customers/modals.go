package customers

type ErrorMessage struct {
	Message string `json:"message"`
}

type Customer struct {
	Customer_Souls_Id  string `json:"customer_souls_id"`
	Customer_Id        int    `json:"customer_id"`
	Customer_Name      string `json:"customer_name"`
	Customer_Mobile_No string `json:"customer_mobile_no"`
	Customer_Gender    string `json:"customer_gender"`
	Pincode            string `json:"pincode"`
	Customer_Email     string `json:"customer_email"`
	Customer_Address   string `json:"customer_address"`
	Status             string `json:"status"`
	Last_Access_Time   string `json:"last_access_time"`
	Registered_Source  string `json:"registered_source" `
	CreatedAt          string `json:"created_at"`
}

type CustomerUpd struct {
	Customer_Id        int    `json:"customer_id"`
	Customer_Name      string `json:"customer_name"`
	Customer_Mobile_No string `json:"customer_mobile_no"`
	Customer_Gender    string `json:"customer_gender"`
	Pincode            string `json:"pincode"`
	Customer_Email     string `json:"customer_email"`
	Customer_Address   string `json:"customer_address"`
	Status             string `json:"status"`
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
	Limit              int
	Page               int
	Customer_Souls_Id  string
	Customer_Name      string
	Customer_Mobile_No string
	Status             string
	Customer_Gender    string
	Customer_Email     string
	Pincode            string
	CreatedAt          string
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

	if customer.Customer_Mobile_No == "" {
		res.Mobile_No = "Mobile_No cannot be empty."
		res.Message = "Error"
	}

	if customer.Customer_Gender == "" {
		res.Customer_Gender = "Customer_Gender cannot be empty."
		res.Message = "Error"
	}

	if customer.Pincode == "" {
		res.Pincode = "Pincode cannot be empty."
		res.Message = "Error"
	}

}
