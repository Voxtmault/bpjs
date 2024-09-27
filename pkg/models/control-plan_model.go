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
