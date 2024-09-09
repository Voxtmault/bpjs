package bpjs

import (
	"context"
	"time"

	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
	"github.com/voxtmault/bpjs-rs-module/pkg/storage"

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

	// If for some reason this function is running longer than 5 seconds, it will be cancelled automatically
	s.controller.ExampleService.SampleCreate(ctx, []string{})

	return nil, nil
}
