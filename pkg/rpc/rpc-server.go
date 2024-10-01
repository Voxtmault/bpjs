package rpc

import (
	subservice "github.com/voxtmault/bpjs-rs-module/pkg/rpc/bpjs"
	_ "google.golang.org/grpc/encoding/gzip"
)

type BPJSService struct {
	ParticipantService *subservice.BPJSParticipantRPCService
	ReferenceService   *subservice.BPJSReferenceRPCService
}

func InitRPCService() *BPJSService {

	// Init Services
	s := BPJSService{
		ParticipantService: subservice.InitParticipantService(),
		ReferenceService:   subservice.InitReferenceService(),
	}

	return &s
}
