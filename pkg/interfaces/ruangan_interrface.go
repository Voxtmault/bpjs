package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

// Aplicares
type RuanganAplicares interface {
	GetReferensiJenisKamar(ctx context.Context) ([]*models.ReferenceJenisKamar, error)
	PostRuangan(ctx context.Context, obj *models.Ruangan, action string) (interface{}, error)
	GetKetersediaanKamar(ctx context.Context, start, limit string) ([]*models.RuanganResponse, error)
}
