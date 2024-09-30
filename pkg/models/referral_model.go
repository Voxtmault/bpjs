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

type ReferredSpecialist struct {
	SpecalistCode string `json:"kodeSpesialis"`
	SpecalistName string `json:"namaSpesialis"`
	Capacity      string `json:"kapasitas"`
	ReferallCount string `json:"jumlahRujukan"`
	Percentage    string `json:"persentase"`
}
type ReferredSpecialistResponse struct {
	Lists []*ReferredSpecialist `json:"list"`
}

type ReferredFacility struct {
	FacilityCode string `json:"kodeSarana"`
	FacilityName string `json:"namaSarana"`
}
type ReferedFacilityResponse struct {
	Lists []*ReferredFacility `json:"list"`
}

type ReferralCreate struct {
	SEPNumber                  string `json:"noSep" validate:"required"`
	ReferralDate               string `json:"tglRujukan" validate:"required,datetime=2006-01-02"`
	PlannedVisitDate           string `json:"tglRencanaKunjungan" validate:"required,datetime=2006-01-02"`
	ReferredHealthFacilityCode string `json:"ppkDirujuk" validate:"required,len=8"`
	ServiceType                string `json:"jnsPelayanan" validate:"required,number,min=1"`
	Note                       string `json:"catatan"`
	ReferralDiagnosis          string `json:"diagRujukan" validate:"required,len=3"`
	ReferralType               string `json:"tipeRujukan" validate:"required,number,min=0"`
	ReferredPoliclinicCode     string `json:"poliRujukan" validate:"required_unless=ReferralType 2"`
	User                       string `json:"user" validate:"required"`
}
type ReferralCreateWrapper struct {
	TReferral *ReferralCreate `json:"t_rujukan"`
}

type ReferralCreateResponse struct {
	ReferralSource         Reference              `json:"AsalRujukan"`
	Diagnosis              Reference              `json:"diagnosa"`
	ReferralNumber         string                 `json:"noRujukan"`
	Participant            SEPParticipantResponse `json:"peserta"`
	ReferedPoliclinic      Reference              `json:"poliTujuan"`
	ValidVisitDate         string                 `json:"tglBerlakuKunjungan"`
	PlannedVisitDate       string                 `json:"tglRencanaKunjungan"`
	ReferralDate           string                 `json:"tglRujukan"`
	ReferredHealthFacility Reference              `json:"tujuanRujukan"`
}
type ReferralCreateResponseWrapper struct {
	Referral *ReferralCreateResponse `json:"rujukan"`
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
