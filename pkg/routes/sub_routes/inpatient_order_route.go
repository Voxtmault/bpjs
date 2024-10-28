package sub_routes

import "github.com/labstack/echo/v4"

func InpatientOrder(group *echo.Group) {

	inpatientOrder := group.Group("/inpatient_order")

	inpatientOrder.POST("", nil)
	inpatientOrder.PUT("", nil)
}
