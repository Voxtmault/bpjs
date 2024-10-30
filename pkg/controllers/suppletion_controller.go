package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
)

type SuppletionController struct {
	service   interfaces.SuplesiJasaRaharja
	validator echo.Validator
}

func NewSuppletionController(service interfaces.SuplesiJasaRaharja, validator echo.Validator) *SuppletionController {
	return &SuppletionController{
		service:   service,
		validator: validator,
	}
}

func (s SuppletionController) GetJasaRaharjaSuppletion(c echo.Context) error {
	var res Response

	var obj models.SEPSuppletionParams
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if errMap := utils.MangleValidateResult(s.validator.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.Suplesi(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetJasaRaharjaSuppletion", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SuppletionController) GetAccidentMasterData(c echo.Context) error {
	var res Response

	var obj models.SEPTrafficAccidentParams
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if errMap := utils.MangleValidateResult(s.validator.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.AccidentMasterData(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetAccidentMasterData", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
