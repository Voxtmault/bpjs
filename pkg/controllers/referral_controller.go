package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
)

type ReferralControllers struct {
	service  interfaces.Referral
	validate echo.Validator
}

func NewReferralControllers(service interfaces.Referral, validate echo.Validator) *ReferralControllers {
	return &ReferralControllers{
		service:  service,
		validate: validate,
	}
}

func (s ReferralControllers) GetReferralViaReferralLetter(c echo.Context) error {
	var res Response

	referralNumber := c.Param("referral_number")
	if referralNumber == "" || referralNumber == ":referral_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no surat rujukan tidak boleh kosong")
	}

	source := c.Param("source")
	if source == "" || source == ":source" {
		return echo.NewHTTPError(http.StatusBadRequest, "asal surat rujukan tidak boleh kosong")
	}
	parsedSource, _ := strconv.Atoi(source)

	var err error
	res.Data, err = s.service.GetParticipantReferralByReferralNumber(c.Request().Context(), referralNumber, uint(parsedSource))
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetReferralViaReferralLetter", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) GetReferralViaCardNumber(c echo.Context) error {
	var res Response

	cardNumber := c.Param("card_number")
	if cardNumber == "" || cardNumber == ":card_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no kartu tidak boleh kosong")
	}

	source := c.Param("source")
	if source == "" || source == ":source" {
		return echo.NewHTTPError(http.StatusBadRequest, "asal surat rujukan tidak boleh kosong")
	}
	parsedSource, _ := strconv.Atoi(source)

	bulk := c.Param("bulk")
	if bulk == "" || bulk == ":bulk" {
		return echo.NewHTTPError(http.StatusBadRequest, "asal surat rujukan tidak boleh kosong")
	}
	parsedBulk, _ := strconv.ParseBool(bulk)

	var err error
	res.Data, err = s.service.GetParticipantReferralByBPJSNumber(c.Request().Context(), cardNumber, uint(parsedSource), parsedBulk)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetReferralViaCardNumber", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) GetOutgoingReferral(c echo.Context) error {
	var res Response

	startDate := c.Param("start_date")
	if startDate == "" || startDate == ":start_date" {
		return echo.NewHTTPError(http.StatusBadRequest, "tanggal mulai tidak boleh kosong")
	}

	endDate := c.Param("end_date")
	if endDate == "" || endDate == ":end_date" {
		return echo.NewHTTPError(http.StatusBadRequest, "tanggal akhir tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetOutgoingReferral(c.Request().Context(), startDate, endDate)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetOutgoingReferral", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) GetOutgoingReferralDetail(c echo.Context) error {
	var res Response

	referralNumber := c.Param("referral_number")
	if referralNumber == "" || referralNumber == ":referral_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no surat rujukan tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetOutgoingReferralDetail(c.Request().Context(), referralNumber)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetOutgoingReferralDetail", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) GetReferralSEPCount(c echo.Context) error {
	var res Response

	referralType := c.Param("referral_type")
	if referralType == "" || referralType == ":referral_type" {
		return echo.NewHTTPError(http.StatusBadRequest, "tipe rujukan tidak boleh kosong")
	}

	referralNumber := c.Param("referral_number")
	if referralNumber == "" || referralNumber == ":referral_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no surat rujukan tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetReferralSEPCount(c.Request().Context(), referralType, referralNumber)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetReferralSEPCount", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) CreateReferral(c echo.Context) error {
	var res Response

	referralType := c.Param("referral_type")
	if referralType == "" || referralType == ":referral_type" {
		return echo.NewHTTPError(http.StatusBadRequest, "tipe rujukan tidak boleh kosong")
	}

	referralNumber := c.Param("referral_number")
	if referralNumber == "" || referralNumber == ":referral_number" {
		return echo.NewHTTPError(http.StatusBadRequest, "no surat rujukan tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetReferralSEPCount(c.Request().Context(), referralType, referralNumber)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetReferralSEPCount", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
