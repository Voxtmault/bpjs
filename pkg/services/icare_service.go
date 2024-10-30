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

type IcareService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.Icare = &IcareService{}

func (s *IcareService) FKRTLIcare(ctx context.Context, obj *models.FKRTLRequest) (interface{}, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.ICarePath
	method := http.MethodPost

	baseUrl = fmt.Sprintf(
		"%s/api/rs/validate",
		baseUrl,
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

	resp, err := s.HttpHandler.SendRequestIcare(ctx, req)
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

	var responseData map[string]interface{}
	if err = json.Unmarshal([]byte(resp), &responseData); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return responseData, nil
}
