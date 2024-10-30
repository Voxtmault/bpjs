package controllers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
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

	var obj models.ReferralAction
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return c.JSON(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.CreateReferral(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> CreateReferral", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) UpdateReferral(c echo.Context) error {
	var res Response

	var obj models.ReferralAction
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return c.JSON(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.UpdateReferral(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> UpdateReferral", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) DeleteReferral(c echo.Context) error {
	var res Response

	var obj struct {
		ReferralNumber string `json:"noRujukan" validate:"required"`
		User           string `json:"user" validate:"required"`
	}
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return c.JSON(http.StatusBadRequest, res)
	}

	err := s.service.DeleteReferral(c.Request().Context(), obj.ReferralNumber, obj.User)
	if err != nil {
		slog.Error("sep_controller -> DeleteReferral", "stack trace", err)
		return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) GetSpecialReferral(c echo.Context) error {
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
	res.Data, err = s.service.GetSpecialReferrals(c.Request().Context(), month, year)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetSpecialReferral", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) CreateSpecialReferral(c echo.Context) error {
	var res Response

	var obj models.SpecialReferralCreate
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.CreateSpecialReferral(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> CreateSpecialReferral", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) DeleteSpecialReferral(c echo.Context) error {
	var res Response

	var obj models.SpecialReferralDelete
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); len(errMap) > 0 {
		res.Message = "Validation Error"
		res.Data = errMap

		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.DeleteSpecialReferral(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> CreateSpecialReferral", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) GetReferredSpecialist(c echo.Context) error {
	var res Response

	hfCode := c.Param("hf_code")
	if hfCode == "" || hfCode == ":hf_code" {
		return echo.NewHTTPError(http.StatusBadRequest, "kode faskes tidak boleh kosong")
	}

	referralDate := c.Param("referral_date")
	if referralDate == "" || referralDate == ":referral_date" {
		return echo.NewHTTPError(http.StatusBadRequest, "tanggal rujukan tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetReferredSpecialist(c.Request().Context(), hfCode, referralDate)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetReferredSpecialist", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ReferralControllers) GetReferredHealthFacility(c echo.Context) error {
	var res Response

	hfCode := c.Param("hf_code")
	if hfCode == "" || hfCode == ":hf_code" {
		return echo.NewHTTPError(http.StatusBadRequest, "kode faskes tidak boleh kosong")
	}

	var err error
	res.Data, err = s.service.GetReferredFacilities(c.Request().Context(), hfCode)
	if err != nil {
		if res.Data != "" {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("sep_controller -> GetReferredHealthFacility", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
