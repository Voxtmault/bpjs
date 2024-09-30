package models

type GeneralProvider struct {
	ProviderCode string `json:"kdProvider"`
	ProviderName string `json:"nmProvider"`
}

type ReferingProvider struct {
	ReferingProviderCode string `json:"kdProviderPerujuk"`
	ReferingProviderName string `json:"nmProviderPerujuk"`
	ReferalSource        string `json:"asalRujukan"`
	ReferalNumber        string `json:"noRujukan"`
	ReferalDate          string `json:"tglRujukan"`
}

type ControlPlanGetViaSEP struct {
	SEPNumber        string                 `json:"noSep"`
	SEPDate          string                 `json:"tglSep"`
	ServiceType      string                 `json:"jnsPelayanan"`
	Policlinic       string                 `json:"poli"`
	Diagnosis        string                 `json:"diagnosa"`
	Participant      SEPParticipantResponse `json:"peserta"`
	GeneralProvider  GeneralProvider        `json:"provUmum"`
	ReferingProvider ReferingProvider       `json:"provPerujuk"`
}

type ControlPlanGetViaControllLetterNumber struct {
	ControlLetterNumber string               `json:"noSuratKontrol"`
	ControlPlanDate     string               `json:"tglRencanaKontrol"`
	IssuedDate          string               `json:"tglTerbit"`
	ControlType         string               `json:"jnsKontrol"`
	TargetPoliCode      string               `json:"poliTujuan"`
	TargetPoliName      string               `json:"namaPoliTujuan"`
	TargetDoctorCode    string               `json:"kodeDokter"`
	TargetDoctorName    string               `json:"namaDokter"`
	ControlFlag         string               `json:"flagKontrol"`
	IssuerDoctorCode    string               `json:"kodeDokterPembuat"`
	IssuerDoctorName    string               `json:"namaDokterPembuat"`
	ControlTypeName     string               `json:"namaJnsKontrol"`
	SEP                 ControlPlanGetViaSEP `json:"sep"`
}

type ClinicControlParams struct {
	ControlType        string `validate:"required,number,min=1"`
	Identifier         string `validate:"required"`
	ControlPlannedDate string `validate:"required,datetime=2006-01-02"`
}

type DoctorScheduleParams struct {
	ControlType        string `validate:"required,number,min=1"`
	PoliCode           string `validate:"required,uppercase,min=3"`
	ControlPlannedDate string `validate:"required,datetime=2006-01-02"`
}

type ClinicControlPlans struct {
	ClinicCode                string `json:"kodePoli"`
	ClinicName                string `json:"namaPoli"`
	Capacity                  string `json:"kapasitas"`
	AssignedControlandReferal string `json:"jmlRencanaKontroldanRujukan"`
	Percentage                string `json:"persentase"`
}
type ClinicControlPlansResponse struct {
	Lists []*ClinicControlPlans `json:"list"`
}

type DoctorPracticeSchedule struct {
	DoctorCode       string `json:"kodeDokter"`
	DoctorName       string `json:"namaDokter"`
	PracticeSchedule string `json:"jadwalPraktek"`
	Capacity         string `json:"kapasitas"`
}
type DoctorPracticeScheduleResponse struct {
	Lists []*DoctorPracticeSchedule `json:"list"`
}

// ControlPlanCreate struct is used for both creating regular controll plan and
// inpatient controll plan / inpatient care order due to it's similarity
type ControlPlanCreate struct {
	CardNumber  string `json:"noKartu,omitempty"`
	SEPNumber   string `json:"noSEP,omitempty"`
	DoctorCode  string `json:"kodeDokter"`
	ClinicCode  string `json:"poliKontrol"`
	ControlDate string `json:"tglRencanaKontrol"`
	User        string `json:"user"`
}

// ControlPlanCreateResponse struct is used for both parsing regular controll plan and
// inpatient controll plan / inpatient care order response from BPJS API due to it's similarity
type ControlPlanCreateResponse struct {
	InpatientCareOrderNumber string `json:"noSPRI,omitempty"`
	ControlLetterNumber      string `json:"noSuratKontrol,omitempty"`
	ControlDate              string `json:"tglRencanaKontrol"`
	DoctorName               string `json:"namaDokter"`
	CardNumber               string `json:"noKartu"`
	ParticipantName          string `json:"nama"`
	ParticipantGender        string `json:"kelamin"`
	ParticipantDOB           string `json:"tglLahir"`
	Diagnosis                string `json:"namaDiagnosa,omitempty"`
}

type ControlPlans struct {
	ControlLetterNumber string `json:"noSuratKontrol"`
	ServiceType         string `json:"jnsPelayanan"`
	ControlTypeCode     string `json:"jnsKontrol"`
	ControlTypeName     string `json:"namaJnsKontrol"`
	ControlDate         string `json:"tglRencanaKontrol"`
	ControlIssuedDate   string `json:"tglTerbitKontrol"`
	SourceSEP           string `json:"noSepAsalKontrol"`
	SourcePoliCode      string `json:"poliAsal"`
	SourcePoliName      string `json:"namaPoliAsal"`
	TargetPoliCode      string `json:"poliTujuan"`
	TargetPoliName      string `json:"namaPoliTujuan"`
	SEPDate             string `json:"tglSEP"`
	DoctorCode          string `json:"kodeDokter"`
	DoctorName          string `json:"namaDokter"`
	CardNumber          string `json:"noKartu"`
	ParticipantName     string `json:"nama"`
	SEPIssued           string `json:"terbitSEP,omitempty"`
}
type ControlPlansResponse struct {
	Lists []*ControlPlans `json:"list"`
}

type ControlPlansFromCardNumberParams struct {
	Month      string `validate:"required,number,len=2"`
	Year       string `validate:"required,number"`
	CardNumber string `validate:"required,len=13"`
	Filter     string `validate:"required,number,min=1"`
}

type ControlPlanParams struct {
	StartDate string `validate:"required,datetime=2006-01-02"`
	EndDate   string `validate:"required,datetime=2006-01-02"`
	Filter    string `validate:"required,number,min=1"`
}

type TControlPlan struct {
	TControlPlan interface{} `json:"t_suratkontrol"`
}

// UpdateControlPlans is used for updating both regular control plan and inpatient care order
// due to it's similarity.
type UpdateControlPlans struct {
	ControlLetterNumber      string `json:"noSuratKontrol"`
	SEPNumber                string `json:"noSEP"`
	InpatientCareOrderNumber string `json:"noSPRI,omitempty"`
	DoctorCode               string `json:"kodeDokter"`
	TargetPoliCode           string `json:"poliKontrol"`
	ControlPlannedDate       string `json:"tglRencanaKontrol"`
	User                     string `json:"user"`
}
