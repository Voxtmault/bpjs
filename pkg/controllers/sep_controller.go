package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
)

type SEPController struct {
	service   interfaces.SEP
	validator echo.Validator
}

func NewSEPController(service interfaces.SEP, validator echo.Validator) *SEPController {
	return &SEPController{
		service:   service,
		validator: validator,
	}
}

func (s SEPController) Get(c echo.Context) error {
	var res Response

	sepNumber := c.Param("sep_number")
	if sepNumber == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No. SEP tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetSEP(c.Request().Context(), sepNumber)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> get", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
