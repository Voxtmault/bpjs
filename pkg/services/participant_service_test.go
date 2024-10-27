package services

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

func TestGetParticipant(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := BPJSParticipantService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetParticipant(context.Background(), &models.ParticipantSearchParams{
		BPJSNumber: "0002088008515",
		// NIK: "1234567890987651",
	})
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting participant: %v", err)
	}

	result, _ := json.Marshal(data)
	log.Print("Data: ", string(result))
}

func TestGetParticipantByReferralNumber(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetParticipantReferralByReferralNumber(context.Background(), "0011336526592", models.PCareSource)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting participant referral by referral number: %v", err)
	}

	log.Println("Data: ", data)
}
