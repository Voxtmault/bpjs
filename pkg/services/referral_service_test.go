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

	data, err := s.CreateReferral(context.Background(), &models.ReferralAction{
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
		t.Errorf("Error creting referral: %v", err)
	}

	log.Println("Data: ", data)
}

func TestUpdateReferral(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.UpdateReferral(context.Background(), &models.ReferralAction{
		ReferralNumber:             "0301R0011117B001126",
		ReferralDate:               time.Now().Format(time.DateOnly),
		PlannedVisitDate:           time.Now().AddDate(0, 0, 2).Format(time.DateOnly),
		ReferredHealthFacilityCode: "0182R009",
		ServiceType:                "1",
		Note:                       "Test Update Referral",
		ReferralDiagnosis:          "A15",
		ReferralType:               "1",
		ReferredPoliclinicCode:     "INT",
		User:                       "Test Update Referral",
	})
	if err != nil {
		t.Errorf("Error updating referral: %v", err)
	}

	log.Println("Data: ", data)
}

func TestDeleteReferral(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	err := s.DeleteReferral(context.Background(), "0301R0011117B001126", "Test Delete Referral")
	if err != nil {
		t.Errorf("Error deleting referral: %v", err)
	}
}

func TestGetOutgoingReferral(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetOutgoingReferral(context.Background(), time.Now().AddDate(0, 0, -1).Format(time.DateOnly), time.Now().Format(time.DateOnly))
	if err != nil {
		t.Errorf("Error getting outgoing referral: %v", err)
	}

	for _, item := range data {
		log.Println("Outgoing Referral: ", item)
	}
}

func TestGetOutgoingReferralDetail(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetOutgoingReferralDetail(context.Background(), "0301R0011117B001126")
	if err != nil {
		t.Errorf("Error getting outgoing referral detail: %v", err)
	}

	log.Println("Outgoing Referral Detail: ", data)
}

func TestCreateSpecialReferral(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.CreateSpecialReferral(context.Background(), &models.SpecialReferralCreate{
		ReferralNumber: "0301R0011117B001126",
		Diagnosises: []*models.SpecialReferralDiagnosis{
			{
				Code: "P;N18",
			},
			{
				Code: "S;N18.1",
			},
		},
		Procedures: []*models.Reference{
			{
				Code: "39.95",
			},
		},
		User: "Test Create Special Referral",
	})
	if err != nil {
		t.Errorf("Error creting special referral: %v", err)
	}

	log.Println("Data: ", data)
}

func TestGetSpecialReferrals(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetSpecialReferrals(context.Background(), "1", "2024")
	if err != nil {
		t.Errorf("Error getting special referrals: %v", err)
	}

	for _, item := range data {
		log.Println("Special Referral: ", item)
	}
}

func TestDeleteSpecialReferral(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferralService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.DeleteSpecialReferral(context.Background(), &models.SpecialReferralDelete{
		ReferralID:     "98865",
		ReferralNumber: "0301U0331019P003283",
		User:           "Test Delete Special Referral",
	})
	if err != nil {
		t.Errorf("Error deleting special referral: %v", err)
	}

	log.Println("Data: ", data)
}
