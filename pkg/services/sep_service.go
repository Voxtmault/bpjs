package services

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type SEPService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.SEP = &SEPService{}

func (s *SEPService) InsertSEP(ctx context.Context, obj *models.SEPCreate) (*models.SEPCreateResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/SEP/2.0/insert"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSSEP{
		Request: &models.TSEP{
			TSEP: obj,
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
		// Meaning that BPJS has a custom message for this case
		// BPJS TOLD US to not meddle with the message, so we're just going to return it as is
		// with no Filter...
		if resp != "" {
			return &models.SEPCreateResponse{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	// log.Println("Response: ", resp)

	var sep models.SEPCreateResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &sep, nil
}

func (s *SEPService) UpdateSEP(ctx context.Context, obj *models.SEPUpdate) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl += "/SEP/2.0/update"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSSEP{
		Request: &models.TSEP{
			TSEP: obj,
		},
	})
	if err != nil {
		return "", eris.Wrap(err, "failed to marshal object")
	}

	log.Println("JSON Data: ", string(jsonData))

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		// Meaning that BPJS has a custom message for this case
		// BPJS TOLD US to not meddle with the message, so we're just going to return it as is
		// with no Filter...
		if resp != "" {
			return "", eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	// BPJS Response with SEP Number
	// TODO find out if the SEP Number from the BPJS Response is the same or different, since we can't test BPJS SEP Service because we do not have SIO

	log.Println("Response: ", resp)

	return resp, nil
}

func (s *SEPService) DeleteSEP(ctx context.Context, obj *models.SEPDelete) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodDelete

	baseUrl += "/SEP/2.0/delete"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSSEP{
		Request: &models.TSEP{
			TSEP: obj,
		},
	})
	if err != nil {
		return "", eris.Wrap(err, "failed to marshal object")
	}

	log.Println("JSON Data: ", string(jsonData))

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		// Meaning that BPJS has a custom message for this case
		// BPJS TOLD US to not meddle with the message, so we're just going to return it as is
		// with no Filter...
		if resp != "" {
			return "", eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	// BPJS Response with SEP Number
	// TODO find out if the SEP Number from the BPJS Response is the same or different, since we can't test BPJS SEP Service because we do not have SIO

	log.Println("Response: ", resp)

	return resp, nil
}

func (s *SEPService) GetSEP(ctx context.Context, sepNumber string) (*models.SEPGet, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl += "/SEP/" + sepNumber

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return &models.SEPGet{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return &models.SEPGet{}, eris.New("a")
	}

	var sep models.SEPGet
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &sep, nil
}
