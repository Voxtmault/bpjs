package routes

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/routes/sub_routes"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitRoute() {
	e := echo.New()
	e.HideBanner = true

	// Register Validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Registers Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())

	// For compression and decompression
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
		Level: 6,
	}))
	e.Use(middleware.DecompressWithConfig(middleware.DecompressConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))

	// For Performance Metrics
	e.Use(echoprometheus.NewMiddleware("bpjs"))
	e.GET("/metrics", echoprometheus.NewHandler())

	// Main Route
	bpjs := e.Group(config.GetConfig().AppRoot)

	// Sub Routes
	sub_routes.ControlPlan(e, bpjs)
	sub_routes.InpatientOrder(e, bpjs)
	sub_routes.Participant(e, bpjs)
	sub_routes.Reference(e, bpjs)
	sub_routes.Referral(e, bpjs)
	sub_routes.SEP(e, bpjs)
	sub_routes.Suppletion(e, bpjs)
	sub_routes.Monitoring(e, bpjs)
	sub_routes.FingerPrint(e, bpjs)

}
