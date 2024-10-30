package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func Aplicares(e *echo.Echo, group *echo.Group) {

	aplicaresController := controllers.NewAplicaresControllerController(
		services.NewRuanganService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	aplicares := group.Group("/aplicares")

	aplicares.GET("/referensi_kamar", aplicaresController.GetReferensiKamar)
	aplicares.GET("/ketersediaan_kamar", aplicaresController.GetKetersediaanKamar)
	aplicares.POST("", aplicaresController.Post)
	aplicares.PUT("", aplicaresController.Put)
	aplicares.DELETE("", aplicaresController.Delete)

}
