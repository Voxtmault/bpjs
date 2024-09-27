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
}

type SuplesiJasaRaharja interface {
	Suplesi(ctx context.Context, params *models.SEPSuppletionParams) ([]*models.SEPSuppletion, error)                      // Suplesi Jasa Raharja
	AccidentMasterData(ctx context.Context, params *models.SEPTrafficAccidentParams) ([]*models.SEPTrafficAccident, error) // Data Induk Kecelakaan
}
