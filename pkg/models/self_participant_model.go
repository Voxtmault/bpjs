package models

type SelfParticipant struct {
	ParticipantName   string         `json:"participant_name"`
	MRNumber          string         `json:"mr_number"`
	CardNumber        string         `json:"card_number"`
	SSN               string         `json:"ssn"`
	Gender            string         `json:"gender"`
	DOB               string         `json:"dob"`
	Age               string         `json:"age"`
	NursingClass      *SelfReference `json:"nursing_class"`
	ParticipantStatus *SelfReference `json:"participant_status"`
	ParticipantType   *SelfReference `json:"participant_type"`
	Provider          *SelfReference `json:"provider"`
}

func (s *SelfParticipant) FromBPJS(obj *BPJSParticipant) {
	s.ParticipantName = obj.Name
	s.MRNumber = obj.MedicalRecord.MRNumber
	s.CardNumber = obj.CardNumber
	s.SSN = obj.NIK
	s.Gender = obj.Sex
	s.DOB = obj.DOB
	s.Age = obj.Age.AgeNow

	// Null Check Safety
	if obj.ClassRights != nil {
		s.NursingClass = &SelfReference{
			Code: obj.ClassRights.Code,
			Name: obj.ClassRights.Note,
		}
	}

	if obj.ParticipantStatus != nil {
		s.ParticipantStatus = &SelfReference{
			Code: obj.ParticipantStatus.Code,
			Name: obj.ParticipantStatus.Note,
		}
	}

	if obj.ParticipantType != nil {
		s.ParticipantType = &SelfReference{
			Code: obj.ParticipantType.Code,
			Name: obj.ParticipantType.Note,
		}
	}

	if obj.ProvUmum != nil {
		s.Provider = &SelfReference{
			Code: obj.ProvUmum.ProviderCode,
			Name: obj.ProvUmum.ProviderName,
		}
	}
}
