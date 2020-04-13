package partner

// import (
// 	"time"
// )

type Partner struct {
	PartnerId                    int    `json:"partnerid"`
	FirstName                    string `json:"firstname"`
	LastName                     string `json:"lastname"`
	MiddleName                   string `json:"middlename"`
	Partner_Email                string `json:"partner_email"`
	Partner_MobileNo             int    `json:"partner_mobileno"`
	Partner_Alternate_MobileNo   int    `json:"partner_alternate_mobileno"`
	Partner_Address              string `json:"partner_address"`
	Pincode                      int    `json:"pincode"`
	Latitude                     string `json:"latitude"`
	Longitude                    string `json:longitude"`
	Per_Visit_Price_Commission   int    `json:per_visit_price_commission"`
	Commission_Type              string `json:"commission_type"`
	Onboard_Date                 string `json:"onboard_date"`
	UpdatedAt                    string `json:"updatedat"`
	CreatedAt                    string `json:"createdat"`
	Created_By                   string `json:"created_by"`
	Updated_By                   string `json:"updated_by"`
	Partner_Gender               string `json:"partner_gender"`
}
 
type UpResponse struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	MiddleName  string `json:"middlename"`
	Partner_Email     string `json:"partner_email"`
	Partner_MobileNo  int `json:"partner_mobileno"`
	Partner_Alternate_MobileNo  int `json:"partner_alternate_mobileno"`
	Partner_Address   string `json:"partner_address"`
	Pincode int `json:"pincode"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Per_Visit_Price_Commission int `"per_visit_price_commission"`
	Commission_Type string `"json:"commission_type"`
    Onboard_Date                 string `json:"onboard_date"`
	UpdatedAt                    string `json:"updatedat"`
	Updated_By string `json:"updated_by"`
}