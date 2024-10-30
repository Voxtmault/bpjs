package controllers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
)

type PRBController struct {
	service  interfaces.PRB
	validate echo.Validator
}

func NewPRBController(service interfaces.PRB, validate echo.Validator) *PRBController {
	return &PRBController{
		service:  service,
		validate: validate,
	}
}

func (s PRBController) Post(c echo.Context) error {
	var res Response
	var obj models.PRBInsertRequest
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.PRBInsert(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("PRB_controller -> post", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s PRBController) PUT(c echo.Context) error {
	var res Response
	var obj models.PRBUpdateRequest
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.PRBUpdate(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("PRB_controller -> PUT", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s PRBController) Delete(c echo.Context) error {
	var res Response
	var obj models.PRBDeleteRequest
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.PRBDelete(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("PRB_controller -> PUT", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s PRBController) GetByNoSRB(c echo.Context) error {
	var res Response
	noSep := c.Param("no_sep")
	if noSep == "" || noSep == ":no_sep" {
		res.Message = "no_sep tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	noPrb := c.Param("no_prb")
	if noPrb == "" || noPrb == ":no_prb" {
		res.Message = "no_sep tidak valid"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetPRBbyNomorSRB(c.Request().Context(), noPrb, noSep)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("PRB_controller -> GetByNoSRB", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s PRBController) GetByTanggal(c echo.Context) error {
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
	var err error
	res.Data, err = s.service.GetPRBbyTanggal(c.Request().Context(), tanggalAwal, tanggalAkhir)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("PRB_controller -> GetByTanggal", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
