package models

type Reference struct {
	Code string `json:"kode"`
	Name string `json:"nama"`
}

type DiagnosisReference struct {
	Diagnosis []*Reference `json:"diagnosa"`
}

type DoctorReference struct {
	Doctor []*Reference `json:"list"`
}

type ListReference struct {
	List []*Reference `json:"list"`
}

type PoliReference struct {
	Poli []*Reference `json:"poli"`
}

type FaskesReference struct {
	Faskes []*Reference `json:"faskes"`
}

type ProcedureReference struct {
	Procedure []*Reference `json:"procedure"`
}
