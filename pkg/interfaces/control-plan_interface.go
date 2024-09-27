package interfaces

import "context"

type ControlPlan interface {

	// Get are divided into 2 different functions since both of them returns different data models
	GetViaSEP(ctx context.Context, sepNumber string) (any, error)
	GetViaControlLetterNumber(ctx context.Context, controlLetterNumber string) (any, error)
}
