package services

import (
	"context"
	"encoding/json"
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
