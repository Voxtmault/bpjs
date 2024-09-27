package services

import (
	"context"

	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
)

type ControlPlanService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.ControlPlan = &ControlPlanService{}

func (s *ControlPlanService) GetViaSEP(ctx context.Context, sepNumber string) (any, error) {
	return nil, nil
}

func (s *ControlPlanService) GetViaControlLetterNumber(ctx context.Context, letterNumber string) (any, error) {
	return nil, nil
}
