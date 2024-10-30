package models

type ReferenceJenisKamar struct {
	KodeKelas string `json:"kodekelas"`
	NamaKelas string `json:"namakelas"`
}

type AplicaresRequestResponse struct {
	Lists []*ReferenceJenisKamar `json:"list"`
}

type AplicaresRuanganResponse struct {
	Lists []*RuanganResponse `json:"list"`
}

type Ruangan struct {
	KodeKelas          string `json:"kodekelas"`
	KodeRuang          string `json:"koderuang"`
	NamaRuang          string `json:"namaruang"`
	Kapasitas          string `json:"kapasitas"`
	Tersedia           string `json:"tersedia"`
	TersediaPria       string `json:"tersediapria"`
	TersediaWanita     string `json:"tersediawanita"`
	TersediaPriaWanita string `json:"tersediapriawanita"`
}

type RuanganResponse struct {
	KodeKelas          string `json:"kodekelas"`
	KodeRuang          string `json:"koderuang"`
	NamaRuang          string `json:"namaruang"`
	Kapasitas          int    `json:"kapasitas"`
	Tersedia           int    `json:"tersedia"`
	TersediaPria       int    `json:"tersediapria"`
	TersediaWanita     int    `json:"tersediawanita"`
	TersediaPriaWanita int    `json:"tersediapriawanita"`
}
