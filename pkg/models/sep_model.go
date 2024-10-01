package models

type TreatmentClass struct {
	TreatmentClassRights  string `json:"klsRawatHak"`
	TreatmentClassUpgrade string `json:"klsRawatNaik"`
	Financing             string `json:"pembiayaan"`
	PIC                   string `json:"penanggungJawab"`
}

type SEPReference struct {
	SourceReference          string `json:"asalRujukan"`
	ReferenceDate            string `json:"tglRujukan"`
	RefernceNumber           string `json:"noRujukan"`
	ReferencedHealthFacility string `json:"ppkRujukan"`
}

type SEPPoliclinics struct {
	PoliclinicCode string `json:"tujuan"` // Policlinics Code from BPJS
	Executive      string `json:"eksekutif"`
}

type SEPCOB struct {
	COB string `json:"cob"`
}

type SEPCatharacs struct {
	Catharacs string `json:"katarak"`
}

type Guarantee struct {
	Accident  string    `json:"lakaLantas"`
	LPNumber  string    `json:"noLP"`
	Guarantor Guarantor `json:"penjamin"`
}

type Guarantor struct {
	IncidentDate string     `json:"tglKejadian"`
	Note         string     `json:"keterangan"`
	Suppletion   Suppletion `json:"suplesi"`
}

type Suppletion struct {
	Suppletion       string           `json:"suplesi"`
	SuppletionNumber string           `json:"noSepSuplesi"`
	AccidentLocation AccidentLocation `json:"lokasiLaka"`
}

type AccidentLocation struct {
	ProvinceCode string `json:"kdPropinsi"`
	RegencyCode  string `json:"kdKabupaten"`
	DistrictCode string `json:"kdKecamatan"`

	// Everything Bellow is used for parsing SEP Accident Locations
	Note         string `json:"ketKejadian,omitempty"`
	Location     string `json:"lokasi,omitempty"`
	AccidentDate string `json:"tglKejadian,omitempty"`
}

type SKDP struct {
	LetterNumber           string `json:"noSurat"`
	AttendingPhysicianCode string `json:"kodeDPJP"`
}

// SEPCreate is used to create a new SEP number from BPJS, Wrap this inside a variable named t_sep
// and then wrap the t_sep inside a variable named request. I know it's weird but what can we do :D
type SEPCreate struct {
	BPJSID                string         `json:"noKartu"`
	ServiceDate           string         `json:"tglSep"`
	HealthFacilityCode    string         `json:"ppkPelayanan"`
	ServiceType           string         `json:"jnsPelayanan"`
	TreatmentClass        TreatmentClass `json:"klsRawat"`
	MRNumber              string         `json:"noMR"`
	Reference             SEPReference   `json:"rujukan"`
	Note                  string         `json:"catatan"`
	InitialDiagnosis      string         `json:"diagAwal"`
	Policlinics           SEPPoliclinics `json:"poli"`
	COB                   SEPCOB         `json:"cob"`
	Catharacs             SEPCatharacs   `json:"katarak"`
	Guarantee             Guarantee      `json:"jaminan"`
	VisitationPurpose     string         `json:"tujuanKunj"`
	ProcedureFlag         string         `json:"flagProcedure"`
	HealthCareSupportCode string         `json:"kdPenunjang"`
	ServiceAssessment     string         `json:"assesmentPel"`
	SKDP                  SKDP           `json:"skdp"`
	ServiceDPJP           string         `json:"dpjpLayan"`
	PhoneNum              string         `json:"noTelp"`
	User                  string         `json:"user"`
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
	PoliclinicsCode       string          `json:"kdPoli"`
	TreatmentClass        string          `json:"kelasRawat"`
	ReferenceNumber       string          `json:"noRujukan"`
	SEPNumber             string          `json:"noSep"`
	Guarantor             string          `json:"penjamin"`
	Participant           BPJSParticipant `json:"peserta"`
	Policlinics           string          `json:"poli"`
	PoliclinicExecutive   string          `json:"poliEksekutif"`
	SEPDate               string          `json:"tglSep"`
	VisitationPurpose     string          `json:"tujuanKunj"`
}

type SEPUpdate struct {
	SEPNumber        string         `json:"noSep"`
	TreatmentClass   TreatmentClass `json:"klsRawat"`
	MRNumber         string         `json:"noMR"`
	Note             string         `json:"catatan"`
	InitialDiagnosis string         `json:"diagAwal"`
	Policlinics      SEPPoliclinics `json:"poli"`
	COB              SEPCOB         `json:"cob"`
	Catharacs        SEPCatharacs   `json:"katarak"`
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
	ReferalNumber      string                 `json:"noRujukan"`
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
	Catharacts         string                 `json:"katarak"`
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
