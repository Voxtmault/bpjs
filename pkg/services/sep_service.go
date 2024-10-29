package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
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

func NewSEPService(httpHandler interfaces.RequestHandler) *SEPService {
	return &SEPService{
		HttpHandler: httpHandler,
	}
}

func (s *SEPService) InsertSEP(ctx context.Context, obj *models.SEPCreate) (*models.SEPCreateResponse, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/SEP/2.0/insert"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
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

	log.Println("Response: ", resp)

	var sep models.SEPCreateResponseWrapper
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.SEP, nil
}

func (s *SEPService) UpdateSEP(ctx context.Context, obj *models.SEPUpdate) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl += "/SEP/2.0/update"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
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

	jsonData, err := json.Marshal(models.BPJSRequest{
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

	slog.Debug("get sep", "url", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return &models.SEPGet{}, eris.Wrap(err, "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return &models.SEPGet{}, eris.New("something went wrong")
	}

	var sep models.SEPGet
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &sep, nil
}

func (s *SEPService) RequestSEP(ctx context.Context, obj *models.SEPRequestCreate) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/Sep/pengajuanSEP"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
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
	log.Println("Response: ", resp)

	return resp, nil
}

func (s *SEPService) ApprovalSEPRequest(ctx context.Context, obj *models.SEPRequestCreate) (string, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl += "/Sep/aprovalSEP"

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
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
	log.Println("Response: ", resp)

	return resp, nil
}

func (s *SEPService) GetSEPRequests(ctx context.Context, month, year string) ([]*models.SEPRequest, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Sep/persetujuanSEP/list/bulan/%s/tahun/%s",
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
			return []*models.SEPRequest{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.SEPRequest{}, eris.New("a")
	}

	var sep models.SEPRequestResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}

func (s *SEPService) UpdateTanggalPulang(ctx context.Context, obj *models.SEPUpdateTanggalPulangRequest) (interface{}, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPut

	baseUrl = fmt.Sprintf(
		"%s/SEP/2.0/updtglplg",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
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
		if resp != "" {
			return []*models.SEPRequest{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.SEPRequest{}, eris.New("a")
	}

	var sep map[string]interface{}
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep, nil
}

func (s *SEPService) GetFingerPrintSEP(ctx context.Context, noKartu, tanggalPelayanan string) (*models.SEPGetFingerPrint, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/SEP/FingerPrint/Peserta/%s/TglPelayanan/%s",
		baseUrl, noKartu, tanggalPelayanan,
	)

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return &models.SEPGetFingerPrint{}, eris.Wrap(err, "failed to create http request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return &models.SEPGetFingerPrint{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return &models.SEPGetFingerPrint{}, eris.New("a")
	}

	var sep models.SEPGetFingerPrint
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return &sep, nil
}

func (s *SEPService) GetListFIngerPrintSEP(ctx context.Context, tanggalPelayanan string) ([]*models.SEPGetListFingerPrint, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/SEP/FingerPrint/List/Peserta/TglPelayanan/%s",
		baseUrl, tanggalPelayanan,
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
			return []*models.SEPGetListFingerPrint{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.SEPGetListFingerPrint{}, eris.New("a")
	}

	var sep models.SEPFingerPrintResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}

func (s *SEPService) GetListRandomQuestion(ctx context.Context, noKartu, tanggalPelayanan string) ([]*models.SEPGetRandomQuestion, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/SEP/FingerPrint/randomquestion/faskesterdaftar/nokapst/%s/tglsep/%s",
		baseUrl, noKartu, tanggalPelayanan,
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
			return []*models.SEPGetRandomQuestion{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.SEPGetRandomQuestion{}, eris.New("a")
	}

	var sep models.SEPRandomQuestionResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Lists, nil
}

func (s *SEPService) PostRandomQuestion(ctx context.Context, obj *models.PostRequestRandomQuestion) (bool, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodPost

	baseUrl = fmt.Sprintf(
		"%s/SEP/FingerPrint/randomanswer",
		baseUrl,
	)

	log.Println("URL: ", baseUrl)

	jsonData, err := json.Marshal(models.BPJSRequest{
		Request: &models.TSEP{
			TSEP: obj,
		},
	})
	if err != nil {
		return false, eris.Wrap(err, "failed to marshal object")
	}

	log.Println("JSON Data: ", string(jsonData))

	req, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, eris.Wrap(err, "failed to create http request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return false, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return false, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	// if resp == "" {
	// 	return false, eris.New("a")
	// }

	// var sep map[string]interface{}
	// if err = json.Unmarshal([]byte(resp), &sep); err != nil {
	// 	return nil, eris.Wrap(err, "failed to unmarshal response")
	// }

	return false, nil
}

