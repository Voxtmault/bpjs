package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type Referral interface {
	GetParticipantReferralByReferralNumber(ctx context.Context, referalNumber string, source uint) ([]*models.Referral, error)
	GetParticipantReferralByBPJSNumber(ctx context.Context, bpjsNumber string, source uint, multi bool) ([]*models.Referral, error)
}
