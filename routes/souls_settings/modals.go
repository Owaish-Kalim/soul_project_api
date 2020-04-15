package souls_settings

type Temp struct {
	Souls_Setting_Id int    `json:"souls_setting_id"`
	Type             string `json:"type"`
	URL              string `json:"url"`
	Description      string `json:"description"`
	HostName         string `json:"hostname"`
	UserName         string `json:"username"`
	Password         string `json:"password"`
}
