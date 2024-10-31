package models

type SelfReferralGet struct {
	ReferralNumber string        `json:"referral_number"`
	Diagnosis      SelfReference `json:"diagnosis"`
	TargetPoly     SelfReference `json:"target_poly"`
	Complaint      string        `json:"complaint"`
	ServiceType    SelfReference `json:"service_type"`
	Referrer       SelfReference `json:"referrer"`
	SEPSource      string        `json:"sep_source"`
	SEPSourceDate  string        `json:"sep_source_date"`
}

func (s *SelfReferralGet) FromBPJS(obj *Referral) {
	s.Diagnosis.Code = obj.Diagnosis.Code
	s.Diagnosis.Name = obj.Diagnosis.Name
	s.TargetPoly.Code = obj.PoliReferral.Code
	s.TargetPoly.Name = obj.PoliReferral.Name
	s.Complaint = obj.Complaint
	s.ServiceType.Code = obj.Service.Code
	s.ServiceType.Name = obj.Service.Name
	s.Referrer.Code = obj.Referrer.Code
	s.Referrer.Name = obj.Referrer.Name
	s.SEPSource = obj.BPJSEncounterID
	s.SEPSourceDate = obj.EncounterDate
}
