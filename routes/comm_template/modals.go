package comm_template

type Temp struct {
	Templ_id      int    `json:"templ_id"`
	Templ_type    string `json:"templ_type"`
	Trigger_time  string `json:"trigger_time"`
	Trigger_for   string `json:"trigger_for"`
	SMS_content   string `json:"sms_content"`
	Subject       string `json:"subject"`
	Email_content string `json:"email_content"`
	Status        string `json:"status"`
}

type query struct {
	Limit         int
	Page          int
	Templ_id      int
	Templ_type    string
	Trigger_time  string
	Trigger_for   string
	SMS_content   string
	Subject       string
	Email_content string
	Status        string
}
