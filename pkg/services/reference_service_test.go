package services

import (
	"context"
	"log"
	"testing"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

func TestGetParticipantByReferralNumber(t *testing.T) {
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	service := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetParticipantReferralByReferralNumber(context.Background(), "030107010217Y001465", models.PCareSource)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting participant referral by referral number: %v", err)
	}

	log.Println("Data: ", data)

}
