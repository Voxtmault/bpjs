package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func InpatientOrder(e *echo.Echo, group *echo.Group) {

	inpatientOrderController := controllers.NewInpatientOrderController(
		services.NewControlPlanService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	inpatientOrder := group.Group("/inpatient_order")

	inpatientOrder.POST("", inpatientOrderController.CreateInpatientOrder)
	inpatientOrder.PUT("", inpatientOrderController.UpdateInpatientOrder)
}
