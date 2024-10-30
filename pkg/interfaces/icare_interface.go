package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

// Aplicares
type Icare interface {
	FKRTLIcare(ctx context.Context, obj *models.FKRTLRequest) (interface{}, error)
}
