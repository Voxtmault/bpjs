package models

type Reference struct {
	Code string `json:"kode"`
	Name string `json:"nama"`
}

type DiagnosisReference struct {
	Diagnosis []*Reference `json:"diagnosa"`
}
