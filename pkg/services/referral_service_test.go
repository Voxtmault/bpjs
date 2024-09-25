package services

import (
	"context"
	"log"
	"testing"

	"github.com/voxtmault/bpjs-rs-module/config"
)

func TestGetParticipantReferralByReferralNumber(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetParticipantReferralByReferralNumber(context.Background(), "0301R0011117B001126", 1)
	if err != nil {
		t.Errorf("Error getting referral: %v", err)
	}

	for _, item := range data {
		log.Println("Referral: ", item)
	}
}
