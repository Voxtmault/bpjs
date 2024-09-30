package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

// Rujukan
type Referral interface {
	GetParticipantReferralByReferralNumber(ctx context.Context, referalNumber string, source uint) ([]*models.Referral, error)
	GetParticipantReferralByBPJSNumber(ctx context.Context, bpjsNumber string, source uint, multi bool) ([]*models.Referral, error)

	// GetReferedSpecialist is used to get lists of specialists available from the refered health facility
	GetReferedSpecialist(ctx context.Context, referedHealthFacilityCode, referalDate string) ([]*models.ReferredSpecialist, error)

	// GetReferedFacilites is used to get lists of health facilities available from the refered health facility, eg: Labs, Radiology, MRI, Hemodialisis, etc...
	GetReferedFacilities(ctx context.Context, referedHealthFacilityCode string) ([]*models.ReferredFacility, error)

	// CUD Operations
	CreateReferral(ctx context.Context, obj *models.ReferralCreate) (*models.ReferralCreateResponse, error)
}
