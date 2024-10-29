package controllers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"
)

type ParticipantController struct {
	service  interfaces.Participant
	validate echo.Validator
}

func NewParticipantController(service interfaces.Participant, validate echo.Validator) *ParticipantController {
	return &ParticipantController{
		service:  service,
		validate: validate,
	}
}

func (s ParticipantController) GetParticipant(c echo.Context) error {
	var res Response

	var obj models.ParticipantSearchParams
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if obj.ServiceDate == "" {
		obj.ServiceDate = time.Now().Format(time.DateTime)
	}

	if errMap := utils.MangleValidateResult(s.validate.Validate(obj)); errMap != nil {
		res.Message = "Validation Error"
		res.Data = errMap
		return echo.NewHTTPError(http.StatusBadRequest, res)
	}

	var err error
	res.Data, err = s.service.GetParticipant(c.Request().Context(), &obj)
	if err != nil {
		if res.Data != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			slog.Error("participant_controller -> get", "stack trace", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	res.Message = "Success"

	return c.JSON(http.StatusOK, res)
}
