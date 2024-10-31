package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

// Rujukan
type Referral interface {
	GetParticipantReferralByReferralNumber(ctx context.Context, referralNumber string, source uint) (*models.Referral, error)
	GetParticipantReferralByReferralNumberV2(ctx context.Context, referralNumber string, source uint) (*models.SelfReferralGet, error)

	GetParticipantReferralByBPJSNumber(ctx context.Context, bpjsNumber string, source uint, multi bool) ([]*models.Referral, error)

	// GetReferredSpecialist is used to get lists of specialists available from the referred health facility
	GetReferredSpecialist(ctx context.Context, referredHealthFacilityCode, referralDate string) ([]*models.ReferredSpecialist, error)

	// GetReferredFacilities is used to get lists of health facilities available from the referred health facility, eg: Labs, Radiology, MRI, Hemodialysis, etc...
	GetReferredFacilities(ctx context.Context, referredHealthFacilityCode string) ([]*models.ReferredFacility, error)

	// Get any referral that's referred to other health facility
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

	GetReferralSEPCount(ctx context.Context, referralType, referralNumber string) (string, error)
}
