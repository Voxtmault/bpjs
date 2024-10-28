package sub_routes

import "github.com/labstack/echo/v4"

func Referral(e *echo.Echo, group *echo.Group) {
	referral := group.Group("/referral")

	// Getter
	referral.GET("/:referral_number", nil)
	referral.GET("/participant/:card_number/:bulk", nil)
	referral.GET("/outgoing/:start_date/:end_date", nil)
	referral.GET("/outgoing/:referral_number", nil)
	referral.GET("/sep_count/:referral_type/:referral_number", nil)

	// CUD
	referral.POST("", nil)
	referral.PUT("", nil)
	referral.DELETE("", nil)

	// Special Referral
	referral.GET("/special/:month/:year", nil)
	referral.POST("/special", nil)
	referral.DELETE("/special", nil)

	// Utility
	referred := group.Group("/referred")
	referred.GET("/specialists/:hf_code/:referral_date", nil)
	referred.GET("/health_facility/:hf_code", nil)
}
