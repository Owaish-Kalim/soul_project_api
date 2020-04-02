package customers

import (
	"time"
)

type Customer struct {
	Customer_Soul_Id string `json:"customersoulid"`
	Customer_Name    string `json:"name"`
	Mobile_No        string `json:"mobile"`
	Customer_Gender  string `json:"gender"`
	Pincode          string `json:"pincode"`
	Customer_Email   string `json:"email"`
	Customer_Address string `json:"address"`
	CreatedAt        time.Time
}
