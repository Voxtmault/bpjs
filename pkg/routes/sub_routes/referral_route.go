package sub_routes

import (
	"github.com/labstack/echo/v4"
	"github.com/voxtmault/bpjs-rs-module/pkg/controllers"
	"github.com/voxtmault/bpjs-rs-module/pkg/services"
)

func Referral(e *echo.Echo, group *echo.Group) {

	referralController := controllers.NewReferralControllers(
		services.NewReferralService(
			services.NewBPJSRequestHandlerService(
				services.NewBPJSSecurityService(),
			),
		),
		e.Validator,
	)

	referral := group.Group("/referral")

	// Getter
	referral.GET("/:referral_number/:source", referralController.GetReferralViaReferralLetter)
	referral.GET("/participant/:card_number/:bulk/:source", referralController.GetReferralViaCardNumber)
	referral.GET("/outgoing/:start_date/:end_date", referralController.GetOutgoingReferral)
	referral.GET("/outgoing/:referral_number", referralController.GetOutgoingReferralDetail)
	referral.GET("/sep_count/:referral_type/:referral_number", referralController.GetReferralSEPCount)

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
