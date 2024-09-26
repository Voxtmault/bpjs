package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	pb "github.com/voxtmault/bpjs-service-proto/go"
)

type Reference struct {
	Code string `json:"kode"`
	Name string `json:"nama"`
}

type DiagnosisReference struct {
	Diagnosis []*Reference `json:"diagnosa"`
}

type DoctorReference struct {
	Doctor []*Reference `json:"list"`
}

type ListReference struct {
	List []*Reference `json:"list"`
}

type PoliReference struct {
	Poli []*Reference `json:"poli"`
}

type FaskesReference struct {
	Faskes []*Reference `json:"faskes"`
}

type ProcedureReference struct {
	Procedure []*Reference `json:"procedure"`
}

// Types used for validation
type DiagnosisReferenceParams struct {
	Code string `validate:"required,min=3"`
}

type DoctorReferenceParams struct {
	ServiceType    string `validate:"required,number,min=1,max=2"`
	ServiceDate    string `validate:"required,datetime=2006-01-02"`
	SpecialistCode string `validate:"required,number,min=1"`
}

type PoliReferenceParams struct {
	PoliCode string `validate:"omitempty,uppercase"`
	PoliName string `validate:"omitempty"`
}

type HealthFacilityReferenceParams struct {
	HealthFacilityName string `validate:"required"`
	HealthFacilityType string `validate:"required,number,min=1,max=2"`
}

type ProcedureReferenceParams struct {
	ProcedureCode string `validate:"required"`
}

type RegencyReferenceParams struct {
	ProvinceCode string `validate:"required,number,min=2"`
}

type DistrictReferenceParams struct {
	RegencyCode string `validate:"required,number,min=4"`
}

type AttendingPhysicianReferenceParams struct {
	DoctorName string `validate:"required,min=3"`
}

// Custom validation function for ICD-10 codes, currently not implemented.
func ICD10Validation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[A-Z][0-9]{2}(\.[0-9]{1,2})?$`)
	return re.MatchString(fl.Field().String())
}

// Custom validation function for ICD-9 codes, currently not implemented.
func ICD9Validation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[0-9]{3}(\.[0-9]{1,2})?$`)
	return re.MatchString(fl.Field().String())
}

// Utils
func (s *DiagnosisReference) ToRPCArr() []*pb.Reference {
	arrObj := []*pb.Reference{}
	for _, item := range s.Diagnosis {
		arrObj = append(arrObj, &pb.Reference{
			Code: item.Code,
			Name: item.Name,
		})
	}

	return arrObj
}

func (s *Reference) ToRPC() *pb.Reference {
	return &pb.Reference{
		Code: s.Code,
		Name: s.Name,
	}
}
