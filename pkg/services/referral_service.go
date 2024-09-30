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

// # According to the documentation there are 3 ways to search for referrals
//
// 1. Is by referral number, which will return only 1 item from BPJS
//
// 2. Is by bpjs id / number, whill will return only 1 item from BPJS
//
// 3. Is by bpjs id / number, but instead it will return multiple item from BPJS
//
// # There are 2 type of sources for referrals
//
// 1. From PCare
//
// 2. From Hospitals
//
// Each source will use different endpoint but the same models are returned
type ReferralService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.Referral = &ReferralService{}

func (s *ReferralService) GetParticipantReferralByReferralNumber(ctx context.Context, referalNumber string, source uint) ([]*models.Referral, error) {
	arrObj := []*models.Referral{}

	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	if source == models.PCareSource {
		baseUrl += "/Rujukan/" + referalNumber
	} else if source == models.HospitalSource {
		baseUrl += "/Rujukan/RS/" + referalNumber
	} else {
		return nil, eris.New("invalid source")
	}

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	var obj models.Referral
	if err = json.Unmarshal([]byte(resp), &obj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	arrObj = append(arrObj, &obj)

	return arrObj, nil
}

func (s *ReferralService) GetParticipantReferralByBPJSNumber(ctx context.Context, bpjsNumber string, source uint, multi bool) ([]*models.Referral, error) {
	arrObj := []*models.Referral{}

	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	if source == models.PCareSource {
		baseUrl += "/Rujukan"
	} else if source == models.HospitalSource {
		baseUrl += "/Rujukan/RS"
	} else {
		return nil, eris.New("invalid source")
	}

	if multi {
		baseUrl += "/List/Peserta" + bpjsNumber
	} else {
		baseUrl += "/Peserta/" + bpjsNumber
	}

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	var obj models.Referral
	if err = json.Unmarshal([]byte(resp), &obj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	arrObj = append(arrObj, &obj)

	return arrObj, nil
}

func (s *ReferralService) GetReferedSpecialist(ctx context.Context, referedHealthFacilityCode, referalDate string) ([]*models.ReferredSpecialist, error) {
	arrObj := models.ReferredSpecialistResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Rujukan/ListSpesialistik/PPKRujukan/%s/TglRujukan/%s",
		baseUrl, referedHealthFacilityCode, referalDate,
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

func (s *ReferralService) GetReferedFacilities(ctx context.Context, referedHealthFacilityCode string) ([]*models.ReferredFacility, error) {
	arrObj := models.ReferedFacilityResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Rujukan/ListSarana/PPKRujukan/%s",
		baseUrl, referedHealthFacilityCode,
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

func (s *ReferralService) CreateReferral(ctx context.Context, obj *models.ReferralCreate) (*models.ReferralCreateResponse, error) {
	referral := models.ReferralCreateResponseWrapper{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/Rujukan/2.0/insert"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.ReferralCreateWrapper{
			TReferral: obj,
		},
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
			return referral.Referral, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return referral.Referral, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &referral); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return referral.Referral, nil
}
