package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type Participant interface {
	GetParticipant(ctx context.Context, params *models.ParticipantSearchParams) (*models.BPJSParticipant, error)
	GetParticipantV2(ctx context.Context, query *models.ParticipantSearchParams) (*models.SelfParticipant, error)
}
