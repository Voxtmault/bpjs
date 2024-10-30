package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type LPK interface {
	LPKInsert(ctx context.Context, obj *models.InsertLPKRequest) (string, error)
	LPKUpdate(ctx context.Context, obj *models.InsertLPKRequest) (string, error)
	LPKDelete(ctx context.Context, noSep string) (string, error)
	LPKGet(ctx context.Context, tanggalMasuk, jenisLayanan string) ([]*models.ListLPKResponse, error)
}
