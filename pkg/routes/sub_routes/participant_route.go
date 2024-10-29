package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func Participant(e *echo.Echo, group *echo.Group) {

	participantController := controllers.NewParticipantController(
		services.NewParticipantService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	participant := group.Group("/participant")
	participant.GET("/nik/:nik/:service_date", participantController.GetParticipant)
	participant.GET("/card_number/:card_number/:service_date", participantController.GetParticipant)
}
