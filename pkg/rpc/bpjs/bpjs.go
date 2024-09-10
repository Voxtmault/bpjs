package bpjs

import (
	"context"
	"time"

	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
	"github.com/voxtmault/bpjs-rs-module/pkg/storage"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/voxtmault/bpjs-service-proto/go"
)

type BPJSSubService struct {
	pb.UnimplementedBPJSServiceServer
	controller BPJSController
}

type BPJSController struct {
	ExampleService interfaces.Example
}

func InitBPJSSubService() *BPJSSubService {
	// Init Services
	exampleImplementation := &services.ExampleService{
		Con:       storage.GetDBConnection(),
		Injection: &services.DependencyInjectionService{},
	}
	authController := BPJSController{ExampleService: exampleImplementation}

	s := BPJSSubService{
		controller: authController,
	}
	return &s
}

func (s *BPJSSubService) SampleService(ctx context.Context, in *pb.SampleServiceRequest) (*pb.SampleServiceResponse, error) {

	// Example of using context with timeout
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// ALWAYS validate inputs from the client
	validator := utils.GetValidator()

	// Example of using validator
	err := validator.Struct(in.GetSample())
	if err != nil {
		// Please use the appropriate error code based on the error
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	// If for some reason this function is running longer than 5 seconds, it will be cancelled automatically
	s.controller.ExampleService.SampleCreate(ctx, []string{})

	return &pb.SampleServiceResponse{
		ResponseCode: int32(codes.OK),
		Message:      "Success",
		Data: &pb.SampleStruct{
			Name: "John Doe",
			// etc
		},
	}, nil
}
