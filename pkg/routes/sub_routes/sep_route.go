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
	sep.POST("/submission", nil)
	sep.POST("/approval", nil)
	sep.GET("/approval/:month/:year", nil)

	// Utility
	sep.PUT("/discharge_date", sepController.UpdateDischargeDate)
	sep.GET("/get_discharged_sep/:month/:year/:filter", sepController.GetDischargedSEP)

	// Internal
	internalSep := sep.Group("/internal")
	internalSep.GET("/:sep_number", nil)
	internalSep.DELETE("", nil)
}

func Suppletion(e *echo.Echo, group *echo.Group) {

	suppletion := group.Group("/suppletion")

	// CRUD
	suppletion.GET("/jasa_raharja/:service_date/:card_number", nil)
	suppletion.GET("/accident_master_data/:card_number", nil)
}
