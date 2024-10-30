package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func PRB(e *echo.Echo, group *echo.Group) {
	prbController := controllers.NewPRBController(
		services.NewPRBService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	prb := group.Group("/prb")

	prb.GET("/srb/:no_sep/:no_prb", prbController.GetByNoSRB)
	prb.GET("/date/:start_date/:end_date", prbController.GetByTanggal)
	prb.POST("", prbController.POST)
	prb.PUT("", prbController.PUT)
	prb.DELETE("", prbController.DELETE)
}
