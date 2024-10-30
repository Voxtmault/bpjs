package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func Monitoring(e *echo.Echo, group *echo.Group) {
	NewMonitoringController := controllers.NewMonitoringController(
		services.NewMonitoringService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	monitoring := group.Group("/monitoring")

	monitoring.GET("/visit/:service_date/:service_type", NewMonitoringController.GetMonitoringDataKunjungan)
	monitoring.GET("/claim/:service_type/:discharge_date/:claim_status", NewMonitoringController.GetMonitoringDataKlaim)
	monitoring.GET("/claim/jasa_raharja/:service_type/:start_date/:end_date", NewMonitoringController.GetMonitoringKlaimJaminanJasaraharja)
	monitoring.GET("/participant/history/:card_number/:start_date/:end_date", NewMonitoringController.GetMonitoringHistoryPelayananPeserta)
}
