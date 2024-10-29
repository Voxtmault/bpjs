package models

// This function is used to provide SEP models to be used for SIM RS
type SelfSEP struct {
	CardNumber   string `json:"cardNumber"`
	ServiceDate  string `json:"serviceDate"`
	HospitalCode string `json:"hospitalCode"`
	ServiceType  string `json:"serviceType"`
}
