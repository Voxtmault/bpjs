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

func TestFkrtlIcares(t *testing.T) {
	config.New("../../.env")

	s := IcareService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}
	// param nomor kartu
	bodyReq := `                                                   
{
"param": "0002088008515",
"kodedokter": 460285
}           
	`
	var obj models.FKRTLRequest
	json.Unmarshal([]byte(bodyReq), &obj)
	data, err := s.FKRTLIcare(context.Background(), &obj)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting participant: %v", err)
	}

	result, _ := json.Marshal(data)
	log.Print("Data: ", string(result))
}
