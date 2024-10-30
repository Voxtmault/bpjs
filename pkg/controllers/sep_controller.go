package controllers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
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

func (s SEPController) SubmitSEPRequest(c echo.Context) error {
	var res Response

	var obj models.SEPRequestCreate
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.RequestSEP(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> SubmitSEPRequest", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) ApproveSEPRequest(c echo.Context) error {
	var res Response

	var obj models.SEPRequestCreate
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.ApprovalSEPRequest(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> ApproveSEPRequest", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) GetSEPRequest(c echo.Context) error {
	var res Response

	month := c.Param("month")
	if month == "" || month == ":month" {
		return echo.NewHTTPError(http.StatusBadRequest, "bulan tidak boleh kosong")
	}

	year := c.Param("year")
	if year == "" || year == ":year" {
		return echo.NewHTTPError(http.StatusBadRequest, "tahun tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetSEPRequests(c.Request().Context(), month, year)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetSEPRequest", "stack trace", err)
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

func (s SEPController) DeleteInternalSEP(c echo.Context) error {
	var res Response

	var obj models.DeleteInternalSEP
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.DeleteInternalSEP(c.Request().Context(), &obj)
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

func (s SEPController) GetFingerprintAuthentication(c echo.Context) error {
	var res Response

	serviceDate := c.Param("service_date")
	if serviceDate == "" || serviceDate == ":service_date" {
		return echo.NewHTTPError(http.StatusBadRequest, "tanggal SEP tidak boleh kosong")
	}

	cardNumber := c.Param("card_number")
	if cardNumber == "" || cardNumber == ":card_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no kartu tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetFingerPrintSEP(c.Request().Context(), cardNumber, serviceDate)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetFingerprintAuthentication", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) GetAuthenticatedFingerprints(c echo.Context) error {
	var res Response

	serviceDate := c.Param("service_date")
	if serviceDate == "" || serviceDate == ":service_date" {
		return echo.NewHTTPError(http.StatusBadRequest, "tanggal SEP tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetListFingerPrintSEP(c.Request().Context(), serviceDate)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetAuthenticatedFingerprints", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) GetRandomQuestion(c echo.Context) error {
	var res Response

	serviceDate := c.Param("service_date")
	if serviceDate == "" || serviceDate == ":service_date" {
		return echo.NewHTTPError(http.StatusBadRequest, "tanggal SEP tidak boleh kosong")
	}

	cardNumber := c.Param("card_number")
	if cardNumber == "" || cardNumber == ":card_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no kartu tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetListRandomQuestion(c.Request().Context(), cardNumber, serviceDate)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetRandomQuestion", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s SEPController) AnswerRandomQuestion(c echo.Context) error {
	var res Response

	var obj models.PostRequestRandomQuestion
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if errMap := utils.MangleValidateResult(s.validator.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.PostRandomQuestion(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> AnswerRandomQuestion", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
