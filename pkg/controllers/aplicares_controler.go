package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type AplicaresController struct {
	service  interfaces.RuanganAplicares
	validate echo.Validator
}

func NewAplicaresControllerController(service interfaces.RuanganAplicares, validate echo.Validator) *AplicaresController {
	return &AplicaresController{
		service:  service,
		validate: validate,
	}
}

func (s AplicaresController) GetReferensiKamar(c echo.Context) error {
	var res Response

	var err error
	res.Data, err = s.service.GetReferensiJenisKamar(c.Request().Context())
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("aplicares_controller -> GetReferensiKamar", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s AplicaresController) GetKetersediaanKamar(c echo.Context) error {
	var res Response
	start := c.FormValue("start")
	if start == "" || start == ":start" {
		res.Message = "start tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	limit := c.FormValue("limit")
	if limit == "" || limit == ":limit" {
		res.Message = "limit tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetKetersediaanKamar(c.Request().Context(), start, limit)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("aplicares_controller -> GetKetersediaanKamar", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s AplicaresController) Post(c echo.Context) error {
	var res Response
	var obj models.Ruangan
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var err error
	res.Data, err = s.service.PostRuangan(c.Request().Context(), &obj, "post")
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("aplicares_controller -> post", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s AplicaresController) Put(c echo.Context) error {
	var res Response
	var obj models.Ruangan
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var err error
	res.Data, err = s.service.PostRuangan(c.Request().Context(), &obj, "put")
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("aplicares_controller -> put", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s AplicaresController) Delete(c echo.Context) error {
	var res Response
	var obj models.Ruangan
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var err error
	res.Data, err = s.service.PostRuangan(c.Request().Context(), &obj, "delete")
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("aplicares_controller -> delete", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
