package models

type ObatPRBReuest struct {
	KdObat  string `json:"kdObat"`
	Signa1  string `json:"signa1"`
	Signa2  string `json:"signa2"`
	JmlObat string `json:"jmlObat"`
}

type PRBInsertRequest struct {
	NoSep      string          `json:"noSep"`
	NoKartu    string          `json:"noKartu"`
	Alamat     string          `json:"alamat"`
	Email      string          `json:"email"`
	ProgramPRB string          `json:"programPRB"`
	KodeDPJP   string          `json:"kodeDPJP"`
	Keterangan string          `json:"keterangan"`
	Saran      string          `json:"saran"`
	User       string          `json:"user"`
	Obat       []ObatPRBReuest `json:"obat"`
}

type PRBUpdateRequest struct {
	NoSRB      string          `json:"noSrb"`
	NoSep      string          `json:"noSep"`
	NoKartu    string          `json:"noKartu"`
	Alamat     string          `json:"alamat"`
	Email      string          `json:"email"`
	ProgramPRB string          `json:"programPRB"`
	KodeDPJP   string          `json:"kodeDPJP"`
	Keterangan string          `json:"keterangan"`
	Saran      string          `json:"saran"`
	User       string          `json:"user"`
	Obat       []ObatPRBReuest `json:"obat"`
}

type PRBDeleteRequest struct {
	NoSRB string `json:"noSrb"`
	NoSep string `json:"noSep"`
	User  string `json:"user"`
}

type TPRB struct {
	TPrb interface{} `json:"t_prb"`
}

type DPJPPRBResponse struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type ObatItemPRBResponse struct {
	JmlObat string `json:"jmlObat"`
	NmObat  string `json:"nmObat"`
	Signa   string `json:"signa"`
}

type ObatPRBResponse struct {
	List []ObatItemPRBResponse `json:"list"`
}

type AsalFaskesPRBResponse struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type PesertaPRBResponse struct {
	Alamat     string                `json:"alamat"`
	AsalFaskes AsalFaskesPRBResponse `json:"asalFaskes"`
	Email      string                `json:"email"`
	Kelamin    string                `json:"kelamin"`
	Nama       string                `json:"nama"`
	NoKartu    string                `json:"noKartu"`
	NoTelepon  string                `json:"noTelepon"`
	TglLahir   string                `json:"tglLahir"`
}

type PRBResponse struct {
	DPJP       DPJPPRBResponse    `json:"DPJP"`
	Keterangan string             `json:"keterangan"`
	NoSRB      string             `json:"noSRB"`
	Obat       ObatPRBResponse    `json:"obat"`
	Peserta    PesertaPRBResponse `json:"peserta"`
	ProgramPRB string             `json:"programPRB"`
	Saran      string             `json:"saran"`
	TglSRB     string             `json:"tglSRB"`
}

type DPJPPRBResponseGet struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type ObatItemPRBResponseGet struct {
	JmlObat string `json:"jmlObat"`
	KdObat  string `json:"kdObat"`
	NmObat  string `json:"nmObat"`
	Signa1  string `json:"signa1"`
	Signa2  string `json:"signa2"`
}

type ObatPRBResponseGet struct {
	Obat []ObatItemPRBResponseGet `json:"obat"`
}

type AsalFaskesPRBResponseGet struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type PesertaPRBResponseGet struct {
	Alamat     string                   `json:"alamat"`
	AsalFaskes AsalFaskesPRBResponseGet `json:"asalFaskes"`
	Email      string                   `json:"email"`
	Kelamin    string                   `json:"kelamin"`
	Nama       string                   `json:"nama"`
	NoKartu    string                   `json:"noKartu"`
	NoTelepon  string                   `json:"noTelepon"`
	TglLahir   string                   `json:"tglLahir"`
}

type ProgramPRBPRBResponseGet struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type PRBResponseGet struct {
	DPJP       DPJPPRBResponseGet       `json:"DPJP"`
	NoSEP      string                   `json:"noSEP"`
	NoSRB      string                   `json:"noSRB"`
	Obat       ObatPRBResponseGet       `json:"obat"`
	Peserta    PesertaPRBResponseGet    `json:"peserta"`
	ProgramPRB ProgramPRBPRBResponseGet `json:"programPRB"`
	Keterangan string                   `json:"keterangan"`
	Saran      string                   `json:"saran"`
	TglSRB     string                   `json:"tglSRB"`
}

type PRBResponseGetList struct {
	List []*PRBResponseGet `json:"list"`
}

type PRBResponseGetByNoSRB struct {
	Prb PRBResponseGet `json:"prb"`
}

type PRBResponseGetByTanggal struct {
	Prb PRBResponseGetList `json:"prb"`
}
