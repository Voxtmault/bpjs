package services

import (
	"context"
	"log"
	"testing"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
)

func TestGetReferensiKamar(t *testing.T) {
	// Load the config
	config.New("../../.env")

	s := RuanganService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetReferensiJenisKamar(context.Background())
	if err != nil {
		log.Println("Errors", err)
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}
