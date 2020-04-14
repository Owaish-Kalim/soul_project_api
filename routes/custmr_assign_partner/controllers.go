package custmr_assign_partner

import (
	"math"
)

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// func Distance(lat1, lon1, lat2, lon2 float64) float64 {

// 	var la1, lo1, la2, lo2, r float64
// 	la1 = lat1 * math.Pi / 180
// 	lo1 = lon1 * math.Pi / 180
// 	la2 = lat2 * math.Pi / 180
// 	lo2 = lon2 * math.Pi / 180

// 	r = 6378100
// 	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

// 	return 2 * r * math.Asin(math.Sqrt(h))
// }

// func CreateCustomerPartner(w http.ResponseWriter, r *http.Request) CustomerPartner {
// 	r.ParseForm()
// 	cust_part := CustomerPartner{}
// 	err := json.NewDecoder(r.Body).Decode(&cust_part)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//Select Partner
// 	temp := Temp{}
// 	sqlStatement := `SELECT ("Latitude"), ("Longitude") FROM slh_customers_pending_orders  WHERE ("Merchant_Transaction_Id")=$1;`
// 	row := config.Db.QueryRow(sqlStatement, cust_part.Merchant_Transaction_Id)
// 	err = row.Scan(&temp.Cust_Lat, &temp.Cust_Long)
// 	if err != nil {
// 		fmt.Print("asfafs")
// 		panic(err)
// 	}
// 	Lat1, _ := strconv.ParseFloat(temp.Cust_Lat, 8)
// 	Long1, _ := strconv.ParseFloat(temp.Cust_Long, 8)

// 	sqlStatement = `SELECT ("Latitude"), ("Longitude"), ("Partner_Souls_Id") FROM slh_partners  WHERE 1=1 ;`
// 	rows, err := config.Db.Query(sqlStatement)
// 	if err != nil {
// 		fmt.Print("asfafs")
// 		panic(err)
// 	}

// 	// MinDist := math.MaxFloat64
// 	var partners []Partner

// 	for rows.Next() {
// 		temp := Temp{}
// 		rows.Scan(&temp.Part_Lat, &temp.Part_Long, &temp.Part_Souls_Id)
// 		Lat2, _ := strconv.ParseFloat(temp.Part_Lat, 8)
// 		Long2, _ := strconv.ParseFloat(temp.Part_Long, 8)

// 		dist := Distance(Lat1, Long1, Lat2, Long2)

// 		partner := Partner{}
// 		partner.Souls_Id = temp.Part_Souls_Id
// 		partner.Dis = dist
// 		partners = append(partners, partner)
// 	}

// 	sort.SliceStable(partners, func(i, j int) bool {
// 		return partners[i].Dis < partners[j].Dis
// 	})

// 	sqlStatement = `SELECT ("Partner_Name"), ("Partner_Mobile_No") FROM slh_partners  WHERE ("Partner_Souls_Id")=$1; `
// 	row = config.Db.QueryRow(sqlStatement, cust_part.Partner_Souls_Id)
// 	err = row.Scan(&cust_part.Partner_Name, &cust_part.Partner_Mobile_No)
// 	if err != nil {
// 		fmt.Print("asfafs")
// 		panic(err)
// 	}

// 	cust_part.Slot_Date = "14-04-2020"
// 	cust_part.Slot_Time = "1 PM"
// 	cust_part.Status = "Pending"

// }

// func UpdateCustomerPartner(w http.ResponseWriter, r *http.Request) (CustomerPartner, Shared.ErrorMsg) {
// 	w.Header().Set("Content-Type", "application/json")
// 	r.ParseForm()
// 	var customerpartner = CustomerPartner{}
// 	err := json.NewDecoder(r.Body).Decode(&customerpartner)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var res Shared.ErrorMsg
// 	res.Message = ""

// 	CheckEmpty(customerpartner, &res)
// 	if res.Message != "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return customerpartner, res
// 	}

// 	userEmail := context.Get(r, middleware.Decoded)

// 	sqlStatement := ` UPDATE slh_teams SET "FirstName" = $1, "LastName" = $2, "Email" = $3, "Password" = $4, "Address" = $5, "MobileNo" = $6, "Status" = $7 WHERE ("Email") = $8`

// 	_, err = config.Db.Exec(sqlStatement, team.FirstName, team.LastName, team.Email, team.Password, team.Address, team.MobileNo, team.Status, userEmail)
// 	if err != nil {
// 		return response, Shared.ErrorMsg{Message: "Email already registered"}
// 		//panic(err)
// 	}
// 	BuildUpdateResponse(&response, team)
// 	return response, Shared.ErrorMsg{Message: ""}
// }
