package models

type SEPSuppletion struct {
	RegistrationNumber    string `json:"noRegister"`
	SEPNumber             string `json:"noSep"`
	InitialSEPNumber      string `json:"noSepAwal"`
	GuaranteeLetterNumber string `json:"noSuratJaminan"`
	AccidentDate          string `json:"tglKejadian"`
	SEPDate               string `json:"tglSep"`
}

type SEPSuppletionResponse struct {
	Guarantee []*SEPSuppletion `json:"jaminan"`
}

type SEPSuppletionParams struct {
	BPJSNumber  string `validate:"required,min=13"`
	ServiceDate string `validate:"required,datetime=2006-01-02"`
}

type SEPTrafficAccident struct {
	SEPNumber           string `json:"noSep"`
	AccidentDate        string `json:"tglKejadian"`
	HealthFacilitySEP   string `json:"ppkPelSEP"`
	ProvinceCode        string `json:"kdProp"`
	RegencyCode         string `json:"kdKab"`
	DistrictCode        string `json:"kdKec"`
	AccidentNote        string `json:"ketKejadian"`
	SuppletionSEPNumber string `json:"noSEPSuplesi"` // Might contain more than 1 SEP Number, combined into 1 string with the delimiter of comma followed by single space
}

type SEPTrafficAccidentResponse struct {
	List []*SEPTrafficAccident `json:"list"`
}

type SEPTrafficAccidentParams struct {
	BPJSNumber string `validate:"required,min=13"`
}
