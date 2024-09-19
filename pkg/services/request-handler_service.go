package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type RequestHandlerService struct {
	Security interfaces.BPJSSecurity
}

var _ interfaces.RequestHandler = &RequestHandlerService{}

func (s *RequestHandlerService) SendRequest(ctx context.Context, req *http.Request) (string, error) {
	cfg := config.GetConfig().BPJSConfig
	// Logic
	// 1. For BPJS, add custom headers before sending the request
	// 2. After receiving the response, decrypt the response before sending it to the caller

	timeStamp := time.Now().UTC().Unix()
	signature, err := s.Security.CreateSignature(ctx, timeStamp)
	if err != nil {
		return "", eris.Wrap(err, "failed to create signature")
	}

	// Add custom headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-cons-id", cfg.ConsumerID)
	req.Header.Add("X-timestamp", fmt.Sprintf("%d", timeStamp))
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", cfg.Userkey)

	// Send the request
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", eris.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", eris.Wrap(err, "failed to read response body")
	}

	// Unmarshall into response obj
	var response models.BPJSResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", eris.Wrap(err, "failed to unmarshall response")
	}

	if response.MetaData.Code != "200" {
		// If the response code is not 200, return the error message
		return response.MetaData.Message, eris.New("failed to get participant")
	}

	// Decrypt the response
	raw, err := s.Security.DecryptResponse(ctx, timeStamp, response.Response)
	if err != nil {
		return "", eris.Wrap(err, "failed to decrypt response")
	}

	return raw, nil
}
