package vars

type EmailConfigSettingS struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type EmailNoticeSettingS struct {
	Receivers []string `json:"receivers"`
}
