package routes

import (
	"strings"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/routes/sub_routes"
)

func InitRoute() {
	e := echo.New()
	e.HideBanner = true

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
	sub_routes.ControlPlan(bpjs)
	sub_routes.InpatientOrder(bpjs)
	sub_routes.Participant(bpjs)
	sub_routes.Reference(bpjs)
	sub_routes.Referral(bpjs)
	sub_routes.SEP(bpjs)
	sub_routes.Suppletion(bpjs)
	sub_routes.Monitoring(bpjs)
	sub_routes.FingerPrint(bpjs)

}
