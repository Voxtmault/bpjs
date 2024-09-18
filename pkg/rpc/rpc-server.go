package rpc

import (
	subservice "github.com/voxtmault/bpjs-rs-module/pkg/rpc/bpjs"
	_ "google.golang.org/grpc/encoding/gzip"
)

type BPJSService struct {
	ParticipantService *subservice.BPJSParticipantRPCService
}

func InitRPCService() *BPJSService {
	s := BPJSService{
		ParticipantService: subservice.InitParticipantService(),
	}
	return &s
}
