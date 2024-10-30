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

type RuanganService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.RuanganAplicares = &RuanganService{}

func NewRuanganService(httpHandler interfaces.RequestHandler) *RuanganService {
	return &RuanganService{
		HttpHandler: httpHandler,
	}
}

func (s *RuanganService) GetReferensiJenisKamar(ctx context.Context) ([]*models.ReferenceJenisKamar, error) {
	baseUrl := config.GetConfig().BPJSConfig.AplicaresUrl + config.GetConfig().BPJSConfig.AplicaresPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/rest/ref/kelas",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.HttpHandler.SendRequestAplicares(ctx, req)
	if err != nil {
		if resp != "" {
			return []*models.ReferenceJenisKamar{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.ReferenceJenisKamar{}, eris.New("a")
	}

	var sep models.AplicaresRequestResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}

func (s *RuanganService) PostRuangan(ctx context.Context, obj *models.Ruangan, action string) (interface{}, error) {
	//action : create,update,delete
	baseUrl := config.GetConfig().BPJSConfig.AplicaresUrl + config.GetConfig().BPJSConfig.AplicaresPath
	ppkCode := config.GetConfig().BPJSConfig.PPKCode
	method := http.MethodPost

	baseUrl = fmt.Sprintf(
		"%s/rest/bed/%s/%s",
		baseUrl, action, ppkCode,
	)

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

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.HttpHandler.SendRequestAplicares(ctx, req)
	if err != nil {
		if resp != "" {
			return nil, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return nil, eris.New("a")
	}

	var sep models.AplicaresRequestResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}

func (s *RuanganService) GetKetersediaanKamar(ctx context.Context, start, limit string) ([]*models.RuanganResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.AplicaresUrl + config.GetConfig().BPJSConfig.AplicaresPath
	ppkCode := config.GetConfig().BPJSConfig.PPKCode
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/rest/bed/read/%s/%s/%s",
		baseUrl, ppkCode, start, limit,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.HttpHandler.SendRequestAplicares(ctx, req)
	if err != nil {
		if resp != "" {
			return []*models.RuanganResponse{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.RuanganResponse{}, eris.New("a")
	}

	var sep models.AplicaresRuanganResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}
