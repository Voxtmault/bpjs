package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type Participant interface {
	GetParticipant(ctx context.Context, params *models.ParticipantSearchParams) (*models.BPJSParticipant, error)
}
