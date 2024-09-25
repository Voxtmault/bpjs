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

// BPJSSEP is used to marshall the SEPCreate struct into a JSON format that is accepted by the BPJS API.
// It's weird, ik dude trust me, but what can i do :D
type BPJSSEP struct {
	Request *TSEP `json:"request"`
}

type TSEP struct {
	TSEP *SEPCreate `json:"t_sep"`
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
	ParticipantType string `json:"jenisPeserta"`
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
	TreatmentClass        string          `json:"klsRawat"`
	ReferenceNumber       string          `json:"noRujukan"`
	SEPNumber             string          `json:"noSEP"`
	Guarantor             string          `json:"penjamin"`
	Participant           BPJSParticipant `json:"peserta"`
	Policlinics           string          `json:"poli"`
	PoliclinicExecutive   string          `json:"poliEksekutif"`
	SEPDate               string          `json:"tglSEP"`
	VisitationPurpose     string          `json:"tujuanKunj"`
}
