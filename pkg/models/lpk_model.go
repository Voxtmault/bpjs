package models

type PoliLPKRequest struct {
	Poli string `json:"poli"`
}

type PerawatanLPKRequest struct {
	RuangRawat    string `json:"ruangRawat"`
	KelasRawat    string `json:"kelasRawat"`
	Spesialistik  string `json:"spesialistik"`
	CaraKeluar    string `json:"caraKeluar"`
	KondisiPulang string `json:"kondisiPulang"`
}

type DiagnosaLPKRequest struct {
	Kode  string `json:"kode"`
	Level string `json:"level"`
}

type ProcedureLPKRequest struct {
	Kode string `json:"kode"`
}

type DirujukKe struct {
	KodePPK string `json:"kodePPK"`
}

type KontrolKembali struct {
	TglKontrol string `json:"tglKontrol"`
	Poli       string `json:"poli"`
}

type RencanaTL struct {
	TindakLanjut   string         `json:"tindakLanjut"`
	DirujukKe      DirujukKe      `json:"dirujukKe"`
	KontrolKembali KontrolKembali `json:"kontrolKembali"`
}

type InsertLPKRequest struct {
	NoSep     string                `json:"noSep"`
	TglMasuk  string                `json:"tglMasuk"`
	TglKeluar string                `json:"tglKeluar"`
	Jaminan   string                `json:"jaminan"`
	Poli      PoliLPKRequest        `json:"poli"`
	Perawatan PerawatanLPKRequest   `json:"perawatan"`
	Diagnosa  []DiagnosaLPKRequest  `json:"diagnosa"`
	Procedure []ProcedureLPKRequest `json:"procedure"`
	RencanaTL RencanaTL             `json:"rencanaTL"`
	DPJP      string                `json:"DPJP"`
	User      string                `json:"user"`
}

type Tlpk struct {
	TLpk interface{} `json:"t_lpk"`
}

type DokterLPKResponse struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type DPJPLPKResponse struct {
	Dokter DokterLPKResponse `json:"dokter"`
}

type ListDiagnosaLPKResponse struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type DiagnosaItemLPKResponse struct {
	Level string                  `json:"level"`
	List  ListDiagnosaLPKResponse `json:"list"`
}

type DiagnosaLPKResponse struct {
	List []DiagnosaItemLPKResponse `json:"list"`
}

type PerawatanDetailLPKResponse struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type PerawatanLPKResponse struct {
	CaraKeluar    PerawatanDetailLPKResponse `json:"caraKeluar"`
	KelasRawat    PerawatanDetailLPKResponse `json:"kelasRawat"`
	KondisiPulang PerawatanDetailLPKResponse `json:"kondisiPulang"`
	RuangRawat    PerawatanDetailLPKResponse `json:"ruangRawat"`
	Spesialistik  PerawatanDetailLPKResponse `json:"spesialistik"`
}

type PesertaLPKResponse struct {
	Kelamin  string `json:"kelamin"`
	Nama     string `json:"nama"`
	NoKartu  string `json:"noKartu"`
	NoMR     string `json:"noMR"`
	TglLahir string `json:"tglLahir"`
}

type PoliDetailLPKResponse struct {
	Kode string `json:"kode"`
}

type PoliLPKResponse struct {
	Eksekutif string                `json:"eksekutif"`
	Poli      PoliDetailLPKResponse `json:"poli"`
}

type ProcedureItemLPKResponse struct {
	List ListDiagnosaLPKResponse `json:"list"`
}

type ProcedureLPKResponse struct {
	List []ProcedureItemLPKResponse `json:"list"`
}

type ListLPKResponse struct {
	DPJP         DPJPLPKResponse      `json:"DPJP"`
	Diagnosa     DiagnosaLPKResponse  `json:"diagnosa"`
	JnsPelayanan string               `json:"jnsPelayanan"`
	NoSep        string               `json:"noSep"`
	Perawatan    PerawatanLPKResponse `json:"perawatan"`
	Peserta      PesertaLPKResponse   `json:"peserta"`
	Poli         PoliLPKResponse      `json:"poli"`
	Procedure    ProcedureLPKResponse `json:"procedure"`
	RencanaTL    *RencanaTL           `json:"rencanaTL"` // Menggunakan `interface{}` untuk mengizinkan `null` value
	TglKeluar    string               `json:"tglKeluar"`
	TglMasuk     string               `json:"tglMasuk"`
}

type LPKResponseData struct {
	Lpk LPKResponseDataList `json:"lpk"`
}
type LPKResponseDataList struct {
	List []*ListLPKResponse `json:"list"`
}
