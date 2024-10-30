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

type LPKService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.LPK = &LPKService{}

func NewLPKService(httpHandler interfaces.RequestHandler) *LPKService {
	return &LPKService{
		HttpHandler: httpHandler,
	}
}

func (s *LPKService) LPKInsert(ctx context.Context, obj *models.InsertLPKRequest) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl = fmt.Sprintf(
		"%s/LPK/insert",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.Tlpk{
			TLpk: obj,
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

func (s *LPKService) LPKUpdate(ctx context.Context, obj *models.InsertLPKRequest) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl = fmt.Sprintf(
		"%s/LPK/update",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.Tlpk{
			TLpk: obj,
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

func (s *LPKService) LPKDelete(ctx context.Context, noSep string) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodDelete

	baseUrl = fmt.Sprintf(
		"%s/LPK/delete",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.Tlpk{
			TLpk: map[string]interface{}{
				"noSep": noSep,
			},
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

func (s *LPKService) LPKGet(ctx context.Context, tanggalMasuk, jenisLayanan string) ([]*models.ListLPKResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/LPK/TglMasuk/%s/JnsPelayanan/%s",
		baseUrl, tanggalMasuk, jenisLayanan,
	)

	req, err := http.NewRequest(method, baseUrl, nil)
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

	var lpk models.LPKResponseData
	if err = json.Unmarshal([]byte(resp), &lpk); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return lpk.Lpk.List, nil
}
