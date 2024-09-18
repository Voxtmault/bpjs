package models

type COB struct {
	InsuranceName   string `json:"nmAsuransi"`
	InsuranceNumber string `json:"noAsuransi"`
	TATDate         string `json:"tglTAT"`
	TMTDate         string `json:"tglTMT"`
}

type Information struct {
	Dinsos      string `json:"dinsos"`
	SKTMNumber  string `json:"noSKTM"`
	ProlanisPRB string `json:"prolanisPRB"`
}

type BPJSMedicalRecord struct {
	MRNumber    string `json:"noMR"`
	PhoneNumber string `json:"noTelepon"`
}

type ProvUmum struct {
	ProviderCode string `json:"kdProvider"`
	ProviderName string `json:"nMProvider"`
}

type PatientAge struct {
	AgeNow       string `json:"umurSekarang"`
	AgeAtService string `json:"umurSaatPelayanan"`
}

type ClassRights struct {
	ReusableNote
}
type ParticipantType struct {
	ReusableNote
}
type ParticipantStatus struct {
	ReusableNote
}

type BPJSParticipant struct {
	Name              string            `json:"nama"`
	NIK               string            `json:"nik"`
	CardNumber        string            `json:"noKartu"`
	Pisa              string            `json:"pisa"`
	Sex               string            `json:"sex"`
	TATDate           string            `json:"tglTAT"`
	TMTDate           string            `json:"tglTMT"`
	DOB               string            `json:"tglLahir"`
	CardPrintDate     string            `json:"tglCetakKartu"`
	COB               COB               `json:"cob"`
	ClassRights       ClassRights       `json:"hakKelas"`
	Information       Information       `json:"informasi"`
	ParticipantType   ParticipantType   `json:"jenisPeserta"`
	MedicalRecord     BPJSMedicalRecord `json:"mr"`
	ProvUmum          ProvUmum          `json:"provUmum"`
	ParticipantStatus ParticipantStatus `json:"statusPeserta"`
	Age               PatientAge        `json:"umur"`
}

// Reusable Class
type ReusableNote struct {
	Code string `json:"kode"`
	Note string `json:"keterangan"`
}

// Search Params
type ParticipantSearchParams struct {
	NIK         string `validate:"omitempty,numeric,min=16" example:"1234567890123456"`
	BPJSNumber  string `validate:"omitempty,numeric,min=13" example:"1234567890123"`
	ServiceDate string `validate:"omitempty,datetime=2006-01-02" example:"2021-01-01"`
}
