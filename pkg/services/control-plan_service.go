package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type ControlPlanService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.ControlPlan = &ControlPlanService{}

func (s *ControlPlanService) GetViaSEP(ctx context.Context, sepNumber string) ([]*models.ControlPlanGetViaSEP, error) {
	arrObj := []*models.ControlPlanGetViaSEP{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl += "/RencanaKontrol/nosep/" + sepNumber

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj, nil
}

func (s *ControlPlanService) GetViaControlLetterNumber(ctx context.Context, letterNumber string) ([]*models.ControlPlanGetViaControllLetterNumber, error) {
	arrObj := []*models.ControlPlanGetViaControllLetterNumber{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl += "/RencanaKontrol/noSuratKontrol/" + letterNumber

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj, nil
}

func (s *ControlPlanService) GetControlPlanFromCardNumber(ctx context.Context, params *models.ControlPlansFromCardNumberParams) ([]*models.ControlPlans, error) {
	arrObj := models.ControlPlansResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/RencanaKontrol/ListRencanaKontrol/Bulan/%s/Tahun/%s/Nokartu/%s/filter/%s",
		baseUrl, params.Month, params.Year, params.CardNumber, params.Filter,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Lists, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj.Lists, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj.Lists, nil
}

func (s *ControlPlanService) GetControlPlans(ctx context.Context, params *models.ControlPlanParams) ([]*models.ControlPlans, error) {
	arrObj := models.ControlPlansResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/RencanaKontrol/ListRencanaKontrol/tglAwal/%s/tglAkhir/%s/filter/%s",
		baseUrl, params.StartDate, params.EndDate, params.Filter,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Lists, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj.Lists, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj.Lists, nil
}

func (s *ControlPlanService) GetClinicControlPlans(ctx context.Context, params *models.ClinicControlParams) ([]*models.ClinicControlPlans, error) {
	arrObj := models.ClinicControlPlansResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/RencanaKontrol/ListSpesialistik/JnsKontrol/%s/nomor/%s/TglRencanaKontrol/%s",
		baseUrl, params.ControlType, params.Identifier, params.ControlPlannedDate,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Lists, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj.Lists, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj.Lists, nil
}

func (s *ControlPlanService) GetDoctorPracticeSchedule(ctx context.Context, params *models.DoctorScheduleParams) ([]*models.DoctorPracticeSchedule, error) {
	arrObj := models.DoctorPracticeScheduleResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/RencanaKontrol/JadwalPraktekDokter/JnsKontrol/%s/KdPoli/%s/TglRencanaKontrol/%s",
		baseUrl, params.ControlType, params.PoliCode, params.ControlPlannedDate,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Lists, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj.Lists, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj.Lists, nil
}

func (s *ControlPlanService) CreateControlPlan(ctx context.Context, obj *models.ControlPlanCreate) (*models.ControlPlanCreateResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/RencanaKontrol/insert"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: obj,
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to marshal object")
	}

	log.Println("JSON Data: ", string(jsonData))

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return nil, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	var controlResponse models.ControlPlanCreateResponse
	if err = json.Unmarshal([]byte(resp), &controlResponse); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &controlResponse, nil
}

func (s *ControlPlanService) UpdateControlPlan(ctx context.Context, obj *models.UpdateControlPlans) (*models.ControlPlanCreateResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl += "/RencanaKontrol/Update"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: obj,
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to marshal object")
	}

	log.Println("JSON Data: ", string(jsonData))

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return nil, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	var controlResponse models.ControlPlanCreateResponse
	if err = json.Unmarshal([]byte(resp), &controlResponse); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &controlResponse, nil
}

func (s *ControlPlanService) DeleteControlPlan(ctx context.Context, controlNumber, user string) error {

	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodDelete

	baseUrl += "/RencanaKontrol/Delete"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: models.TControlPlan{
			TControlPlan: struct {
				ControlLetterNumber string `json:"noSuratKontrol"`
				User                string `json:"user"`
			}{
				ControlLetterNumber: controlNumber,
				User:                user,
			},
		},
	})
	if err != nil {
		return eris.Wrap(err, "failed to marshal object")
	}

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	// For Delete Control Plan, BPJS does not return any message

	return nil
}

func (s *ControlPlanService) CreateInpatientCareOrder(ctx context.Context, obj *models.ControlPlanCreate) (*models.ControlPlanCreateResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/RencanaKontrol/InsertSPRI"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: obj,
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to marshal object")
	}

	log.Println("JSON Data: ", string(jsonData))

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return nil, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	var controlResponse models.ControlPlanCreateResponse
	if err = json.Unmarshal([]byte(resp), &controlResponse); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &controlResponse, nil
}

func (s *ControlPlanService) UpdateInpatientCareOrder(ctx context.Context, obj *models.UpdateControlPlans) (*models.ControlPlanCreateResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl += "/RencanaKontrol/UpdateSPRI"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: obj,
	})
	if err != nil {
		return nil, eris.Wrap(err, "failed to marshal object")
	}

	log.Println("JSON Data: ", string(jsonData))

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return nil, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	var controlResponse models.ControlPlanCreateResponse
	if err = json.Unmarshal([]byte(resp), &controlResponse); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &controlResponse, nil
}
