package comm_template

import (
	"fmt"
	"soul_api/email"
	"time"
)

func SendResponseAfter(d time.Duration, target string, emailContent string, smsContent string) {
	time.AfterFunc(d, func() {
		fmt.Println(emailContent)
		fmt.Println(time.Now())
		if smsContent!="" {
			fmt.Println("SEND SMS")	
		}
		if emailContent != "" {
			fmt.Println("CALLING SEND EMAIL")
			email.SendEmailToAll(target, emailContent)
		}
	})
}