package Interfaces

type SystemInterface struct {
	CorsAllowOrigin  string
	CorsAllowHeader  string
	CorsAllowMethod  string
	AppName          string
	AppPrefork       bool
	AppCaseSensitive bool
	AppStrictRouting bool
	AppListenHost    string
	AppListenPort    int
	ServerHeader     string
}

type AddCustomer struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}
