package controllers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type LPKController struct {
	service  interfaces.LPK
	validate echo.Validator
}

func NewLPKController(service interfaces.LPK, validate echo.Validator) *LPKController {
	return &LPKController{
		service:  service,
		validate: validate,
	}
}

func (s LPKController) POST(c echo.Context) error {
	var res Response
	var obj models.InsertLPKRequest
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.LPKInsert(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("lpk_controller -> post", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s LPKController) PUT(c echo.Context) error {
	var res Response
	var obj models.InsertLPKRequest
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.LPKUpdate(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("lpk_controller -> put", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s LPKController) DELETE(c echo.Context) error {
	var res Response
	noSep := c.Param("no_sep")
	if noSep == "" || noSep == ":no_sep" {
		res.Message = "no_sep tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.LPKDelete(c.Request().Context(), noSep)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("lpk_controller -> delete", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s LPKController) GET(c echo.Context) error {
	var res Response

	tanggalMasuk := c.Param("tanggal_masuk")
	if tanggalMasuk == "" || tanggalMasuk == ":tanggal_masuk" {
		res.Message = "Tanggal tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, tanggalMasuk)
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
	res.Data, err = s.service.LPKGet(c.Request().Context(), tanggalMasuk, jenisPelayanan)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("lpk_controller -> delete", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
