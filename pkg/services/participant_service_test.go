package services

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/logger"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/storage"
)

var envPath = "../../.env"

func TestGetParticipant(t *testing.T) {
	cfg := config.New(envPath)
	if err := storage.InitMariaDB(&cfg.DBConfig); err != nil {
		t.Errorf("Error initializing MariaDB: %v", err)
		return
	}
	if err := storage.InitRedis(&cfg.RedisConfig); err != nil {
		t.Errorf("Error initializing Redis: %v", err)
		return
	}
	logger.InitRequestLogger()

	service := NewParticipantService(
		NewBPJSRequestHandlerService(
			NewBPJSSecurityService(),
		),
	)

	data, err := service.GetParticipant(context.Background(), &models.ParticipantSearchParams{
		BPJSNumber: "0002088008515a",
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
