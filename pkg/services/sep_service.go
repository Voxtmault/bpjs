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
		if resp != "" {
			return &models.SEPCreateResponse{}, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	// log.Println("Response: ", resp)

	var sep models.SEPCreateResponse
	if err = json.Unmarshal([]byte(resp), &obj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &sep, nil
}

func (s *SEPService) UpdateSEP(ctx context.Context) (any, error) {
	panic("not implemented") // TODO: Implement
}

func (s *SEPService) DeleteSEP(ctx context.Context) (any, error) {
	panic("not implemented") // TODO: Implement
}
