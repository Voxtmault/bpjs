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
	referral.GET("/v2/:referral_number/:source", referralController.GetReferralViaReferralLetterV2)
	referral.GET("/participant/:card_number/:bulk/:source", referralController.GetReferralViaCardNumber)
	referral.GET("/outgoing/:start_date/:end_date", referralController.GetOutgoingReferral)
	referral.GET("/outgoing/:referral_number", referralController.GetOutgoingReferralDetail)
	referral.GET("/sep_count/:referral_type/:referral_number", referralController.GetReferralSEPCount)

	// CUD
	referral.POST("", referralController.CreateReferral)
	referral.PUT("", referralController.UpdateReferral)
	referral.DELETE("", referralController.DeleteReferral)

	// Special Referral
	referral.GET("/special/:month/:year", referralController.GetSpecialReferral)
	referral.POST("/special", referralController.CreateSpecialReferral)
	referral.DELETE("/special", referralController.DeleteSpecialReferral)

	// Utility
	referred := group.Group("/referred")
	referred.GET("/specialists/:hf_code/:referral_date", referralController.GetReferredSpecialist)
	referred.GET("/health_facility/:hf_code", referralController.GetReferredHealthFacility)
}
