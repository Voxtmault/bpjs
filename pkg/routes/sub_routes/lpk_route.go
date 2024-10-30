package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func LPK(e *echo.Echo, group *echo.Group) {
	lpkController := controllers.NewLPKController(
		services.NewLPKService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	lpk := group.Group("/lpk")

	lpk.GET("/:service_type/:tanggal_masuk", lpkController.GET)
	lpk.POST("", lpkController.POST)
	lpk.PUT("", lpkController.PUT)
	lpk.DELETE("/:no_sep", lpkController.DELETE)
}
