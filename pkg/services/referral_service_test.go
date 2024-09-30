package services

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
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

func TestGetReferedSpecialist(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetReferedSpecialist(context.Background(), "0182R009", time.Now().Format(time.DateOnly))
	if err != nil {
		t.Errorf("Error getting refered specialist: %v", err)
	}

	for _, item := range data {
		log.Println("Specialist: ", item)
	}
}

func TestGetReferedFacilities(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetReferedFacilities(context.Background(), "0182R009")
	if err != nil {
		t.Errorf("Error getting refered facilities: %v", err)
	}

	for _, item := range data {
		log.Println("Facility: ", item)
	}
}

func TestCreateReferral(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.CreateReferral(context.Background(), &models.ReferralCreate{
		SEPNumber:                  "0301R0010321V000003",
		ReferralDate:               time.Now().Format(time.DateOnly),
		PlannedVisitDate:           time.Now().AddDate(0, 0, 2).Format(time.DateOnly),
		ReferredHealthFacilityCode: "0182R009",
		ServiceType:                "1",
		Note:                       "Test",
		ReferralDiagnosis:          "A15",
		ReferralType:               "1",
		ReferredPoliclinicCode:     "INT",
		User:                       "Test Create Referral",
	})
	if err != nil {
		t.Errorf("Error getting refered facilities: %v", err)
	}

	log.Println("Data: ", data)
}
