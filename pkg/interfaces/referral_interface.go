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

	// Get any referral that's refered to other health facility
	GetOutgoingReferral(ctx context.Context, startDate, endDate string) ([]*models.OutgoingReferral, error)

	GetOutgoingReferralDetail(ctx context.Context, referralNumber string) (*models.ReferralDetail, error)

	// CUD Operations
	CreateReferral(ctx context.Context, obj *models.ReferralAction) (*models.ReferralCreateResponse, error)
	UpdateReferral(ctx context.Context, obj *models.ReferralAction) (string, error)
	DeleteReferral(ctx context.Context, referralNumber, user string) error

	// Special Referral Operations
	CreateSpecialReferral(ctx context.Context, obj *models.SpecialReferralCreate) (*models.SpecialReferralCreateResponse, error)
	GetSpecialReferrals(ctx context.Context, month, year string) ([]*models.SpecialReferrals, error)
	DeleteSpecialReferral(ctx context.Context, obj *models.SpecialReferralDelete) (string, error)
}
