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

type InpatientOrderController struct {
	service  interfaces.ControlPlan
	validate echo.Validator
}

func NewInpatientOrderController(service interfaces.ControlPlan, validate echo.Validator) *InpatientOrderController {
	return &InpatientOrderController{
		service:  service,
		validate: validate,
	}
}

func (s InpatientOrderController) CreateInpatientOrder(c echo.Context) error {
	var res Response

	var obj models.ControlPlanCreate
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); errMap != nil {
		res.Message = "Validation Error"
		res.Data = errMap

		return echo.NewHTTPError(http.StatusBadRequest, errMap)
	}

	var err error
	res.Data, err = s.service.CreateInpatientCareOrder(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("inpatient_order_controller -> CreateInpatientOrder", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}

func (s InpatientOrderController) UpdateInpatientOrder(c echo.Context) error {
	var res Response

	var obj models.UpdateControlPlans
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); errMap != nil {
		res.Message = "Validation Error"
		res.Data = errMap

		return echo.NewHTTPError(http.StatusBadRequest, errMap)
	}

	var err error
	res.Data, err = s.service.UpdateInpatientCareOrder(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, eris.Cause(err).Error())
		} else {
			slog.Error("inpatient_order_controller -> UpdateInpatientOrder", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, eris.Cause(err).Error())
		}
	}
	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
