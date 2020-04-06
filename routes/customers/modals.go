package customers

import (
	"time"
)

type Customer struct {
	Customer_Soul_Id  string `json:"customersoulid"`
	Customer_Id       int    `json:"customerid"`
	Customer_Name     string `json:"name"`
	Mobile_No         string `json:"mobile"`
	Customer_Gender   string `json:"gender"`
	Pincode           string `json:"pincode"`
	Customer_Email    string `json:"email"`
	Customer_Address  string `json:"address"`
	Status            string `json:"status"`
	Last_Access_Time  time.Time
	Registered_Source string `json:"registeredsource" `
	CreatedAt         time.Time
}

type ErrorMsg struct {
	Message           string `json:"message"`
	Customer_Name     string `json:"name"`
	Mobile_No         string `json:"mobile"`
	Customer_Gender   string `json:"gender"`
	Pincode           string `json:"pincode"`
	Customer_Email    string `json:"email"`
	Customer_Address  string `json:"address"`
	Status            string `json:"status"`
	Registered_Source string `json:"registeredsource" `
}

type CustomerOrder struct {
	Customer_Primary_Order_Id int    `json:"primaryorderid"`
	Customer_Id               int    `json:"customerid"`
	Customer_Soul_Id          string `json:"customersoulid"`
	Pincode                   string `json:"pincode"`
	Address                   string `json:"address"`
	Num_Therapist             int    `json:"numtherapist"`
	Therapist_Gender          string `json:"therapistgender"`
	Massage_Duration          string `json:"massageduration"`
	Massage_For               string `json:"massagefor"`
	Slot_Time                 time.Time
	Slot_Date                 time.Time
	Status                    string `json:"status"`
	Latitude                  string `json:"latitude"`
	Longitude                 string `json:"longitude"`
	Is_Order_Confirmed        string `json:"isorderconfirmed"`
	Transaction_Id            string `json:"transactionid"`
	CreatedAt                 time.Time
	Mobile_No                 string `json:"mobileno"`
}

type ErorMsg struct {
	Message            string `json:"message"`
	Num_Therapist      string `json:"numtherapist"`
	Therapist_Gender   string `json:"therapistgender"`
	Massage_Duration   string `json:"massageduration"`
	Massage_For        string `json:"massagefor"`
	Status             string `json:"status"`
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
	Is_Order_Confirmed string `json:"isorderconfirmed"`
	Mobile_No          string `json:"mobileno"`
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

	if customer.Mobile_No == "" {
		res.Mobile_No = "Mobile_No cannot be empty."
		res.Message = "Error"
	}

	if customer.Customer_Gender == "" {
		res.Customer_Gender = "Customer_Gender cannot be empty."
		res.Message = "Error"
	}

	if customer.Status == "" {
		res.Status = "Status cannot be empty."
		res.Message = "Error"
	}

	if customer.Pincode == "" {
		res.Pincode = "Pincode cannot be empty."
		res.Message = "Error"
	}

	if customer.Registered_Source == "" {
		res.Registered_Source = "Registered_Source cannot be empty."
		res.Message = "Error"
	}

}

func CheckEmpty(customer CustomerOrder, res *ErorMsg) {

	if customer.Mobile_No == "" {
		res.Mobile_No = "Mobile_No cannot be empty."
		res.Message = "Error"
	}

	if customer.Num_Therapist == 0 {
		res.Num_Therapist = "Number of therapist cannot be empty."
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

	if customer.Status == "" {
		res.Status = "Status cannot be empty."
		res.Message = "Error"
	}

	if customer.Latitude == "" {
		res.Latitude = "Latitude cannot be empty."
		res.Message = "Error"
	}

	if customer.Longitude == "" {
		res.Longitude = "Longitude_Source cannot be empty."
		res.Message = "Error"
	}

	if customer.Is_Order_Confirmed == "" {
		res.Is_Order_Confirmed = "Is_Order_Confirmed cannot be empty."
		res.Message = "Error"
	}

}
