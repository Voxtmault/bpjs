package routes

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/models"
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

	validate := validator.New()
	validate.RegisterValidation("specialDiag", models.ValidateSpecialReferralDiagnosisCode)

	// Register Validator
	e.Validator = &CustomValidator{validator: validate}

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
	sub_routes.PRB(e, bpjs)
	sub_routes.LPK(e, bpjs)
	sub_routes.ICare(e, bpjs)
	sub_routes.Aplicares(e, bpjs)

	// Finally Start the Service
	StartHTTPService(e)
}

func StartHTTPService(e *echo.Echo) {
	cfg := config.GetConfig()
	go func() {
		slog.Info("Starting Service...")

		// HTTPS
		if cfg.SSLConfig.CertPath != "" && cfg.SSLConfig.KeyPath != "" {
			slog.Info("Starting in HTTPS")
			slog.Info("Port: " + cfg.AppPort)

			// Redirect all HTTP Traffic to HTTPS since the API is running in HTTPS
			e.Pre(middleware.HTTPSRedirect())

			s := http.Server{
				Addr:    fmt.Sprintf(":%s", cfg.AppPort),
				Handler: e,
			}
			if err := s.ListenAndServeTLS(cfg.SSLConfig.CertPath, cfg.SSLConfig.KeyPath); err != nil {
				slog.Warn("Shutting down the service...")
			}
		} else {
			//HTTP
			slog.Info("Starting in HTTP")
			slog.Info("Port: " + cfg.AppPort)
			if err := e.Start(":" + cfg.AppPort); err != nil {
				slog.Warn("Shutting down the service...")
			}
		}
	}()
}
