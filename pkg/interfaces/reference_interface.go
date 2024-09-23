package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type Reference interface {
	DiagnoseReference(ctx context.Context) ([]*models.Reference, error)
	DoctorReference(ctx context.Context) ([]*models.Reference, error)
	PoliclinicsReference(ctx context.Context) ([]*models.Reference, error)
	HealthFacilityReference(ctx context.Context) ([]*models.Reference, error)
	ProcedureReference(ctx context.Context) ([]*models.Reference, error)
	NursingClassReference(ctx context.Context) ([]*models.Reference, error)
	SpecialistReference(ctx context.Context) ([]*models.Reference, error)
	DischargeMethodReference(ctx context.Context) ([]*models.Reference, error)
	PostDischargeReference(ctx context.Context) ([]*models.Reference, error)
	ProvinceReference(ctx context.Context) ([]*models.Reference, error)           // Provinsi
	RegencyReference(ctx context.Context) ([]*models.Reference, error)            // Kabupaten
	DistrictReference(ctx context.Context) ([]*models.Reference, error)           // Kecamatan
	AttendingPhysicianReference(ctx context.Context) ([]*models.Reference, error) // DPJP
}
