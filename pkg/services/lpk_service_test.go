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

func TestLPKInsert(t *testing.T) {
	config.New("../../.env")

	s := LPKService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}
	// param nomor kartu
	bodyReq := `                                                   
 {
          
                 "noSep": "0182R0091024V000002",
                 "tglMasuk": "2024-10-30",
                 "tglKeluar": "2024-10-30",
                 "jaminan": "1",
                 "poli": {
                    "poli": "INT"
                 },
                 "perawatan": {
                    "ruangRawat": "1",
                    "kelasRawat": "1",
                    "spesialistik": "1",
                    "caraKeluar": "1",
                    "kondisiPulang": "1"
                 },
                 "diagnosa": [
                    {
                       "kode": "N88.0",
                       "level": "1"
                    },
                    {
                       "kode": "A00.1",
                       "level": "2"
                    }
                 ],
                 "procedure": [
                    {
                       "kode": "00.82"
                    },
                    {
                       "kode": "00.83"
                    }
                 ],
                 "rencanaTL": {
                    "tindakLanjut": "1",
                    "dirujukKe": {
                       "kodePPK": ""
                    },
                    "kontrolKembali": {
                       "tglKontrol": "2017-11-10",
                       "poli": ""
                    }
                 },
                 "DPJP": "3",
                 "user": "Coba Ws"
              
        }                         
	`
	var obj models.InsertLPKRequest
	json.Unmarshal([]byte(bodyReq), &obj)
	data, err := s.LPKInsert(context.Background(), &obj)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error getting participant: %v", err)
	}

	result, _ := json.Marshal(data)
	log.Print("Data: ", string(result))
}
