package souls_settings

type Temp struct {
	Templ_id      int    `json:"templ_id"`
	Templ_type    string `json:"templ_type"`
	Trigger_time  string `json:"trigger_time"`
	Trigger_for   string `json:"Trigger_for"`
	SMS_content   string `json:"sms_content"`
	Subject       string `json:"subject"`
	Email_content string `json:"email_content"`
}
