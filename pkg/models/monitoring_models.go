package models

type MonitoringDataKunjungan struct {
	Diagnosa     string  `json:"diagnosa"`
	JnsPelayanan string  `json:"jnsPelayanan"`
	KelasRawat   string  `json:"kelasRawat"`
	Nama         string  `json:"nama"`
	NoKartu      string  `json:"noKartu"`
	NoSep        string  `json:"noSep"`
	NoRujukan    string  `json:"noRujukan"`
	Poli         *string `json:"poli"` // Using *string to allow null values
	TglPlgSep    string  `json:"tglPlgSep"`
	TglSep       string  `json:"tglSep"`
}

type MonitoringDataKunjunganResponse struct {
	Sep []*MonitoringDataKunjungan `json:"sep"`
}
type GetMonitoringHistoryPelayananPeserta struct {
	Sep []*MonitoringDataKunjungan `json:"histori"`
}

type InacbgKlaim struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type BiayaKlaim struct {
	ByPengajuan   string `json:"byPengajuan"`
	BySetujui     string `json:"bySetujui"`
	ByTarifGruper string `json:"byTarifGruper"`
	ByTarifRS     string `json:"byTarifRS"`
	ByTopup       string `json:"byTopup"`
}

type PesertaKlaim struct {
	Nama    string `json:"nama"`
	NoKartu string `json:"noKartu"`
	NoMR    string `json:"noMR"`
}

type DataKlaim struct {
	Inacbg     InacbgKlaim  `json:"Inacbg"`
	Biaya      BiayaKlaim   `json:"biaya"`
	KelasRawat string       `json:"kelasRawat"`
	NoFPK      string       `json:"noFPK"`
	NoSEP      string       `json:"noSEP"`
	Peserta    PesertaKlaim `json:"peserta"`
	Poli       string       `json:"poli"`
	Status     string       `json:"status"`
	TglPulang  string       `json:"tglPulang"`
	TglSep     string       `json:"tglSep"`
}

type MonitoringKlaimResponse struct {
	Klaim []*DataKlaim `json:"klaim"`
}

type PesertaJasaraharja struct {
	NoKartu string `json:"noKartu"`
	Nama    string `json:"nama"`
	NoMR    string `json:"noMR"`
}

type SEPJasaraharja struct {
	NoSEP        string             `json:"noSEP"`
	TglSEP       string             `json:"tglSEP"`
	TglPlgSEP    string             `json:"tglPlgSEP"`
	NoMR         string             `json:"noMr"`
	JnsPelayanan string             `json:"jnsPelayanan"`
	Poli         string             `json:"poli"`
	Diagnosa     string             `json:"diagnosa"`
	Peserta      PesertaJasaraharja `json:"peserta"`
}

type JasaRaharja struct {
	TglKejadian        string `json:"tglKejadian"`
	NoRegister         string `json:"noRegister"`
	KetStatusDijamin   string `json:"ketStatusDijamin"`
	KetStatusDikirim   string `json:"ketStatusDikirim"`
	BiayaDijamin       string `json:"biayaDijamin"`
	Plafon             string `json:"plafon"`
	JmlDibayar         string `json:"jmlDibayar"`
	ResultsJasaRaharja string `json:"resultsJasaRaharja"`
}

type DataJasaraharja struct {
	SEP         SEPJasaraharja `json:"sep"`
	JasaRaharja JasaRaharja    `json:"jasaRaharja"`
}

type ResponseKlaimJaminanJasaraharja struct {
	Data []*DataJasaraharja `json:"data"`
}
