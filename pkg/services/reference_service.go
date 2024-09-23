package services

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type ReferenceService struct {
}

var _ interfaces.Reference = &ReferenceService{}

func (s *ReferenceService) DiagnoseReference(ctx context.Context) ([]*models.Reference, error) {
	return nil, nil
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
