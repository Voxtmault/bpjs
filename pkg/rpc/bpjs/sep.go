package bpjs

import (
	"context"
	"encoding/json"

	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
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
	var data models.SelfSEPCreate
	if err := json.Unmarshal(in.GetSepJsonStr(), &data); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bpjsFormat := data.ToBPJS()
	bpjsFormat.HealthFacilityCode = config.GetConfig().BPJSConfig.PPKCode

	if bpjsFormat.Referral.ReferencedHealthFacility == "" {
		bpjsFormat.Referral.ReferencedHealthFacility = bpjsFormat.HealthFacilityCode
	}

	// Call the service implementation
	result, err := s.Service.InsertSEP(ctx, bpjsFormat)
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

func (s *BPJSSEPRPCService) GetSEP(ctx context.Context, in *pb.SEPGetRequest) (response *pb.SEPGetResponse, err error) {

	result, err := s.Service.GetSEP(ctx, in.GetSepNumber())
	if err != nil {
		data, marshallErr := json.Marshal(result)
		if marshallErr != nil {
			err = status.Error(codes.Internal, marshallErr.Error())
			return
		}
		response = &pb.SEPGetResponse{
			StatusCode: int32(code.Code_INTERNAL),
			SepJsonStr: data,
		}
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return
	}

	response = &pb.SEPGetResponse{
		StatusCode: int32(code.Code_OK),
		SepJsonStr: data,
		Message:    "success getting SEP",
	}

	return
}
