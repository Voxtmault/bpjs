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

func TestInsertSEP(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SEPService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	// obj := models.SEPCreate{
	// 	BPJSID:             "0002017051402",
	// 	ServiceDate:        time.Now().Format(time.DateOnly),
	// 	HealthFacilityCode: "0301R011",
	// 	ServiceType:        "2",
	// 	TreatmentClass: &models.TreatmentClass{
	// 		TreatmentClassRights: "3",
	// 	},
	// 	MRNumber:         "000001",
	// 	Reference:        &models.SEPReference{},
	// 	InitialDiagnosis: "E10",
	// 	Policlinics: &models.SEPPoliclinics{
	// 		PoliclinicCode: "",
	// 		Executive:      "0",
	// 	},
	// 	COB: &models.SEPCOB{
	// 		COB: "0",
	// 	},
	// 	Catharacs: &models.SEPCatharacs{
	// 		Catharacs: "0",
	// 	},
	// 	Guarantee: &models.Guarantee{
	// 		Accident: "0",
	// 		LPNumber: "12345",
	// 		Guarantor: &models.Guarantor{
	// 			Suppletion: &models.Suppletion{
	// 				AccidentLocation: &models.AccidentLocation{},
	// 			},
	// 		},
	// 	},
	// 	VisitationPurpose:     "0",
	// 	ServiceAssessment:     "",
	// 	SKDP:                  &models.SKDP{},
	// 	ServiceDPJP:           "000002",
	// 	PhoneNum:              "123456789098767",
	// 	User:                  "Testing",
	// 	Note:                  "",
	// 	ProcedureFlag:         "",
	// 	HealthCareSupportCode: "",
	// }

	sample := `{
                 "noKartu":"0002017051402",
                 "tglSep":"2024-09-26",
                 "ppkPelayanan":"0182R009",
                 "jnsPelayanan":"2",
                 "klsRawat":{
                    "klsRawatHak":"3",
                    "klsRawatNaik":"",
                    "pembiayaan":"",
                    "penanggungJawab":""
                 },
                 "noMR":"000001",
                 "rujukan":{
                    "asalRujukan":"",
                    "tglRujukan":"",
                    "noRujukan":"",
                    "ppkRujukan":""
                 },
                 "catatan":"testinsert RJ",
                 "diagAwal":"E10",
                 "poli":{
                    "tujuan":"INT",
                    "eksekutif":"0"
                 },
                 "cob":{
                    "cob":"0"
                 },
                 "katarak":{
                    "katarak":"0"
                 },
                 "jaminan":{
                    "lakaLantas":"0",
                    "noLP":"",
                    "penjamin":{
                       "tglKejadian":"",
                       "keterangan":"",
                       "suplesi":{
                          "suplesi":"0",
                          "noSepSuplesi":"",
                          "lokasiLaka":{
                             "kdPropinsi":"",
                             "kdKabupaten":"",
                             "kdKecamatan":""
                          }
                       }
                    }
                 },
                 "tujuanKunj":"0",
                 "flagProcedure":"",
                 "kdPenunjang":"",
                 "assesmentPel":"",
                 "skdp":{
                    "noSurat":"",
                    "kodeDPJP":""
                 },
                 "dpjpLayan":"31486",
                 "noTelp":"081111111101",
                 "user":"Coba Ws"
              }`

	var obj models.SEPCreate
	if err := json.Unmarshal([]byte(sample), &obj); err != nil {
		t.Errorf("Error unmarshalling the object: %v", err)
	}

	data, err := s.InsertSEP(context.Background(), &obj)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}

func TestUpdateSEP(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SEPService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	sample := `
	 {
      "noSep": "0301R0110521V000037",
      "klsRawat": {
        "klsRawatHak": "3",
        "klsRawatNaik": "",
        "pembiayaan": "",
        "penanggungJawab": ""
      },
      "noMR": "00469120",
      "catatan": "",
      "diagAwal": "E10",
      "poli": {
        "tujuan": "IGD",
        "eksekutif": "0"
      },
      "cob": {
        "cob": "0"
      },
      "katarak": {
        "katarak": "0"
      },
      "jaminan": {
        "lakaLantas": "0",
        "penjamin": {
          "tglKejadian": "",
          "keterangan": "",
          "suplesi": {
            "suplesi": "0",
            "noSepSuplesi": "",
            "lokasiLaka": {
              "kdPropinsi": "",
              "kdKabupaten": "",
              "kdKecamatan": ""
            }
          }
        }
      },
      "dpjpLayan": "46",
      "noTelp": "08522038363",
      "user": "Cobaws"
    }
	`

	var obj models.SEPUpdate
	if err := json.Unmarshal([]byte(sample), &obj); err != nil {
		t.Errorf("Error unmarshalling the object: %v", err)
	}

	data, err := s.UpdateSEP(context.Background(), &obj)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}

func TestDeleteSEP(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SEPService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	sample := `
	  {
             "noSep": "0301R0011017V000007",
             "user": "Coba Ws"
          }
	`

	var obj models.SEPDelete
	if err := json.Unmarshal([]byte(sample), &obj); err != nil {
		t.Errorf("Error unmarshalling the object: %v", err)
	}

	data, err := s.DeleteSEP(context.Background(), &obj)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}

func TestGetSEP(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SEPService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetSEP(context.Background(), "0001300759569")
	if err != nil {
		log.Println("Errors", err)
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}

func TestRequestSEP(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SEPService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	sample := `
	{
	"noKartu": "0001300759569",
	"tglSep": "2024-08-26",
	"jnsPelayanan": "1",
	"jnsPengajuan": "2",
	"keterangan": "Hari libur",
	"user": "Coba Ws"
	}
	`

	var obj models.SEPRequestCreate
	if err := json.Unmarshal([]byte(sample), &obj); err != nil {
		t.Errorf("Error unmarshalling the object: %v", err)
	}

	data, err := s.RequestSEP(context.Background(), &obj)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}

func TestApprovalSEPRequest(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SEPService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	sample := `
	{
	"noKartu": "0001300759569",
	"tglSep": "2024-08-26",
	"jnsPelayanan": "1",
	"jnsPengajuan": "2",
	"keterangan": "Hari libur",
	"user": "Coba Ws"
	}
	`

	var obj models.SEPRequestCreate
	if err := json.Unmarshal([]byte(sample), &obj); err != nil {
		t.Errorf("Error unmarshalling the object: %v", err)
	}

	data, err := s.ApprovalSEPRequest(context.Background(), &obj)
	if err != nil {
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}

func TestGetSEPRequests(t *testing.T) {
	// Load the config
	config.New("/home/andy/go-projects/rs/bpjs/.env")

	s := SEPService{
		HttpHandler: &RequestHandlerService{
			Security: &BPJSSecurityService{},
		},
	}

	data, err := s.GetSEPRequests(context.Background(), "08", "2024")
	if err != nil {
		log.Println("Errors", err)
		log.Println("Root Error", eris.Cause(err))
		t.Errorf("Error creating SEP: %v", err)
	} else {
		log.Println("Data: ", data)
	}
}
