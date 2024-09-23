package services

import (
	"context"
	"log"
	"testing"

	"github.com/voxtmault/bpjs-rs-module/config"
)

func TestGetDiagnosis(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferenceService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.DiagnoseReference(context.Background(), "")
	if err != nil {
		t.Errorf("Error getting diagnosis: %v", err)
	}

	for _, item := range data {
		log.Println("Diagnosis: ", item)
	}

}
