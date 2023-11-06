package models

type User struct {
	LoginId     string `json:"login_id,omitempty"`
	Type        int    `json:"type,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

type UserOut struct {
	Response struct {
		Status       string `json:"status"`
		Message      string `json:"message"`
		EndpointName string `json:"endpoint_name"`
		Data         struct {
			UserRegistrationInfo struct {
				Registered        bool     `json:"registered"`
				ResendOtpWaitTime int      `json:"resend_otp_wait_time"`
				LoginTypes        []string `json:"login_types"`
			} `json:"user_registration_info"`
		} `json:"data"`
		ClientDetail struct {
			ActiveExperiments struct {
			} `json:"active_experiments"`
			UserAgentAnalysis struct {
				BrowserName    string `json:"browser_name"`
				BrowserVersion string `json:"browser_version"`
				DeviceType     string `json:"device_type"`
				OsName         string `json:"os_name"`
				OsVersion      string `json:"os_version"`
			} `json:"user_agent_analysis"`
			AppVersion    string `json:"app_version"`
			BatchExp      bool   `json:"batch_exp"`
			Country       string `json:"country"`
			Language      string `json:"language"`
			Currency      string `json:"currency"`
			CountryFromIp string `json:"country_from_ip"`
			Locale        string `json:"locale"`
			RegionFromIp  string `json:"region_from_ip"`
			Stage         string `json:"stage"`
			UserAgent     string `json:"user_agent"`
		} `json:"client_detail"`
	} `json:"response"`
}
