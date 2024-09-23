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
                 "noKartu":"0001105689835",
                 "tglSep":"2021-07-30",
                 "ppkPelayanan":"0301R011",
                 "jnsPelayanan":"1",
                 "klsRawat":{
                    "klsRawatHak":"2",
                    "klsRawatNaik":"1",
                    "pembiayaan":"1",
                    "penanggungJawab":"Pribadi"
                 },
                 "noMR":"MR9835",
                 "rujukan":{
                    "asalRujukan":"2",
                    "tglRujukan":"2021-07-23",
                    "noRujukan":"RJKMR9835001",
                    "ppkRujukan":"0301R011"
                 },
                 "catatan":"testinsert RI",
                 "diagAwal":"E10",
                 "poli":{
                    "tujuan":"",
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
                    "noLP":"12345",
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
                    "noSurat":"0301R0110721K000021",
                    "kodeDPJP":"31574"
                 },
                 "dpjpLayan":"",
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
	}

	log.Println("Data: ", data)
}
