package controllers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
)

type MonitoringController struct {
	service  interfaces.Monitoring
	validate echo.Validator
}

func NewMonitoringController(service interfaces.Monitoring, validate echo.Validator) *MonitoringController {
	return &MonitoringController{
		service:  service,
		validate: validate,
	}
}

func (s MonitoringController) GetMonitoringDataKunjungan(c echo.Context) error {
	var res Response
	tanggalPelayanan := c.Param("service_date")

	if tanggalPelayanan == "" || tanggalPelayanan == ":service_date" {
		res.Message = "Tanggal tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, tanggalPelayanan)
		if err != nil {
			res.Message = "Tanggal tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}
	jenisPelayanan := c.Param("service_type")
	if jenisPelayanan != "1" && jenisPelayanan != "2" {
		res.Message = "Jenis Pelayanan tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.GetMonitoringDataKunjungan(c.Request().Context(), tanggalPelayanan, jenisPelayanan)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("monitoring_controller -> GetMonitoringDataKunjungan", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s MonitoringController) GetMonitoringDataKlaim(c echo.Context) error {
	var res Response
	tanggalPulang := c.Param("discharge_date")

	if tanggalPulang == "" || tanggalPulang == ":discharge_date" {
		res.Message = "Tanggal tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, tanggalPulang)
		if err != nil {
			res.Message = "Tanggal tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}
	jenisPelayanan := c.Param("jenis_pelayanan")
	if jenisPelayanan != "1" && jenisPelayanan != "2" {
		res.Message = "Jenis Pelayanan tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	status_klaim := c.Param("status_klaim")
	if jenisPelayanan != "1" && jenisPelayanan != "2" && jenisPelayanan != "3" {
		res.Message = "Jenis Pelayanan tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetMonitoringDataKlaim(c.Request().Context(), tanggalPulang, jenisPelayanan, status_klaim)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("monitoring_controller -> GetMonitoringDataKlaim", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s MonitoringController) GetMonitoringHistoryPelayananPeserta(c echo.Context) error {
	var res Response
	tanggalAwal := c.Param("start_date")

	if tanggalAwal == "" || tanggalAwal == ":start_date" {
		res.Message = "Tanggal awal tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, tanggalAwal)
		if err != nil {
			res.Message = "Tanggal awal tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}
	tanggalAkhir := c.Param("end_date")

	if tanggalAwal == "" || tanggalAwal == ":end_date" {
		res.Message = "Tanggal akhir tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, tanggalAkhir)
		if err != nil {
			res.Message = "Tanggal akhir tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}
	nomorKartu := c.Param("card_number")
	if tanggalAwal == "" || tanggalAwal == ":card_number" {
		res.Message = "Nomor kartu tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetMonitoringHistoryPelayananPeserta(c.Request().Context(), nomorKartu, tanggalAwal, tanggalAkhir)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("monitoring_controller -> GetMonitoringHistoryPelayananPeserta", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s MonitoringController) GetMonitoringKlaimJaminanJasaraharja(c echo.Context) error {
	var res Response
	tanggalAwal := c.Param("start_date")

	if tanggalAwal == "" || tanggalAwal == ":start_date" {
		res.Message = "Tanggal awal tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, tanggalAwal)
		if err != nil {
			res.Message = "Tanggal awal tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}
	tanggalAkhir := c.Param("end_date")

	if tanggalAwal == "" || tanggalAwal == ":end_date" {
		res.Message = "Tanggal akhir tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, tanggalAkhir)
		if err != nil {
			res.Message = "Tanggal akhir tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}

	jenisPelayanan := c.Param("service_type")
	if jenisPelayanan != "1" && jenisPelayanan != "2" {
		res.Message = "Jenis Pelayanan tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.GetMonitoringKlaimJaminanJasaraharja(c.Request().Context(), jenisPelayanan, tanggalAwal, tanggalAkhir)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("monitoring_controller -> GetMonitoringKlaimJaminanJasaraharja", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
