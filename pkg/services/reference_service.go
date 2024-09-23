package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type ReferenceService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.Reference = &ReferenceService{}

func (s *ReferenceService) DiagnoseReference(ctx context.Context, diagnosisCode string) ([]*models.Reference, error) {
	arrObj := models.DiagnosisReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/diagnosa"

	if diagnosisCode != "" {
		baseUrl += "/" + diagnosisCode
	}

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Diagnosis, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.Diagnosis, nil
}

func (s *ReferenceService) DoctorReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

func (s *ReferenceService) PoliclinicsReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

func (s *ReferenceService) HealthFacilityReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

func (s *ReferenceService) ProcedureReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

func (s *ReferenceService) NursingClassReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

func (s *ReferenceService) SpecialistReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

func (s *ReferenceService) DischargeMethodReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

func (s *ReferenceService) PostDischargeReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

// Provinsi
func (s *ReferenceService) ProvinceReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

// Kabupaten
func (s *ReferenceService) RegencyReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

// Kecamatan
func (s *ReferenceService) DistrictReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}

// DPJP
func (s *ReferenceService) AttendingPhysicianReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
}
