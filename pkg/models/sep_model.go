package models

type TreatmentClass struct {
	TreatmentClassRights  string `json:"klsRawatHak" validate:"required,number"`
	TreatmentClassUpgrade string `json:"klsRawatNaik" validate:"omitempty,number"`
	Financing             string `json:"pembiayaan" validate:"required_with=TreatmentClassUpgrade,len=0|number"`
	PIC                   string `json:"penanggungJawab" validate:"required_with=TreatmentClassUpgrade,len=0|number"`
}

type SEPReference struct {
	SourceReference          string `json:"asalRujukan" validate:"required,number"`
	ReferenceDate            string `json:"tglRujukan" validate:"required,datetime=2006-01-02"`
	ReferenceNumber          string `json:"noRujukan" validate:"omitempty"`
	ReferencedHealthFacility string `json:"ppkRujukan" validate:"required"`
}

type SEPPolyclinics struct {
	PoliclinicCode string `json:"tujuan" validate:"required,alphanum"` // Polyclinics Code from BPJS
	Executive      string `json:"eksekutif" validate:"required,number"`
}

type SEPCOB struct {
	COB string `json:"cob" validate:"required,number"`
}

type SEPCataract struct {
	Cataract string `json:"katarak" validate:"required,number"`
}

type Guarantee struct {
	Accident  string    `json:"lakaLantas" validate:"required,number"`
	LPNumber  string    `json:"noLP" validate:"omitempty"`
	Guarantor Guarantor `json:"penjamin" validate:"required"`
}

type Guarantor struct {
	IncidentDate string     `json:"tglKejadian" validate:"omitempty,datetime=2006-01-02"`
	Note         string     `json:"keterangan" validate:"omitempty"`
	Suppletion   Suppletion `json:"suplesi" validate:"required"`
}

type Suppletion struct {
	Suppletion       string           `json:"suplesi" validate:"required,number"`
	SuppletionNumber string           `json:"noSepSuplesi" validate:"omitempty"`
	AccidentLocation AccidentLocation `json:"lokasiLaka" validate:"required"`
}

type AccidentLocation struct {
	ProvinceCode string `json:"kdPropinsi" validate:"omitempty"`
	RegencyCode  string `json:"kdKabupaten" validate:"omitempty"`
	DistrictCode string `json:"kdKecamatan" validate:"omitempty"`

	// Everything Bellow is used for parsing SEP Accident Locations
	Note         string `json:"ketKejadian,omitempty"`
	Location     string `json:"lokasi,omitempty"`
	AccidentDate string `json:"tglKejadian,omitempty"`
}

type SKDP struct {
	LetterNumber           string `json:"noSurat" validate:"omitempty"`
	AttendingPhysicianCode string `json:"kodeDPJP" validate:"omitempty"`
}

// SEPCreate is used to create a new SEP number from BPJS, Wrap this inside a variable named t_sep
// and then wrap the t_sep inside a variable named request. I know it's weird but what can we do :D
type SEPCreate struct {
	BPJSID                string         `json:"noKartu" validate:"required,len=13,number"`
	ServiceDate           string         `json:"tglSep" validate:"required,datetime=2006-01-02"`
	HealthFacilityCode    string         `json:"ppkPelayanan" validate:"required"`
	ServiceType           string         `json:"jnsPelayanan" validate:"required,number,len=1"`
	TreatmentClass        TreatmentClass `json:"klsRawat" validate:"required"`
	MRNumber              string         `json:"noMR" validate:"required"`
	Reference             SEPReference   `json:"rujukan" validate:"required"`
	Note                  string         `json:"catatan"`
	InitialDiagnosis      string         `json:"diagAwal" validate:"required"`
	Polyclinics           SEPPolyclinics `json:"poli" validate:"required"`
	COB                   SEPCOB         `json:"cob" validate:"required"`
	Cataracts             SEPCataract    `json:"katarak" validate:"required"`
	Guarantee             Guarantee      `json:"jaminan" validate:"required"`
	VisitationPurpose     string         `json:"tujuanKunj" validate:"required,number"`
	ProcedureFlag         string         `json:"flagProcedure" validate:"required_unless=VisitationPurpose 0"`
	HealthCareSupportCode string         `json:"kdPenunjang" validate:"required_unless=VisitationPurpose 0"`
	ServiceAssessment     string         `json:"assesmentPel" validate:"omitempty"`
	SKDP                  SKDP           `json:"skdp" validate:"required"`
	ServiceDPJP           string         `json:"dpjpLayan" validate:"required_unless=ServiceType 1"`
	PhoneNum              string         `json:"noTelp" validate:"omitempty"`
	User                  string         `json:"user" validate:"required"`
}

type TSEP struct {
	TSEP interface{} `json:"t_sep"`
}

type SEPInformation struct {
	Dinsos      string `json:"dinsos"`
	ESEP        string `json:"eSEP"`
	SKTMNumber  string `json:"noSKTM"`
	ProlanisPRB string `json:"prolanisPRB"`
}

// SEPCreateResponse is... well.. let's just say it is there and very different from the BPJS Participant
// Don't ask me why, I don't know either
type SEPParticipantResponse struct {
	Insurance       string `json:"asuransi"`
	ClassRights     string `json:"hakKelas"`
	ParticipantType string `json:"jnsPeserta"`
	Sex             string `json:"kelamin"`
	Name            string `json:"nama"`
	CardNumber      string `json:"noKartu"`
	MRNumber        string `json:"noMr"`
	DOB             string `json:"tglLahir"`
}

type SEPCreateResponse struct {
	ServiceAssessment     string          `json:"assesmentPel"`
	Note                  string          `json:"catatan"`
	Diagnosis             string          `json:"diagnosa"`
	ProcedureFlag         string          `json:"flagProcedure"`
	Information           Information     `json:"informasi"`
	ServiceType           string          `json:"jnsPelayanan"`
	HealthCareSupportCode string          `json:"kdPenunjang"`
	PolyclinicsCode       string          `json:"kdPoli"`
	TreatmentClass        string          `json:"kelasRawat"`
	ReferenceNumber       string          `json:"noRujukan"`
	SEPNumber             string          `json:"noSep"`
	Guarantor             string          `json:"penjamin"`
	Participant           BPJSParticipant `json:"peserta"`
	Polyclinics           string          `json:"poli"`
	PoliclinicExecutive   string          `json:"poliEksekutif"`
	SEPDate               string          `json:"tglSep"`
	VisitationPurpose     string          `json:"tujuanKunj"`
}
type SEPCreateResponseWrapper struct {
	SEP *SEPCreateResponse `json:"sep"`
}

type SEPUpdate struct {
	SEPNumber        string         `json:"noSep"`
	TreatmentClass   TreatmentClass `json:"klsRawat"`
	MRNumber         string         `json:"noMR"`
	Note             string         `json:"catatan"`
	InitialDiagnosis string         `json:"diagAwal"`
	Polyclinics      SEPPolyclinics `json:"poli"`
	COB              SEPCOB         `json:"cob"`
	Cataracts        SEPCataract    `json:"katarak"`
	Guarantee        Guarantee      `json:"jaminan"`
	ServiceDPJP      string         `json:"dpjpLayan"`
	PhoneNum         string         `json:"noTelp"`
	User             string         `json:"user"`
}

type SEPDelete struct {
	SEPNumber string `json:"noSep"`
	User      string `json:"user"`
}

type DPJP struct {
	Code string `json:"kdDPJP"`
	Name string `json:"nmDPJP"`
}

type SEPControl struct {
	DoctorCode          string `json:"kdDokter"`
	DoctorName          string `json:"nmDokter"`
	ControlLetterNumber string `json:"noSurat"`
}

type SEPGet struct {
	SEPNumber          string                 `json:"noSep"`
	SEPDate            string                 `json:"tglSep"`
	ServiceType        string                 `json:"jnsPelayanan"`
	NursingClass       string                 `json:"kelasRawat"`
	Diagnosis          string                 `json:"diagnosa"`
	ReferralNumber     string                 `json:"noRujukan"`
	Policlinic         string                 `json:"poli"`
	PoliExecutive      string                 `json:"poliEksekutif"`
	Note               string                 `json:"catatan"`
	Guarantor          Guarantor              `json:"penjamin"`
	AccidentStatusCode string                 `json:"kdStatusKecelakaan"`
	AccidentStatus     string                 `json:"nmstatusKecelakaan"`
	AccidentLocation   AccidentLocation       `json:"lokasiKecelakaan"`
	DPJP               DPJP                   `json:"dpjp"`
	Participant        SEPParticipantResponse `json:"peserta"`
	TreatmentClass     TreatmentClass         `json:"klsRawat"`
	Control            SEPControl             `json:"kontrol"`
	COB                string                 `json:"cob"`
	Cataracts          string                 `json:"katarak"`
	VisitationPurpose  ReusableNote           `json:"tujuanKunj"`
	FlagProcedure      ReusableNote           `json:"flagProcedure"`
	HealthCareCode     ReusableNote           `json:"kdPenunjang"`
	ServiceAssessment  ReusableNote           `json:"assestmenPel"`
	DigitalSEP         string                 `json:"eSEP"`
}

// SEPRequestCreate is used for creating SEP Backdate and Fingerprint Request to BPJS
type SEPRequestCreate struct {
	CardNumber  string `json:"noKartu"`
	SEPDate     string `json:"tglSep"`
	ServiceType string `json:"jnsPelayanan"`
	RequestType string `json:"jnsPengajuan"`
	Note        string `json:"keterangan"`
	User        string `json:"user"`
}

type SEPRequest struct {
	CardNumber      string `json:"noKartu"`
	ParticipantName string `json:"nama"`
	SEPDate         string `json:"tglSep"`
	ServiceType     string `json:"jnsPelayanan"`
	Approval        string `json:"persetujuan"`
	Status          string `json:"status"`
}
type SEPRequestResponse struct {
	Lists []*SEPRequest `json:"list"`
}

type SEPUpdateTanggalPulangRequest struct {
	NoSep            string `json:"noSep"`
	StatusPulang     string `json:"statusPulang"`
	NoSuratMeninggal string `json:"noSuratMeninggal"`
	TglMeninggal     string `json:"tglMeninggal"`
	TglPulang        string `json:"tglPulang"`
	NoLPManual       string `json:"noLPManual"`
	User             string `json:"user"`
}

type DischargedSEP struct {
	SEPNumber          string `json:"noSep"`
	UpdatingSEPNumber  string `json:"noSepUpdating"`
	ServiceType        string `json:"jnsPelayanan"`
	HealthFacilityCode string `json:"ppkTujuan"`
	CardNumber         string `json:"noKartu"`
	ParticipantName    string `json:"nama"`
	ServiceDate        string `json:"tglSep"`
	DischargedDate     string `json:"tglPulang"`
	Status             string `json:"status"`
	DateOfDeath        string `json:"tglMeninggal"`
	LetterNumber       string `json:"noSurat"`
	Note               string `json:"keterangan"`
	User               string `json:"user"`
}
type DischargedSEPWrapper struct {
	List []*DischargedSEP `json:"list"`
}

type SEPFingerPrintResponse struct {
	Lists []*SEPGetListFingerPrint `json:"list"`
}

type SEPRandomQuestionResponse struct {
	Lists []*SEPGetRandomQuestion `json:"faskes"`
}

type SEPGetFingerPrint struct {
	Kode   string `json:"kode"`
	Status string `json:"status"`
}

type SEPGetRandomQuestion struct {
	Kode   string `json:"kode"`
	Status string `json:"nama"`
}

type SEPGetListFingerPrint struct {
	NoKartu string `json:"noKartu"`
	NoSEP   string `json:"noSEP"`
}

type PostRequestRandomQuestion struct {
	NoKartu   string `json:"noKartu"`
	TglSep    string `json:"tglSep"`
	JenPel    string `json:"jenPel"`
	PpkPelSep string `json:"ppkPelSep"`
	TglLahir  string `json:"tglLahir"`
	PpkPst    string `json:"ppkPst"`
	User      string `json:"user"`
}

type InternalSEP struct {
	// To be honest, i feel like the first 2 are the same variables, it's just that i feel like
	// BPJS forgot that these exists and decided to just eye ball it. Man talk about spaghetti code
	// And please for the love of god, use a consistent json naming scheme will you, almost everything is using camelCase
	// and now, what is this? snake_case? what is this? python?
	TargetReferralPolyclinic     string `json:"tujuanrujuk"`
	TargetReferralPolyclinicCode string `json:"kdpolituj"`
	TargetReferralPolyclinicName string `json:"namatujuanrujuk"`
	SourcePolyName               string `json:"nmpoliasal"`
	SourcePolyCode               string `json:"kdpoliasal"`
	InternalServiceDate          string `json:"tglrujukinternal"`
	SEPNumber                    string `json:"nosep"`
	ReferredSEPNumber            string `json:"nosepref"`
	HealthFacilityCode           string `json:"ppkpelsep"`
	NoKapSt                      string `json:"nokapst"`
	ServiceDate                  string `json:"tglsep"`
	LetterNumber                 string `json:"nosurat"`
	InternalFlag                 string `json:"flaginternal"`
	SupportCode                  string `json:"kdpenunjang"`
	SupportName                  string `json:"nmpenunjang"`
	Diagnosis                    string `json:"diagppk"`
	DiagnosisName                string `json:"nmdiag"`
	DoctorCode                   string `json:"kodedokter"`
	DoctorName                   string `json:"nmdokter"`
	FlagProcedure                string `json:"flagprocedure"`
	ConsultationOption           string `json:"opsikonsul"`
	SEPFlag                      string `json:"flagsep"`
	FUser                        string `json:"fuser"`
	FDate                        string `json:"fdate"`
}
type InternalSEPWrapper struct {
	List  []*InternalSEP `json:"list"`
	Count string         `json:"count"`
}
