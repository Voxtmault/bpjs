package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

func (s *SEPService) GetReferensiKamar(ctx context.Context) ([]*models.ReferenceKamar, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/aplicaresws/rest/ref/kelas",
		baseUrl,
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
			return []*models.ReferenceKamar{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.ReferenceKamar{}, eris.New("a")
	}

	var sep models.AplicaresRequestResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}


func (s *SEPService) PostKamar(ctx context.Context) ([]*models.ReferenceKamar, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/aplicaresws/rest/ref/kelas",
		baseUrl,
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
			return []*models.ReferenceKamar{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.ReferenceKamar{}, eris.New("a")
	}

	var sep models.AplicaresRequestResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}
