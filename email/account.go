// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package email

import (
	"fmt"
	"log"
	"os"
	"soul_api/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(userName string, userEmail string) {
	fmt.Println(userName)
	fmt.Println(userEmail)
	from := mail.NewEmail("Souls","code.aks.010@gmail.com")
	subject := "Souls Team"
	to := mail.NewEmail(userName,userEmail)
	plainTextContent := `Hello`+userName
	htmlContent := `Hello User!`
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		// fmt.Println(response.Body)
		// fmt.Println(response.Headers)
	}
}

func SendEmailToAll(target string, content string) {
	fmt.Println(target)
	var sqlStatement string;
	if target == "CUSTOMER" {
		sqlStatement =`SELECT ("Customer_Name"), ("Customer_Email") FROM slh_customers WHERE 1=1`
	} else if target == "PARTNER" {
		sqlStatement =`SELECT (Partner_Name), ("Partner_Email") FROM slh_partners WHERE 1=1`
	} else {
		// return error
		fmt.Println("UNKNOWN TARGET TYPE!")
		return
	}

	rows, err := config.Db.Query(sqlStatement)
	// fmt.Println(rows)
	if err != nil {
		fmt.Print("asfafs")
		panic(err)
	}
	var name string;
	var email string
	for rows.Next() {
		rows.Scan(&name, &email)
		SendEmail(name, email)
	}

}