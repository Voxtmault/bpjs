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

type PRBService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.PRB = &PRBService{}

func NewPRBService(httpHandler interfaces.RequestHandler) *PRBService {
	return &PRBService{
		HttpHandler: httpHandler,
	}
}

func (s *PRBService) PRBInsert(ctx context.Context, obj *models.PRBInsertRequest) (*models.PRBResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl = fmt.Sprintf(
		"%s/PRB/insert",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.TPRB{
			TPrb: obj,
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
			return nil, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return nil, eris.New("a")
	}

	var sep models.PRBResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &sep, nil
}

func (s *PRBService) PRBUpdate(ctx context.Context, obj *models.PRBUpdateRequest) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl = fmt.Sprintf(
		"%s/PRB/update",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.TPRB{
			TPrb: obj,
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
		if resp != "" {
			return "", eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return "", eris.New("a")
	}

	// var sep models.MonitoringDataKunjunganResponse
	// if err = json.Unmarshal([]byte(resp), &sep); err != nil {
	// 	return "", eris.Wrap(err, "failed to unmarshal response")
	// }

	return resp, nil
}

func (s *PRBService) PRBDelete(ctx context.Context, obj *models.PRBDeleteRequest) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodDelete

	baseUrl = fmt.Sprintf(
		"%s/PRB/delete",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.TPRB{
			TPrb: obj,
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
		if resp != "" {
			return "", eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return "", eris.New("a")
	}

	// var sep models.MonitoringDataKunjunganResponse
	// if err = json.Unmarshal([]byte(resp), &sep); err != nil {
	// 	return "", eris.Wrap(err, "failed to unmarshal response")
	// }

	return resp, nil
}

func (s *PRBService) GetPRBbyNomorSRB(ctx context.Context, noPrb, noSep string) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/prb/%s/nosep/%s",
		baseUrl, noPrb, noSep,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return "", eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return "", eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return "", eris.New("a")
	}

	var sep models.PRBResponseGetByNoSRB
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return "", eris.Wrap(err, "failed to unmarshal response")
	}

	return resp, nil
}

func (s *PRBService) GetPRBbyTanggal(ctx context.Context, tanggalAwal, tanggalAkhir string) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/prb/tglMulai/%s/tglAkhir/%s",
		baseUrl, tanggalAwal, tanggalAkhir,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return "", eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return "", eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return "", eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return "", eris.New("a")
	}

	var sep models.PRBResponseGetByTanggal
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return "", eris.Wrap(err, "failed to unmarshal response")
	}

	return resp, nil
}
