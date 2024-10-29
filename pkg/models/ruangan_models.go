package models

type ReferenceKamar struct {
	KodeKelas string `json:"kodekelas"`
	NamaKelas string `json:"namakelas"`
}

type AplicaresRequestResponse struct {
	Lists []*ReferenceKamar `json:"list"`
}
