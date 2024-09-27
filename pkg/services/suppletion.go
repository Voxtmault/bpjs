package services

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type SuppletionService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.SuplesiJasaRaharja = &SuppletionService{}

func (s *SuppletionService) Suplesi(ctx context.Context, params *models.SEPSuppletionParams) ([]*models.SEPSuppletion, error) {
	arrObj := models.SEPSuppletionResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	// Default the Service Date if empty to today
	if params.ServiceDate == "" {
		params.ServiceDate = time.Now().Format(time.DateOnly)
	}

	baseUrl += "/sep/JasaRaharja/Suplesi/" + params.BPJSNumber + "/tglPelayanan/" + params.ServiceDate

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.Guarantee, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj.Guarantee, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj.Guarantee, nil
}

func (s *SuppletionService) AccidentMasterData(ctx context.Context, params *models.SEPTrafficAccidentParams) ([]*models.SEPTrafficAccident, error) {
	arrObj := models.SEPTrafficAccidentResponse{}
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	if params.BPJSNumber == "" {
		return nil, eris.New("BPJS Number is required")
	}

	baseUrl += "/sep/KllInduk/List/" + params.BPJSNumber

	log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return arrObj.List, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return arrObj.List, nil
	} else {
		if err = json.Unmarshal([]byte(resp), &arrObj); err != nil {
			return nil, eris.Wrap(err, "failed to unmarshal response")
		}
	}

	return arrObj.List, nil
}
