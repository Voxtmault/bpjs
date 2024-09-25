package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type ReferenceService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.Reference = &ReferenceService{}

func (s *ReferenceService) DiagnoseReference(ctx context.Context, diagnosisCode string) ([]*models.Reference, error) {
	arrObj := models.DiagnosisReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/diagnosa"

	if diagnosisCode != "" {
		baseUrl += "/" + diagnosisCode
	}

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Diagnosis, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.Diagnosis, nil
}

func (s *ReferenceService) DoctorReference(ctx context.Context, jenisPelayanan, tglPelayanan, kodeSPesialis string) ([]*models.Reference, error) {
	arrObj := models.DoctorReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/dokter/pelayanan/" + jenisPelayanan + "/tglPelayanan/" + tglPelayanan + "/Spesialis/" + kodeSPesialis

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Doctor, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.Doctor, nil
	// return nil, nil
}

func (s *ReferenceService) PoliclinicsReference(ctx context.Context, poliCode string) ([]*models.Reference, error) {
	arrObj := models.PoliReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/poli"

	if poliCode != "" {
		baseUrl += "/" + poliCode
	}

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Poli, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.Poli, nil
	// return nil, nil
}

func (s *ReferenceService) HealthFacilityReference(ctx context.Context, namaFaskes, jenisFaskes string) ([]*models.Reference, error) {
	arrObj := models.FaskesReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/faskes" + namaFaskes + "/" + jenisFaskes

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Faskes, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.Faskes, nil
	// return nil, nil
}

func (s *ReferenceService) ProcedureReference(ctx context.Context, procedure string) ([]*models.Reference, error) {
	arrObj := models.ProcedureReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/procedure/" + procedure

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Procedure, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.Procedure, nil
}

func (s *ReferenceService) NursingClassReference(ctx context.Context) ([]*models.Reference, error) {
	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/kelasrawat"
	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
	// return nil, nil
}

func (s *ReferenceService) SpecialistReference(ctx context.Context) ([]*models.Reference, error) {
	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/spesialistik"

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
	// return nil, nil
}

func (s *ReferenceService) DischargeMethodReference(ctx context.Context) ([]*models.Reference, error) {

	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/carakeluar"

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
}

func (s *ReferenceService) PostDischargeReference(ctx context.Context) ([]*models.Reference, error) {
	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/pascapulang"

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
}

// Provinsi
func (s *ReferenceService) ProvinceReference(ctx context.Context) ([]*models.Reference, error) {
	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/propinsi"

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
	// return nil, nil
}

// Kabupaten
func (s *ReferenceService) RegencyReference(ctx context.Context, kodeProvince string) ([]*models.Reference, error) {
	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/kabupaten/propinsi/" + kodeProvince

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
}

// Kecamatan
func (s *ReferenceService) DistrictReference(ctx context.Context, kodeKota string) ([]*models.Reference, error) {
	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/kecamatan/kabupaten/" + kodeKota

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
}

// DPJP
func (s *ReferenceService) AttendingPhysicianReference(ctx context.Context, kodeDokter string) ([]*models.Reference, error) {
	arrObj := models.ListReference{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Uses ICD-10 diagnosis code
	baseUrl += "/referensi/dokter/" + kodeDokter

	// if diagnosisCode != "" {
	// 	baseUrl += "/" + diagnosisCode
	// }

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return arrObj.List, nil
}
