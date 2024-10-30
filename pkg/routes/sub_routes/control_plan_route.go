package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func ControlPlan(e *echo.Echo, group *echo.Group) {
	NewControlPlanService := controllers.NewControlPlanController(
		services.NewControlPlanService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	controlPlan := group.Group("/control_plan")

	controlPlan.POST("", NewControlPlanService.Post)
	controlPlan.PUT("", NewControlPlanService.Put)
	controlPlan.DELETE("/:control_number", NewControlPlanService.Delete)

	// Returns 1 object
	controlPlan.GET("/control_plan/:control_plan_number", NewControlPlanService.GetByControlPlanNumber)
	controlPlan.GET("/sep/:sep", NewControlPlanService.GetBySEP)

	// Returns an array of object(s)
	controlPlan.GET("/list/card_number/:card_number/:month/:year/:filter", NewControlPlanService.GetListCard)
	controlPlan.GET("/list/:start_date/:end_date/:filter", NewControlPlanService.GetByControlPlanList)

	// Utility
	controlPlan.GET("/specialist/:control_type/:number/:expected_date", NewControlPlanService.GetClinicControl)
	controlPlan.GET("/doctor_schedule/:control_type/:poly_code/:expected_date", NewControlPlanService.GetDoctorSchedule)
}
