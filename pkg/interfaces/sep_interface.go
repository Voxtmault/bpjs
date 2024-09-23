package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type SEP interface {
	InsertSEP(ctx context.Context, obj *models.SEPCreate) (*models.SEPCreateResponse, error)
	UpdateSEP(ctx context.Context) (any, error)
	DeleteSEP(ctx context.Context) (any, error)
}
