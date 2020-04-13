// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package email

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(userName string, userEmail string) {
	fmt.Println(userName)
	fmt.Println(userEmail)
	from := mail.NewEmail("","asingh2@ch.iitr.ac.in")
	subject := "CONFESSION"
	to := mail.NewEmail("",userEmail)
	plainTextContent := `DIVIT, YOU ARE THE LOVE OF LIFE. I WANT TO SPEND MY ENTIRE LIFE WITH YOU.
	PLEASE BE MINE...
	PLZZ LOVE ME THE WAY YOU DO, DIVIT...
	WAITING FOR YOUR REPLY...
	PLZZ CALL ME`
	fmt.Println(plainTextContent)


	htmlContent := `	
		<div style="margin:auto, background:yellow, height: 100px">
		DIVIT, YOU ARE THE LOVE OF LIFE. I WANT TO SPEND MY ENTIRE LIFE WITH YOU.
		PLEASE BE MINE...
		PLZZ LOVE ME THE WAY YOU DO, DIVIT...
		WAITING FOR YOUR REPLY...
		PLZZ CALL ME
		</div>

	`
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}