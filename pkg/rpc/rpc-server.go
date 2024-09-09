package rpc

import (
	subservice "github.com/voxtmault/bpjs-rs-module/pkg/rpc/bpjs"
	_ "google.golang.org/grpc/encoding/gzip"
)

type BPJSService struct {
	BPJSSubService *subservice.BPJSSubService
}

func InitRPCService() *BPJSService {
	s := BPJSService{
		BPJSSubService: subservice.InitBPJSSubService(),
	}
	return &s
}
