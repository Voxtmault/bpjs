package services

import (
	"context"
	"log"
	"testing"

	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

func TestSuplesi(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SuppletionService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.Suplesi(context.Background(), &models.SEPSuppletionParams{
		BPJSNumber: "0002017051402",
		// ServiceDate: time.Now().Format(time.DateOnly),
		ServiceDate: "2018-08-06",
	})
	if err != nil {
		t.Errorf("Error getting suppletion: %v", err)
	}

	for _, item := range data {
		log.Println("Suplesi: ", item)
	}
}

func TestAccidentMasterData(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SuppletionService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.AccidentMasterData(context.Background(), &models.SEPTrafficAccidentParams{
		BPJSNumber: "0001335263984",
	})
	if err != nil {
		t.Errorf("Error getting accident master data: %v", err)
	}

	for _, item := range data {
		log.Println("List: ", item)
	}
}
