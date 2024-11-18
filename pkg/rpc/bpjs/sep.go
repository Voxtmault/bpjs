package bpjs

import (
	"context"
	"encoding/json"

	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
	pb "github.com/voxtmault/bpjs-service-proto/go"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BPJSSEPRPCService struct {
	pb.UnimplementedSEPServiceServer
	Service interfaces.SEP
}

func InitSEPService() *BPJSSEPRPCService {
	// Init Services
	s := BPJSSEPRPCService{
		Service: services.NewSEPService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
	}
	return &s
}

func (s *BPJSSEPRPCService) CreateSEP(ctx context.Context, in *pb.SEPCreateRequest) (*pb.SEPCreateResponse, error) {
	// Parse the data
	var data models.SEPCreate
	if err := json.Unmarshal(in.GetSepJsonStr(), &data); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Validate the data
	if err := utils.GetValidator().StructCtx(ctx, data); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Call the service implementation
	result, err := s.Service.InsertSEP(ctx, &data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := pb.SEPCreateResponse{
		StatusCode: int32(code.Code_OK),
		Message:    "success creating SEP",
		Sep:        result.SEPNumber,
	}

	return &response, nil
}
