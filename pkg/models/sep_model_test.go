package models

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
)

func TestInsertSEPModel(t *testing.T) {

	validate := validator.New()

	testCases := []struct {
		name      string
		jsonInput string
		expectErr bool
	}{
		{
			name: "Valid SEPCreate",
			jsonInput: `{
				"noKartu": "0002088008515",
				"tglSep": "2024-10-22",
				"ppkPelayanan": "0182R009",
				"jnsPelayanan": "2",
				"klsRawat": {
					"klsRawatHak": "3",
					"klsRawatNaik": "",
					"pembiayaan": "",
					"penanggungJawab": ""
				},
				"noMR": "000001",
				"rujukan": {
					"asalRujukan": "2",
					"tglRujukan": "2024-10-22",
					"noRujukan": "",
					"ppkRujukan": "0182R009"
				},
				"catatan": "testinsert RJ",
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
					"noLP": "",
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
				"tujuanKunj": "0",
				"flagProcedure": "",
				"kdPenunjang": "",
				"assesmentPel": "",
				"skdp": {
					"noSurat": "",
					"kodeDPJP": ""
				},
				"dpjpLayan": "460285",
				"noTelp": "081111111101",
				"user": "Coba Ws"
			}`,
			expectErr: false,
		},
		{
			name: "Invalid SEPCreate - Missing Required Field",
			jsonInput: `{
				"tglSep": "2024-10-22",
				"ppkPelayanan": "0182R009",
				"jnsPelayanan": "2",
				"klsRawat": {
					"klsRawatHak": "3",
					"klsRawatNaik": "",
					"pembiayaan": "",
					"penanggungJawab": ""
				},
				"noMR": "000001",
				"rujukan": {
					"asalRujukan": "2",
					"tglRujukan": "2024-10-22",
					"noRujukan": "",
					"ppkRujukan": "0182R009"
				},
				"catatan": "testinsert RJ",
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
					"noLP": "",
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
				"tujuanKunj": "0",
				"flagProcedure": "",
				"kdPenunjang": "",
				"assesmentPel": "",
				"skdp": {
					"noSurat": "",
					"kodeDPJP": ""
				},
				"dpjpLayan": "460285",
				"noTelp": "081111111101",
				"user": "Coba Ws"
			}`,
			expectErr: true,
		},
		// Add more test cases as needed
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var obj SEPCreate
			if err := json.Unmarshal([]byte(tc.jsonInput), &obj); err != nil {
				t.Fatalf("Error unmarshalling JSON: %v", err)
			}

			err := validate.Struct(obj)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectErr && err != nil {
				if errMap := utils.MangleValidateResult(err); err != nil {
					log.Println("Error Map: ", errMap)
				}
				t.Errorf("Did not expect error but got: %v", err)
			}
		})
	}
}

func TestTreatmentClassValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Valid input with TreatmentClassUpgrade",
			input:   `{"klsRawatHak": "1", "klsRawatNaik": "2", "pembiayaan": "1", "penanggungJawab": "1"}`,
			wantErr: false,
		},
		{
			name:    "Valid input without TreatmentClassUpgrade",
			input:   `{"klsRawatHak": "1", "klsRawatNaik": "", "pembiayaan": "", "penanggungJawab": ""}`,
			wantErr: false,
		},
		{
			name:    "Invalid input missing Financing with TreatmentClassUpgrade",
			input:   `{"klsRawatHak": "1", "klsRawatNaik": "2", "pembiayaan": "", "penanggungJawab": "1"}`,
			wantErr: true,
		},
		{
			name:    "Invalid input missing PIC with TreatmentClassUpgrade",
			input:   `{"klsRawatHak": "1", "klsRawatNaik": "2", "pembiayaan": "1", "penanggungJawab": ""}`,
			wantErr: true,
		},
	}

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tc TreatmentClass
			err := json.Unmarshal([]byte(tt.input), &tc)
			if err != nil {
				t.Fatalf("Failed to unmarshal input: %v", err)
			}

			err = validate.Struct(tc)
			if (err != nil) != tt.wantErr {
				log.Println("Test case index: ", index)
				t.Errorf("Validation error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
