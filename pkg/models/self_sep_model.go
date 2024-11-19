package models

type SelfSEPCreateReferral struct {
	ReferralSource             string `json:"referral_source" validate:"required,number"`
	ReferralDate               string `json:"referral_date" validate:"required,datetime=2006-01-02"`
	ReferralNumber             string `json:"referral_number" validate:"omitempty,numeric"`
	ReferrerHealthFacilityCode string `json:"referrer_health_facility_code" validate:"required,numeric"`
}

type SelfSEPCreateControl struct {
	ControllLetterNumber string `json:"controll_letter_number" validate:"omitempty,numeric"`
	ReferredDoctorCode   string `json:"referred_doctor_code" validate:"omitempty,numeric"`
}

type SelfSEPCreateGuarantee struct {
	AccidentType             string `json:"accident_type" validate:"required,number"`
	RegisteredAccidentNumber string `json:"registered_accident_number" validate:"omitempty"`
	Guarantor                struct {
		AccidentDate string `json:"accident_date" validate:"omitempty,datetime=2006-01-02"`
		Note         string `json:"note" validate:"omitempty"`
		Suppletion   struct {
			Suppletion       string `json:"suppletion" validate:"required,number"`
			SuppletionNumber string `json:"suppletion_number" validate:"omitempty"`
			AccidentLocation struct {
				ProvinceCode string `json:"province_code" validate:"omitempty,numeric"`
				RegencyCode  string `json:"regency_code" validate:"omitempty,numeric"`
				DistrictCode string `json:"district_code" validate:"omitempty,numeric"`
			} `json:"accident_location" validate:"required"`
		} `json:"suppletion" validate:"required"`
	} `json:"guarantor" validate:"required"`
}

type SelfSEPCreatePoli struct {
	PolyclinicCode string `json:"polyclinic_code" validate:"required"`
	Executive      string `json:"executive" validate:"required,number"`
}

type BPJSParticipantTreatmentClass struct {
	TreatmentClass        string `json:"treatment_class" validate:"required,number"`
	UpgradeTreatmentClass string `json:"upgrade_treatment_class" validate:"omitempty,number"`
	Financing             string `json:"financing" validate:"omitempty"`
	PIC                   string `json:"pic" validate:"omitempty"`
}

type BPJSAccounting struct {
	OperatorID    uint   `json:"operator_id" validate:"required,number,min=1"`
	OperatorName  string `json:"operator_name" validate:"required"`
	OperatorPhone string `json:"operator_phone" validate:"required"`
}

type SelfSEPCreate struct {
	CardNumber          string                         `json:"card_number" validate:"required,number,min=13"`
	ServiceDate         string                         `json:"service_date" validate:"required,datetime=2006-01-02"`
	ServiceType         string                         `json:"service_type" validate:"required,number"`
	MRNumber            string                         `json:"mr_number" validate:"required"`
	TreatmentClass      *BPJSParticipantTreatmentClass `json:"treatment_class" validate:"required"`
	Referral            *SelfSEPCreateReferral         `json:"referral" validate:"required"`
	Control             *SelfSEPCreateControl          `json:"control" validate:"required"`
	Guarantee           *SelfSEPCreateGuarantee        `json:"guarantee" validate:"required"`
	Note                string                         `json:"note" validate:"omitempty"`
	InitialDiagnosis    string                         `json:"initial_diagnosis" validate:"required"`
	Poli                *SelfSEPCreatePoli             `json:"poli" validate:"required"`
	COB                 string                         `json:"cob" validate:"required,number"`
	Catharact           string                         `json:"catharact" validate:"required,number"`
	VisitationPurpose   string                         `json:"visitation_purpose" validate:"required,number"`
	FlagProcedure       string                         `json:"flag_procedure" validate:"omitempty,number"`
	SupportCode         string                         `json:"support_code" validate:"omitempty,number"`
	ServiceAssessment   string                         `json:"service_assessment" validate:"omitempty,number"`
	AttendingDoctorCode string                         `json:"attending_doctor_code" validate:"omitempty,numeric"`
	Accounting          *BPJSAccounting                `json:"accounting" validate:"required"`
}

func (s *SelfSEPCreate) ToBPJS() *SEPCreate {
	return &SEPCreate{
		BPJSID:                s.CardNumber,
		ServiceDate:           s.ServiceDate,
		ServiceType:           s.ServiceType,
		MRNumber:              s.MRNumber,
		Note:                  s.Note,
		InitialDiagnosis:      s.InitialDiagnosis,
		VisitationPurpose:     s.VisitationPurpose,
		ProcedureFlag:         s.FlagProcedure,
		HealthCareSupportCode: s.SupportCode,
		ServiceAssessment:     s.ServiceAssessment,
		ServiceDPJP:           s.AttendingDoctorCode,
		UserID:                s.Accounting.OperatorID,
		User:                  s.Accounting.OperatorName,
		PhoneNum:              s.Accounting.OperatorPhone,
		COB: SEPCOB{
			COB: s.COB,
		},
		Cataracts: SEPCataract{
			Cataract: s.Catharact,
		},
		SKDP: SKDP{
			LetterNumber:           s.Control.ControllLetterNumber,
			AttendingPhysicianCode: s.Control.ReferredDoctorCode,
		},
		Polyclinics: SEPPolyclinics{
			PoliclinicCode: s.Poli.PolyclinicCode,
			Executive:      s.Poli.Executive,
		},
		Referral: SEPReferral{
			SourceReference:          s.Referral.ReferralSource,
			ReferenceDate:            s.Referral.ReferralDate,
			ReferenceNumber:          s.Referral.ReferralNumber,
			ReferencedHealthFacility: s.Referral.ReferrerHealthFacilityCode,
		},
		TreatmentClass: TreatmentClass{
			TreatmentClassRights:  s.TreatmentClass.TreatmentClass,
			TreatmentClassUpgrade: s.TreatmentClass.UpgradeTreatmentClass,
			Financing:             s.TreatmentClass.Financing,
			PIC:                   s.TreatmentClass.PIC,
		},
		Guarantee: Guarantee{
			Accident: s.Guarantee.AccidentType,
			LPNumber: s.Guarantee.RegisteredAccidentNumber,
			Guarantor: Guarantor{
				IncidentDate: s.Guarantee.Guarantor.AccidentDate,
				Note:         s.Guarantee.Guarantor.Note,
				Suppletion: Suppletion{
					Suppletion:       s.Guarantee.Guarantor.Suppletion.Suppletion,
					SuppletionNumber: s.Guarantee.Guarantor.Suppletion.SuppletionNumber,
					AccidentLocation: AccidentLocation{
						ProvinceCode: s.Guarantee.Guarantor.Suppletion.AccidentLocation.ProvinceCode,
						RegencyCode:  s.Guarantee.Guarantor.Suppletion.AccidentLocation.RegencyCode,
						DistrictCode: s.Guarantee.Guarantor.Suppletion.AccidentLocation.DistrictCode,
					},
				},
			},
		},
	}
}
