package services

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type BPJSParticipantService struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.Participant = &BPJSParticipantService{}

func (s *BPJSParticipantService) GetParticipant(ctx context.Context, query *models.ParticipantSearchParams) (*models.BPJSParticipant, error) {

	// Logic
	// 1. Build the http request according to the BPJS docs
	// 2. Send it through the http handler service
	// 3. Process the response

	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	if query.ServiceDate == "" {
		query.ServiceDate = time.Now().Format(time.DateOnly)
	}

	if query.BPJSNumber != "" {
		// Search By BPJS Number
		baseUrl += "/Peserta/nokartu/" + query.BPJSNumber + "/tglSEP/" + query.ServiceDate
	} else if query.NIK != "" {
		// Search by NIK
		baseUrl += "/Peserta/nik/" + query.NIK + "/tglSEP/" + query.ServiceDate
	} else {
		// Invalid query
		return &models.BPJSParticipant{}, eris.New("invalid query params")
	}

	// log.Println("URL: ", baseUrl)

	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, eris.Wrap(err, "failed to create http request")
	}

	resp, err := s.HttpHandler.SendRequest(ctx, req)
	if err != nil {
		if resp != "" {
			return &models.BPJSParticipant{}, eris.Wrap(eris.New(resp), "failed to send http request")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	var obj models.BPJSParticipantResponse
	if err = json.Unmarshal([]byte(resp), &obj); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return obj.Participant, nil
}
