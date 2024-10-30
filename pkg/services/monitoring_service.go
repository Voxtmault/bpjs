package services

import (
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

type Monitoring struct {
	HttpHandler interfaces.RequestHandler
}

var _ interfaces.Monitoring = &Monitoring{}

func NewMonitoringService(httpHandler interfaces.RequestHandler) *Monitoring {
	return &Monitoring{
		HttpHandler: httpHandler,
	}
}

func (s *Monitoring) GetMonitoringDataKunjungan(ctx context.Context, tanggalPelayanan, jenisPelayanan string) ([]*models.MonitoringDataKunjungan, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Monitoring/Kunjungan/Tanggal/%s/JnsPelayanan/%s",
		baseUrl, tanggalPelayanan, jenisPelayanan,
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
			return []*models.MonitoringDataKunjungan{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.MonitoringDataKunjungan{}, eris.New("a")
	}

	var sep models.MonitoringDataKunjunganResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Sep, nil
}

//Jenis Pelayanan (1. Inap 2. Jalan)
//Status Klaim (1. Proses Verifikasi 2. Pending Verifikasi 3. Klaim)

func (s *Monitoring) GetMonitoringDataKlaim(ctx context.Context, tanggalPelayanan, jenisPelayanan, statusKlaim string) ([]*models.DataKlaim, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/Monitoring/Klaim/Tanggal/%s/JnsPelayanan/%s/Status/%s",
		baseUrl, tanggalPelayanan, jenisPelayanan, statusKlaim,
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
			return []*models.DataKlaim{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.DataKlaim{}, eris.New("a")
	}

	var sep models.MonitoringKlaimResponse
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Klaim, nil
}

func (s *Monitoring) GetMonitoringHistoryPelayananPeserta(ctx context.Context, noKartu, tanggalMulai, tanggalAkhir string) ([]*models.MonitoringDataKunjungan, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/monitoring/HistoriPelayanan/NoKartu/%s/tglMulai/%s/tglAkhir/%s",
		baseUrl, noKartu, tanggalMulai, tanggalAkhir,
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
			return []*models.MonitoringDataKunjungan{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.MonitoringDataKunjungan{}, eris.New("a")
	}

	var sep models.GetMonitoringHistoryPelayananPeserta
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Sep, nil
}

func (s *Monitoring) GetMonitoringKlaimJaminanJasaraharja(ctx context.Context, jenisPelayanan, tanggalMulai, tanggalAkhir string) ([]*models.DataJasaraharja, error) {
	baseUrl := config.GetConfig().BPJSConfig.BPJSURL + config.GetConfig().BPJSConfig.VClaimPath
	method := http.MethodGet

	baseUrl = fmt.Sprintf(
		"%s/monitoring/JasaRaharja/JnsPelayanan/%s/tglMulai/%s/tglAkhir/%s",
		baseUrl, jenisPelayanan, tanggalMulai, tanggalAkhir,
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
			return []*models.DataJasaraharja{}, eris.Wrap(eris.New(resp), "BPJS Message")
		} else {
			return nil, eris.Wrap(err, "failed to send http request")
		}
	}

	log.Println("Response: ", resp)

	if resp == "" {
		return []*models.DataJasaraharja{}, eris.New("a")
	}

	var sep models.ResponseKlaimJaminanJasaraharja
	if err = json.Unmarshal([]byte(resp), &sep); err != nil {
		return nil, eris.Wrap(err, "failed to unmarshal response")
	}

	return sep.Data, nil
}
