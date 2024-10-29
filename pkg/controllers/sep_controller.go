package controllers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
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

func (s SEPController) Post(c echo.Context) error {
	var res Response

	var obj models.SEPCreate
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.InsertSEP(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> post", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) Put(c echo.Context) error {
	var res Response

	var obj models.SEPUpdate
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.UpdateSEP(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> put", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) Delete(c echo.Context) error {
	var res Response

	var obj models.SEPDelete
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.DeleteSEP(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> delete", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) UpdateDischargeDate(c echo.Context) error {
	var res Response

	var obj models.SEPUpdateTanggalPulangRequest
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := s.service.UpdateTanggalPulang(c.Request().Context(), &obj)
	if err != nil {
		if strings.Contains(err.Error(), "BPJS Message") {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> UpdateDischargeDate", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) GetDischargedSEP(c echo.Context) error {
	var res Response

	month := c.Param("month")
	if month == "" || month == ":month" {
		return echo.NewHTTPError(http.StatusBadRequest, "bulan tidak boleh kosong")
	}

	year := c.Param("year")
	if year == "" || year == ":year" {
		return echo.NewHTTPError(http.StatusBadRequest, "tahun tidak boleh kosong")
	}

	filter := c.Param("filter")

	var err error
	res.Data, err = s.service.GetDischargedSEP(c.Request().Context(), month, year, filter)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetDischargedSEP", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) GetInternalSEP(c echo.Context) error {
	var res Response

	sepNumber := c.Param("sep_number")
	if sepNumber == "" || sepNumber == ":sep_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no SEP tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetInternalSEP(c.Request().Context(), sepNumber)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetInternalSEP", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
