package services

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/voxtmault/bpjs-rs-module/config"
)

func TestGetDiagnosis(t *testing.T) {
	// Load the config
	config.New(`D:\Goland\bpjs\.env`)

	s := ReferenceService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.PoliclinicsReference(context.Background(), "")
	if err != nil {
		t.Errorf("Error getting diagnosis: %v", err)
	}
	// time.Sleep(1 * time.Second)
	fmt.Println(data)
	for _, item := range data {
		log.Println("Diagnosis: ", item)
	}

}
