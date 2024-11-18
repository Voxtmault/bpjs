package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/logger"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type RequestHandlerService struct {
	Security      interfaces.BPJSSecurity
	RequestLogger *logger.RequestLogger
}

var _ interfaces.RequestHandler = &RequestHandlerService{}

func NewBPJSRequestHandlerService(security interfaces.BPJSSecurity) *RequestHandlerService {
	return &RequestHandlerService{
		Security:      security,
		RequestLogger: logger.GetRequestLogger(),
	}
}

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

	// If none is set from the caller then use the default
	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/json")
	}

	// Add custom headers
	req.Header.Add("X-cons-id", cfg.ConsumerID)
	req.Header.Add("X-timestamp", fmt.Sprintf("%d", timeStamp))
	req.Header.Add("X-signature", signature)
	req.Header.Add("user_key", cfg.Userkey)

	var requestBody []byte
	if req.Body != nil {
		requestBody, err = io.ReadAll(req.Body)
		if err != nil {
			return "", eris.Wrap(err, "failed to read request body")
		}
		// Reset the request body so it can be read again by the HTTP client
		req.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}

	// Send the request
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", eris.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	slog.Debug("request handler -> SendRequest", "received http status code", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", eris.Wrap(err, "failed to read response body")
	}
	// log.Println("Response: ", string(body))

	// Unmarshall into response obj

	// Unmarshal into response obj
	var response models.BPJSResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", eris.Wrap(err, "failed to unmarshal response")
	}

	if response.MetaData.Code != "200" {
		slog.Debug("request handler -> SendRequest", "received response", response.MetaData)
		// If the response code is not 200, return the error message

		// If the response message is "Data Tidak Ada" or something similar, you can treat this as a 404 response code
		// IDK why they insists in returning code 201 :/

		// [Update] they will occasionally use the MetaData.Code as the error message, it's weird man, i tell you

		// Log the request and response
		req.Body = io.NopCloser(bytes.NewBuffer(requestBody))   // Restore the request body
		resp.Body = io.NopCloser(bytes.NewBuffer([]byte(body))) // Restore the response body

		if err = s.RequestLogger.LogEggressRequest(ctx, req, resp); err != nil {
			return "", eris.Wrap(err, "failed to log egress request")
		}

		return response.MetaData.Message, eris.New(response.MetaData.Code)
	}

	// For some cases, the response is empty or null
	if response.Response == "" {
		return "", nil
	}

	// Decrypt the response
	raw, err := s.Security.DecryptResponse(ctx, timeStamp, response.Response)
	if err != nil {
		return "", eris.Wrap(err, "failed to decrypt response")
	}

	// Log the request and response
	req.Body = io.NopCloser(bytes.NewBuffer(requestBody))  // Restore the request body
	resp.Body = io.NopCloser(bytes.NewBuffer([]byte(raw))) // Restore the response body

	if err = s.RequestLogger.LogEggressRequest(ctx, req, resp); err != nil {
		return "", eris.Wrap(err, "failed to log egress request")
	}

	return raw, nil
}

func (s *RequestHandlerService) SendRequestIcare(ctx context.Context, req *http.Request) (string, error) {
	cfg := config.GetConfig().BPJSConfig
	// Logic
	// 1. For BPJS, add custom headers before sending the request
	// 2. After receiving the response, decrypt the response before sending it to the caller

	timeStamp := time.Now().UTC().Unix()
	signature, err := s.Security.CreateSignature(ctx, timeStamp)
	if err != nil {
		return "", eris.Wrap(err, "failed to create signature")
	}

	// If none is set from the caller then use the default
	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/json")
	}

	// Add custom headers
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

	log.Println("Status Code: ", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", eris.Wrap(err, "failed to read response body")
	}
	log.Println("Response: ", string(body))

	// Unmarshall into response obj
	var response models.IcareResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", eris.Wrap(err, "failed to unmarshall response")
	}

	if response.MetaData.Code != 200 {
		log.Println("Response: ", response.MetaData)
		// If the response code is not 200, return the error message

		// If the response message is "Data Tidak Ada" or something similar, you can treat this as a 404 response code
		// IDK why they insists in returning code 201 :/

		return response.MetaData.Message, eris.New(strconv.Itoa(response.MetaData.Code))
	}

	// For some cases, the response is empty or null
	if response.Response == "" {
		return "", nil
	}

	// Decrypt the response
	raw, err := s.Security.DecryptResponse(ctx, timeStamp, response.Response)
	if err != nil {
		return "", eris.Wrap(err, "failed to decrypt response")
	}

	return raw, nil
}

func (s *RequestHandlerService) SendRequestAplicares(ctx context.Context, req *http.Request) (string, error) {
	cfg := config.GetConfig().BPJSConfig
	// Logic
	// 1. For BPJS, add custom headers before sending the request
	// 2. After receiving the response, decrypt the response before sending it to the caller

	timeStamp := time.Now().UTC().Unix()
	signature, err := s.Security.CreateSignature(ctx, timeStamp)
	if err != nil {
		return "", eris.Wrap(err, "failed to create signature")
	}

	// If none is set from the caller then use the default
	if req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/json")
	}

	// Add custom headers
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

	log.Println("Status Code: ", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", eris.Wrap(err, "failed to read response body")
	}
	log.Println("Response: ", string(body))

	// Unmarshall into response obj
	var response models.AplicaresResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", eris.Wrap(err, "failed to unmarshall response")
	}

	if response.MetaData.Code != 1 {
		log.Println("Response: ", response.MetaData)
		// If the response code is not 200, return the error message

		// If the response message is "Data Tidak Ada" or something similar, you can treat this as a 404 response code
		// IDK why they insists in returning code 201 :/

		return response.MetaData.Message, eris.New(strconv.Itoa(response.MetaData.Code))
	}

	// // For some cases, the response is empty or null
	// if response.Response == "" {
	// 	return "", nil
	// }

	// Decrypt the response
	raw, err := json.Marshal(response.Response)
	if err != nil {
		return "", eris.Wrap(err, "failed to marshall response")
	}
	return string(raw), nil
}
