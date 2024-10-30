package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type PRB interface {
	PRBInsert(ctx context.Context, obj *models.PRBInsertRequest) (*models.PRBResponse, error)
	PRBUpdate(ctx context.Context, obj *models.PRBUpdateRequest) (string, error)
	PRBDelete(ctx context.Context, obj *models.PRBDeleteRequest) (string, error)
	GetPRBbyNomorSRB(ctx context.Context, noPrb, noSep string) (string, error)
	GetPRBbyTanggal(ctx context.Context, tanggalAwal, tanggalAkhir string) (string, error)
}
