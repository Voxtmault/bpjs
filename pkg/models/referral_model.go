package models

// Rujukan
type Referral struct {
	Diagnosis       Reference `json:"diagnosa"`
	Complaint       string    `json:"keluhan"`
	BPJSEncounterID string    `json:"noKunjungan"`
	Service         Reference `json:"pelayanan"`
	PoliReferral    Reference `json:"poliRujukan"`
	Referrer        Reference `json:"provPerujuk"`
	EncounterDate   string    `json:"tglKunjungan"`
}

// Unused for now since the function are divided into 2 different function
// type ReferralParams struct {
// 	ReferralNumber string
// 	BPJSNumber     string
// 	MultipleRecord *bool
// 	Source         string
// }

// Consts
const (
	PCareSource    uint = 1
	HospitalSource uint = 2
)
