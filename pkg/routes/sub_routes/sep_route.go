package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func SEP(e *echo.Echo, group *echo.Group) {

	sep := group.Group("/sep")

	sepController := controllers.NewSEPController(
		services.NewSEPService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	// CRUD
	sep.GET("/:sep_number", sepController.Get)
	sep.POST("", sepController.Post)
	sep.PUT("", sepController.Put)
	sep.DELETE("", sepController.Delete)

	// Approval, Usually for backdate SEP
	sep.POST("/submission", sepController.SubmitSEPRequest)
	sep.POST("/approval", sepController.ApproveSEPRequest)
	sep.GET("/approval/:month/:year", sepController.GetSEPRequest)

	// Utility
	sep.PUT("/discharge_date", sepController.UpdateDischargeDate)
	sep.GET("/get_discharged_sep/:month/:year/:filter", sepController.GetDischargedSEP)

	// Internal
	internalSep := sep.Group("/internal")
	internalSep.GET("/:sep_number", sepController.GetInternalSEP)
	internalSep.DELETE("", sepController.DeleteInternalSEP)
}

func Suppletion(e *echo.Echo, group *echo.Group) {

	suppletion := group.Group("/suppletion")

	suppletionController := controllers.NewSuppletionController(
		services.NewSuppletionService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	// CRUD
	suppletion.GET("/jasa_raharja/:service_date/:card_number", suppletionController.GetJasaRaharjaSuppletion)
	suppletion.GET("/accident_master_data/:card_number", suppletionController.GetAccidentMasterData)
}
