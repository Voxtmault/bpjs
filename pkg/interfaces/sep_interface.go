package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type SEP interface {
	InsertSEP(ctx context.Context, obj *models.SEPCreate) (*models.SEPCreateResponse, error)
	UpdateSEP(ctx context.Context, obj *models.SEPUpdate) (string, error)
	DeleteSEP(ctx context.Context, obj *models.SEPDelete) (string, error)
	GetSEP(ctx context.Context, sepNumber string) (*models.SEPGet, error)

	// For SEP Backdate and Fingerprint Request
	RequestSEP(ctx context.Context, obj *models.SEPRequestCreate) (string, error)

	// For SEP Approval, Only Fingerprint request may be accepted by the HealthCare (Rumah Sakit).
	// For other types of request, it will be done by BPJS Personnel
	ApprovalSEPRequest(ctx context.Context, obj *models.SEPRequestCreate) (string, error)
	GetSEPRequests(ctx context.Context, month, year string) ([]*models.SEPRequest, error)

	GetInternalSEP(ctx context.Context, sepNumber string) ([]*models.InternalSEP, error)

	UpdateTanggalPulang(ctx context.Context, obj *models.SEPUpdateTanggalPulangRequest) error
	GetDischargedSEP(ctx context.Context, month, year, filter string) ([]*models.DischargedSEP, error)
	GetFingerPrintSEP(ctx context.Context, noKartu, tanggalPelayanan string) (*models.SEPGetFingerPrint, error)
	GetListFingerPrintSEP(ctx context.Context, tanggalPelayanan string) ([]*models.SEPGetListFingerPrint, error)
	GetListRandomQuestion(ctx context.Context, noKartu, tanggalPelayanan string) ([]*models.SEPGetRandomQuestion, error)
}

type SuplesiJasaRaharja interface {
	Suplesi(ctx context.Context, params *models.SEPSuppletionParams) ([]*models.SEPSuppletion, error)                      // Suplesi Jasa Raharja
	AccidentMasterData(ctx context.Context, params *models.SEPTrafficAccidentParams) ([]*models.SEPTrafficAccident, error) // Data Induk Kecelakaan
}
