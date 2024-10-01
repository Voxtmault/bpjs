package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

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

// Referral Action is used for both creating and updating a referral due to the similarity of the fields
type ReferralAction struct {
	ReferralNumber             string `json:"noRujukan,omitempty" validate:"required_without=SEPNumber"`
	SEPNumber                  string `json:"noSep,omitempty" validate:"required_without=ReferralNumber"`
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
type ReferralActionWrapper struct {
	TReferral any `json:"t_rujukan"`
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

type OutgoingReferral struct {
	ReferralNumber             string `json:"noRujukan"`
	ReferralDate               string `json:"tglRujukan"`
	ServiceType                string `json:"jnsPelayanan"`
	SEPNumber                  string `json:"noSep"`
	CardNumber                 string `json:"noKartu"`
	ParticpantName             string `json:"nama"`
	ReferredHealthFacilityCode string `json:"ppkDirujuk"`
	ReferredHealthFacilityName string `json:"namaPpkDirujuk"`
}
type OutgoingReferralResponse struct {
	Lists []*OutgoingReferral `json:"list"`
}

type ReferralDetail struct {
	ReferralNumber             string `json:"noRujukan"`
	SEPNumber                  string `json:"noSep"`
	CardNumber                 string `json:"noKartu"`
	ParticipantName            string `json:"nama"`
	NursingClass               string `json:"kelasRawat"`
	Gender                     string `json:"kelamin"`
	DOB                        string `json:"tglLahir"`
	SEPDate                    string `json:"tglSep"`
	ReferralDate               string `json:"tglRujukan"`
	PlannedVisitDate           string `json:"tglRencanaKunjungan"`
	ReferredHealthFacilityCode string `json:"ppkDirujuk"`
	ReferredHealthFacilityName string `json:"namaPpkDirujuk"`
	ServiceType                string `json:"jnsPelayanan"`
	Note                       string `json:"catatan"`
	ReferralDiagnosis          string `json:"diagRujukan"`
	ReferralDiagnosisName      string `json:"namaDiagRujukan"`
	ReferralType               string `json:"tipeRujukan"`
	ReferralTypeName           string `json:"namaTipeRujukan"`
	ReferredPoliclinicCode     string `json:"poliRujukan"`
	ReferredPoliclinicName     string `json:"namaPoliRujukan"`
}
type ReferralDetailResponse struct {
	Referral *ReferralDetail `json:"rujukan"`
}

type SpecialReferralDiagnosis struct {
	Code string `json:"kode" validate:"required,specialDiag"`
}
type SpecialReferralCreate struct {
	ReferralNumber string                      `json:"noRujukan"`
	Diagnosises    []*SpecialReferralDiagnosis `json:"diagnosa"`
	Procedures     []*Reference                `json:"procedure"`
	User           string                      `json:"user"`
}

type SpecialReferralCreateResponse struct {
	ReferralNumber    string `json:"noRujukan"`
	CardNumber        string `json:"nokapst"`
	ParticipantName   string `json:"nmpst"`
	Diagnosis         string `json:"diagppk"`
	ReferralStartDate string `json:"tglrujukan_awal"`
	ReferralEndDate   string `json:"tglrujukan_berakhir"`
}
type SpecialReferralCreateResponseWrapper struct {
	Referral *SpecialReferralCreateResponse `json:"rujukan"`
}

type SpecialReferrals struct {
	ReferralID string `json:"idRujukan"`
	SpecialReferralCreate
}
type SpecialReferralsResponse struct {
	Referrals []*SpecialReferrals `json:"rujukan"`
}

type SpecialReferralDelete struct {
	ReferralID     string `json:"idRujukan"`
	ReferralNumber string `json:"noRujukan"`
	User           string `json:"user"`
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

// Custom Validation
func ValidateSpecialReferralDiagnosisCode(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[PS];`)
	return re.MatchString(fl.Field().String())
}
