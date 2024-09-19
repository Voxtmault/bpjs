package models

import (
	pb "github.com/voxtmault/bpjs-service-proto/go"
)

type COB struct {
	InsuranceName   string `json:"nmAsuransi"`
	InsuranceNumber string `json:"noAsuransi"`
	TATDate         string `json:"tglTAT"`
	TMTDate         string `json:"tglTMT"`
}

type Information struct {
	Dinsos      string `json:"dinsos"`
	SKTMNumber  string `json:"noSKTM"`
	ProlanisPRB string `json:"prolanisPRB"`
}

type BPJSMedicalRecord struct {
	MRNumber    string `json:"noMR"`
	PhoneNumber string `json:"noTelepon"`
}

type ProvUmum struct {
	ProviderCode string `json:"kdProvider"`
	ProviderName string `json:"nMProvider"`
}

type PatientAge struct {
	AgeNow       string `json:"umurSekarang"`
	AgeAtService string `json:"umurSaatPelayanan"`
}

type ClassRights struct {
	ReusableNote
}
type ParticipantType struct {
	ReusableNote
}
type ParticipantStatus struct {
	ReusableNote
}

type BPJSParticipant struct {
	Name              string            `json:"nama"`
	NIK               string            `json:"nik"`
	CardNumber        string            `json:"noKartu"`
	Pisa              string            `json:"pisa"`
	Sex               string            `json:"sex"`
	TATDate           string            `json:"tglTAT"`
	TMTDate           string            `json:"tglTMT"`
	DOB               string            `json:"tglLahir"`
	CardPrintDate     string            `json:"tglCetakKartu"`
	COB               COB               `json:"cob"`
	ClassRights       ClassRights       `json:"hakKelas"`
	Information       Information       `json:"informasi"`
	ParticipantType   ParticipantType   `json:"jenisPeserta"`
	MedicalRecord     BPJSMedicalRecord `json:"mr"`
	ProvUmum          ProvUmum          `json:"provUmum"`
	ParticipantStatus ParticipantStatus `json:"statusPeserta"`
	Age               PatientAge        `json:"umur"`
}

type BPJSParticipantResponse struct {
	Participant *BPJSParticipant `json:"peserta"`
}

// Reusable Class
type ReusableNote struct {
	Code string `json:"kode"`
	Note string `json:"keterangan"`
}

// Search Params
type ParticipantSearchParams struct {
	NIK         string `validate:"omitempty,numeric,min=16" example:"1234567890123456"`
	BPJSNumber  string `validate:"omitempty,numeric,min=13" example:"1234567890123"`
	ServiceDate string `validate:"omitempty,datetime=2006-01-02" example:"2021-01-01"`
}

func (p *BPJSParticipant) ToProto() *pb.BPJSParticipant {
	return &pb.BPJSParticipant{
		Name:          p.Name,
		Nik:           p.NIK,
		CardNumber:    p.CardNumber,
		Pisa:          p.Pisa,
		Sex:           p.Sex,
		TatDate:       p.TATDate,
		TmtDate:       p.TMTDate,
		Dob:           p.DOB,
		CardPrintDate: p.CardPrintDate,
		Cob: &pb.COB{
			InsuranceName:   p.COB.InsuranceName,
			InsuranceNumber: p.COB.InsuranceNumber,
			TatDate:         p.COB.TATDate,
			TmtDate:         p.COB.TMTDate,
		},
		ClassRights: &pb.ReusableNote{
			Code: p.ClassRights.Code,
			Note: p.ClassRights.Note,
		},
		Information: &pb.Information{
			Dinsos:      p.Information.Dinsos,
			SKTMNumber:  p.Information.SKTMNumber,
			ProlanisPRB: p.Information.ProlanisPRB,
		},
		ParticipantType: &pb.ReusableNote{
			Code: p.ParticipantType.Code,
			Note: p.ParticipantType.Note,
		},
		MedicalRecord: &pb.BPJSMedicalRecord{
			MRNumber:    p.MedicalRecord.MRNumber,
			PhoneNumber: p.MedicalRecord.PhoneNumber,
		},
		ProvUmum: &pb.ProvUmum{
			ProviderCode: p.ProvUmum.ProviderCode,
			ProviderName: p.ProvUmum.ProviderName,
		},
		ParticipantStatus: &pb.ReusableNote{
			Code: p.ParticipantStatus.Code,
			Note: p.ParticipantStatus.Note,
		},
		Age: &pb.PatientAge{
			AgeNow:       p.Age.AgeNow,
			AgeAtService: p.Age.AgeAtService,
		},
	}
}

func (p *BPJSParticipant) FromProto(obj *pb.BPJSParticipant) {
	p.Name = obj.GetName()
	p.NIK = obj.GetNik()
	p.CardNumber = obj.GetCardNumber()
	p.Pisa = obj.GetPisa()
	p.Sex = obj.GetSex()
	p.TATDate = obj.GetTatDate()
	p.TMTDate = obj.GetTmtDate()
	p.DOB = obj.GetDob()
	p.CardPrintDate = obj.GetCardPrintDate()
	p.COB = COB{
		InsuranceName:   obj.GetCob().GetInsuranceName(),
		InsuranceNumber: obj.GetCob().GetInsuranceNumber(),
		TATDate:         obj.GetCob().GetTatDate(),
		TMTDate:         obj.GetCob().GetTmtDate(),
	}
	p.ClassRights = ClassRights{
		ReusableNote: ReusableNote{
			Code: obj.GetClassRights().GetCode(),
			Note: obj.GetClassRights().GetNote(),
		},
	}
	p.Information = Information{
		Dinsos:      obj.GetInformation().GetDinsos(),
		SKTMNumber:  obj.GetInformation().GetSKTMNumber(),
		ProlanisPRB: obj.GetInformation().GetProlanisPRB(),
	}
	p.ParticipantType = ParticipantType{
		ReusableNote: ReusableNote{
			Code: obj.GetClassRights().GetCode(),
			Note: obj.GetClassRights().GetNote(),
		},
	}
	p.MedicalRecord = BPJSMedicalRecord{
		MRNumber:    obj.GetMedicalRecord().GetMRNumber(),
		PhoneNumber: obj.GetMedicalRecord().GetPhoneNumber(),
	}
	p.ProvUmum = ProvUmum{
		ProviderCode: obj.GetProvUmum().GetProviderCode(),
		ProviderName: obj.GetProvUmum().GetProviderName(),
	}
	p.ParticipantStatus = ParticipantStatus{
		ReusableNote: ReusableNote{
			Code: obj.GetClassRights().GetCode(),
			Note: obj.GetClassRights().GetNote(),
		},
	}
	p.Age = PatientAge{
		AgeNow:       obj.GetAge().GetAgeNow(),
		AgeAtService: obj.GetAge().GetAgeAtService(),
	}
}
