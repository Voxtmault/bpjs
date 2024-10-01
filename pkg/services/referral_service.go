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

func (s *ReferralService) GetOutgoingReferral(ctx context.Context, startDate, endDate string) ([]*models.OutgoingReferral, error) {
	arrObj := models.OutgoingReferralResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Rujukan/Keluar/List/tglMulai/%s/tglAkhir/%s",
		baseUrl, startDate, endDate,
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

func (s *ReferralService) GetOutgoingReferralDetail(ctx context.Context, referralNumber string) (*models.ReferralDetail, error) {
	obj := models.ReferralDetailResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Rujukan/Keluar/%s",
		baseUrl, referralNumber,
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
			return obj.Referral, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return obj.Referral, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &obj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return obj.Referral, nil
}

func (s *ReferralService) CreateReferral(ctx context.Context, obj *models.ReferralAction) (*models.ReferralCreateResponse, error) {
	referral := models.ReferralCreateResponseWrapper{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/Rujukan/2.0/insert"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.ReferralActionWrapper{
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

func (s *ReferralService) UpdateReferral(ctx context.Context, obj *models.ReferralAction) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl += "/Rujukan/2.0/Update"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.ReferralActionWrapper{
			TReferral: obj,
		},
	})
	if err != nil {
		return "", eris.Wrap(err, "failed to marshal object")
	}

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return resp, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	return resp, nil
}

func (s *ReferralService) DeleteReferral(ctx context.Context, referralNumber, user string) error {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodDelete

	baseUrl += "/Rujukan/delete"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: models.ReferralActionWrapper{
			TReferral: struct {
				ReferralNumber string `json:"noRujukan"`
				User           string `json:"user"`
			}{
				ReferralNumber: referralNumber,
				User:           user,
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

	return nil
}

func (s *ReferralService) CreateSpecialReferral(ctx context.Context, obj *models.SpecialReferralCreate) (*models.SpecialReferralCreateResponse, error) {
	referral := models.SpecialReferralCreateResponseWrapper{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/Rujukan/Khusus/insert"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(obj)
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

func (s *ReferralService) GetSpecialReferrals(ctx context.Context, month, year string) ([]*models.SpecialReferrals, error) {
	arrObj := models.SpecialReferralsResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Rujukan/Khusus/List/Bulan/%s/Tahun/%s",
		baseUrl, month, year,
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
			return arrObj.Referrals, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj.Referrals, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj.Referrals, nil
}

func (s *ReferralService) DeleteSpecialReferral(ctx context.Context, obj *models.SpecialReferralDelete) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodDelete

	baseUrl += "/Rujukan/Khusus/delete"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.ReferralActionWrapper{
			TReferral: obj,
		},
	})
	if err != nil {
		return "", eris.Wrap(err, "failed to marshal object")
	}

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return resp, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	return resp, nil
}
