package bpjs

import (
	"context"
	"fmt"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/voxtmault/bpjs-service-proto/go"
)

type BPJSParticipantRPCService struct {
	pb.UnimplementedParticipantServiceServer
	Service interfaces.Participant
}

func InitParticipantService() *BPJSParticipantRPCService {
	// Init Services
	s := BPJSParticipantRPCService{
		Service: &services.BPJSParticipantService{
			HttpHandler: &services.RequestHandlerService{
				Security: &services.BPJSSecurityService{},
			},
		},
	}
	return &s
}

func (s *BPJSParticipantRPCService) GetParticipant(ctx context.Context, in *pb.GetParticipantRequest) (*pb.GetParticipantResponse, error) {
	// Logic
	// 1. Convert the request into internal model
	// 2. Validate the model
	// 3. Call the get service

	params := &models.ParticipantSearchParams{
		NIK:         in.GetNIK(),
		BPJSNumber:  in.GetBpjsID(),
		ServiceDate: in.GetServiceDate(),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("invalid request: %w", err).Error())
	}

	data, err := s.Service.GetParticipant(ctx, params)
	if err != nil {
		// log.Println(eris.Cause(err))

		if data != nil {
			// Meaning that there is no internal error, but a business error
			return nil, status.Error(codes.InvalidArgument, eris.Cause(err).Error())
		} else {
			return nil, status.Errorf(codes.Internal, fmt.Errorf("failed to get participant: %w", err).Error())
		}
	}

	return &pb.GetParticipantResponse{
		StatusCode:  int32(codes.OK),
		Message:     "Success",
		Participant: data.ToProto(),
	}, nil
}
