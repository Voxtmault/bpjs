package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func FingerPrint(e *echo.Echo, group *echo.Group) {

	sepController := controllers.NewSEPController(
		services.NewSEPService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	finger_print := group.Group("/finger_print")
	finger_print.GET("/:service_date/:card_number", sepController.GetFingerprintAuthentication)
	finger_print.GET("/:service_date", sepController.GetAuthenticatedFingerprints)

	finger_print.GET("/random_question/:service_date/:card_number", sepController.GetRandomQuestion)
	finger_print.POST("/random_question/answer", sepController.AnswerRandomQuestion)
}
