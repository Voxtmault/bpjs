package services

import (
	"context"
	"log"
	"strconv"
	"testing"
	"time"

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

	data, err := s.DiagnoseReference(context.Background(), "B201")
	if err != nil {
		t.Errorf("Error getting diagnosis: %v", err)
	}

	for _, item := range data {
		log.Println("Diagnosis: ", item)
	}
}

func TestGetPoliclinics(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferenceService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.PoliclinicsReference(context.Background(), "Penyakit")
	if err != nil {
		t.Errorf("Error getting policlinics reference: %v", err)
	}

	for _, item := range data {
		log.Println("Policlinics: ", item)
	}
}

func TestGetHealthFacilityReference(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferenceService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.HealthFacilityReference(context.Background(), "0182R009", "2")
	if err != nil {
		t.Errorf("Error getting health facility reference: %v", err)
	}

	for _, item := range data {
		log.Println("Health Facilities: ", item)
	}
}

func TestGetSpecialistReference(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferenceService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.SpecialistReference(context.Background())
	if err != nil {
		t.Errorf("Error getting specialist reference: %v", err)
	}

	for _, item := range data {
		log.Println("Specialists: ", item)
	}
}

func TestGetDoctorReference(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := ReferenceService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	for i := 1; i < 30; i++ {
		data, err := s.DoctorReference(context.Background(), "1", time.Now().Format(time.DateOnly), strconv.Itoa(i))
		if err != nil {
			t.Errorf("Error getting doctor reference: %v", err)
		}

		for _, item := range data {
			log.Println("Doctors: ", item)
		}
	}
}
