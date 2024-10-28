package sub_routes

import "github.com/labstack/echo/v4"

func ControlPlan(group *echo.Group) {

	controlPlan := group.Group("/control_plan")

	controlPlan.POST("", nil)
	controlPlan.PUT("", nil)
	controlPlan.DELETE("", nil)

	// Returns 1 object
	controlPlan.GET("/control_plan/:control_plan_number", nil)
	controlPlan.GET("/sep/:sep", nil)

	// Returns an array of object(s)
	controlPlan.GET("/list/card_number/:card_number/:month/:year/:filter", nil)
	controlPlan.GET("/list/:start_date/:end_date/:filter", nil)

	// Utility
	controlPlan.GET("/specialist/:control_type/:number/:expected_date", nil)
	controlPlan.GET("/doctor_schedule/:control_type/:poly_code/:expected_date", nil)
}
