package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type ICareController struct {
	service  interfaces.Icare
	validate echo.Validator
}

func NewICareController(service interfaces.Icare, validate echo.Validator) *ICareController {
	return &ICareController{
		service:  service,
		validate: validate,
	}
}

func (s ICareController) FKRTLIcare(c echo.Context) error {
	var res Response
	var obj models.FKRTLRequest
	var err error
	obj.Param = c.Param("no_card")

	if obj.Param == "" || obj.Param == ":no_card" {
		res.Message = "no_card tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.KodeDokter, err = strconv.Atoi(c.Param("kode_dokter"))
	if err != nil {
		res.Message = "kode_dokter tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	if obj.KodeDokter == 0 {
		res.Message = "kode_dokter tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	res.Data, err = s.service.FKRTLIcare(c.Request().Context(), &obj)
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
