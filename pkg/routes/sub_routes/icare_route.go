package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func ICare(e *echo.Echo, group *echo.Group) {
	iCareController := controllers.NewICareController(
		services.NewIcareService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	iCare := group.Group("/i_care")

	iCare.GET("/:no_card/:kode_dokter", iCareController.FKRTLIcare)
}
