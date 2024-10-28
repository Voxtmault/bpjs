package sub_routes

import "github.com/labstack/echo/v4"

func Monitoring(group *echo.Group) {
	monitoring := group.Group("/monitoring")

	monitoring.GET("/visit/:service_date/:service_type", nil)
	monitoring.GET("/claim/:service_date/:discharge_date/:claim_status", nil)
	monitoring.GET("/claim/jasa_raharja/:service_type/:start_date/:end_date", nil)
	monitoring.GET("/participant/history/:card_number/:start_date/:end_date", nil)
}
