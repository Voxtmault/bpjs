package services

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
)

func TestGetMonitoringData(t *testing.T) {
	config.New("../../.env")

	service := Monitoring{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := service.GetMonitoringDataKlaim(context.Background(), "2024-10-22", "2", "1")
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting participant: %v", err)
	}

	result, _ := json.Marshal(data)
	log.Print("Data: ", string(result))
}
