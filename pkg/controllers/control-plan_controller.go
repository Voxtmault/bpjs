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

type ControlPlanController struct {
	service  interfaces.ControlPlan
	validate echo.Validator
}

func NewControlPlanController(service interfaces.ControlPlan, validate echo.Validator) *ControlPlanController {
	return &ControlPlanController{
		service:  service,
		validate: validate,
	}
}

func (s ControlPlanController) Post(c echo.Context) error {
	var res Response
	var obj models.ControlPlanCreate
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.CreateControlPlan(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> post", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) Put(c echo.Context) error {
	var res Response
	var obj models.UpdateControlPlans
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	res.Data, err = s.service.UpdateControlPlan(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> put", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) Delete(c echo.Context) error {
	var res Response
	controlNumber := c.Param("control_number")

	if controlNumber == "" || controlNumber == ":control_number" {
		res.Message = "control_number tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	user := c.Param("user")

	if controlNumber == "" || controlNumber == ":user" {
		res.Message = "user tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	err := s.service.DeleteControlPlan(c.Request().Context(), controlNumber, user)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> delete", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) GetByControlPlanNumber(c echo.Context) error {
	var res Response
	controlNumber := c.Param("control_plan_number")

	if controlNumber == "" || controlNumber == ":control_plan_number" {
		res.Message = "control_plan_number tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetViaControlLetterNumber(c.Request().Context(), controlNumber)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> GetByControlPlanNumber", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) GetBySEP(c echo.Context) error {
	var res Response
	controlNumber := c.Param("sep")

	if controlNumber == "" || controlNumber == ":sep" {
		res.Message = "nomor sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetViaSEP(c.Request().Context(), controlNumber)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> GetBySEP", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) GetByControlPlanList(c echo.Context) error {
	var res Response
	params := models.ControlPlanParams{}

	params.StartDate = c.Param("start_date")

	if params.StartDate == "" || params.StartDate == ":start_date" {
		res.Message = "start_date sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, params.StartDate)
		if err != nil {
			res.Message = "start_date tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}

	params.EndDate = c.Param("end_date")

	if params.StartDate == "" || params.StartDate == ":end_date" {
		res.Message = "end_date sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	} else {
		_, err := time.Parse(time.DateOnly, params.StartDate)
		if err != nil {
			res.Message = "end_date tidak Valid"
			return echo.NewHTTPError(http.StatusBadRequest, res)
		}
	}

	params.Filter = c.Param("filter")

	if params.Filter == "" || params.Filter == ":filter" {
		res.Message = "filter sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.GetControlPlans(c.Request().Context(), &params)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> GetByControlPlanList", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) GetDoctorSchedule(c echo.Context) error {
	var res Response
	var obj models.DoctorScheduleParams

	obj.ControlType = c.Param("control_type")

	if obj.ControlType == "" || obj.ControlType == ":control_type" {
		res.Message = "control_type sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.PoliCode = c.Param("poly_code")

	if obj.PoliCode == "" || obj.PoliCode == ":poly_code" {
		res.Message = "poly_code sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.ControlPlannedDate = c.Param("expected_date")

	if obj.ControlPlannedDate == "" || obj.ControlPlannedDate == ":expected_date" {
		res.Message = "expected_date sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetDoctorPracticeSchedule(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> GetDoctorSchedule", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) GetClinicControl(c echo.Context) error {
	var res Response
	var obj models.ClinicControlParams
	// /:control_type/:number/:expected_date
	obj.ControlType = c.Param("control_type")

	if obj.ControlType == "" || obj.ControlType == ":control_type" {
		res.Message = "control_type sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.Identifier = c.Param("number")

	if obj.Identifier == "" || obj.Identifier == ":number" {
		res.Message = "number sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.ControlPlannedDate = c.Param("expected_date")

	if obj.ControlPlannedDate == "" || obj.ControlPlannedDate == ":expected_date" {
		res.Message = "expected_date sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetClinicControlPlans(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> GetClinicControl", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s ControlPlanController) GetListCard(c echo.Context) error {
	var res Response
	var obj models.ControlPlansFromCardNumberParams
	// :card_number/:month/:year/:filter
	obj.CardNumber = c.Param("card_number")

	if obj.CardNumber == "" || obj.CardNumber == ":card_number" || len(obj.CardNumber) < 13 {
		res.Message = "card_number sep tidak boleh kosong / tidak valid "
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.Month = c.Param("month")

	if obj.Month == "" || obj.Month == ":month" || len(obj.Month) < 2 {
		res.Message = "month tidak boleh kosong / tidak valid "
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.Year = c.Param("year")

	if obj.Year == "" || obj.Year == ":year" {
		res.Message = "year sep tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	obj.Filter = c.Param("filter")
	if obj.Year == "" || obj.Year == ":filter" {
		res.Message = "filter tidak boleh kosong"
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}
	var err error
	res.Data, err = s.service.GetControlPlanFromCardNumber(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("control-plan_controller -> GetClinicControl", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
