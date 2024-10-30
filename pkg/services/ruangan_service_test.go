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

func TestGetReferensiKamar(t *testing.T) {
	// Load the config
	config.New("../../.env")

	s := RuanganService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetReferensiJenisKamar(context.Background())
	if err != nil {
		log.Println("Errors", err)
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		result, _ := json.Marshal(data)
		log.Print("Data: ", string(result))
	}
}

func TestGetKeterdesiaanKamar(t *testing.T) {
	// Load the config
	config.New("../../.env")

	s := RuanganService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetKetersediaanKamar(context.Background(), "1", "10")
	if err != nil {
		log.Println("Errors", err)
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		result, _ := json.Marshal(data)
		log.Print("Data: ", string(result))
	}
}

func TestPostRuangan(t *testing.T) {
	// Load the config
	config.New("../../.env")

	s := RuanganService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	sample := `                                
  { 
    "kodekelas":"VIP", 
    "koderuang":"RG01", 
    "namaruang":"Ruang Anggrek VIP", 
    "kapasitas":"20", 
    "tersedia":"10",
    "tersediapria":"0", 
    "tersediawanita":"0", 
    "tersediapriawanita":"0"
   }`

	var obj models.Ruangan
	if err := json.Unmarshal([]byte(sample), &obj); err != nil {
		t.Errorf("Error unmarshalling the object: %v", err)
	}

	data, err := s.PostRuangan(context.Background(), &obj, "create")
	if err != nil {
		log.Println("Errors", err)
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		result, _ := json.Marshal(data)
		log.Print("Data: ", string(result))
	}
}
