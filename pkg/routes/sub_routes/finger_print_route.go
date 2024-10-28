package sub_routes

import "github.com/labstack/echo/v4"

func FingerPrint(e *echo.Echo, group *echo.Group) {

	finger_print := group.Group("/finger_print")
	finger_print.GET("/:service_date/:card_number", nil)
	finger_print.GET("/:service_date", nil)

	finger_print.GET("/random_question/:service_date/:card_number", nil)
	finger_print.POST("/random_question/answer", nil)
}
