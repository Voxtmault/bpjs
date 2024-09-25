package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type Reference interface {
	DiagnoseReference(ctx context.Context, diagnosisCode string) ([]*models.Reference, error)
	DoctorReference(ctx context.Context, jenisPelayanan, tglPelayanan, kodeSpesialis string) ([]*models.Reference, error)
	PoliclinicsReference(ctx context.Context, poliCode string) ([]*models.Reference, error)
	HealthFacilityReference(ctx context.Context, namaFaskes, jenisFaskes string) ([]*models.Reference, error)
	ProcedureReference(ctx context.Context, procedure string) ([]*models.Reference, error)
	NursingClassReference(ctx context.Context) ([]*models.Reference, error)
	SpecialistReference(ctx context.Context) ([]*models.Reference, error)
	DischargeMethodReference(ctx context.Context) ([]*models.Reference, error)
	PostDischargeReference(ctx context.Context) ([]*models.Reference, error)
	ProvinceReference(ctx context.Context) ([]*models.Reference, error)                              // Provinsi
	RegencyReference(ctx context.Context, kodeProvinsi string) ([]*models.Reference, error)          // Kabupaten
	DistrictReference(ctx context.Context, kodeKota string) ([]*models.Reference, error)             // Kecamatan
	AttendingPhysicianReference(ctx context.Context, kodeDokter string) ([]*models.Reference, error) // DPJP
}
