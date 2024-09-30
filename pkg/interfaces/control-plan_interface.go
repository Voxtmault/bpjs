package interfaces

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type ControlPlan interface {

	// Get are divided into 2 different functions since both of them returns different data models
	GetViaSEP(ctx context.Context, sepNumber string) ([]*models.ControlPlanGetViaSEP, error)
	GetViaControlLetterNumber(ctx context.Context, controlLetterNumber string) ([]*models.ControlPlanGetViaControllLetterNumber, error)

	// Used to get lists of control plans registered to a card number
	GetControlPlanFromCardNumber(ctx context.Context, params *models.ControlPlansFromCardNumberParams) ([]*models.ControlPlans, error)

	// Used to get registered control plans, the BPJS Documentation didn't specify from where
	// but it's safe to assume it's from the health care / hospital
	GetControlPlans(ctx context.Context, params *models.ControlPlanParams) ([]*models.ControlPlans, error)

	// Used to get how many controlls are assigned to the said clinic
	GetClinicControlPlans(ctx context.Context, params *models.ClinicControlParams) ([]*models.ClinicControlPlans, error)

	// Used to get the schedule of a doctor in a clinic
	GetDoctorPracticeSchedule(ctx context.Context, params *models.DoctorScheduleParams) ([]*models.DoctorPracticeSchedule, error)

	// CUD Operations
	CreateControlPlan(ctx context.Context, obj *models.ControlPlanCreate) (*models.ControlPlanCreateResponse, error)
	UpdateControlPlan(ctx context.Context, obj *models.UpdateControlPlans) (*models.ControlPlanCreateResponse, error)
	DeleteControlPlan(ctx context.Context, controlNumber, user string) error

	// Used to insert / create an inpatient care order (SPRI / Surat Perintah Rawat Inap)
	CreateInpatientCareOrder(ctx context.Context, obj *models.ControlPlanCreate) (*models.ControlPlanCreateResponse, error)
	UpdateInpatientCareOrder(ctx context.Context, obj *models.UpdateControlPlans) (*models.ControlPlanCreateResponse, error)
}
