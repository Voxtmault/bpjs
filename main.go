package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/logger"
	"github.com/voxtmault/bpjs-rs-module/pkg/routes"
	"github.com/voxtmault/bpjs-rs-module/pkg/rpc"
	"github.com/voxtmault/bpjs-rs-module/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pbBPJS "github.com/voxtmault/bpjs-service-proto/go"
)

func main() {
	slog.Info("starting service...")
	AppConfig := config.New(".env")
	timeLoc, _ := time.LoadLocation(AppConfig.AppTimezone)
	time.Local = timeLoc

	if AppConfig.AppMode == "debug" {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else {
		slog.SetLogLoggerLevel(slog.LevelInfo)
	}

	// Adjust to your needs
	slog.Info("opening mariadb connection")
	if err := storage.InitMariaDB(&AppConfig.DBConfig); err != nil {
		panic(err)
	}
	slog.Info("opening redis connection")
	if err := storage.InitRedis(&AppConfig.RedisConfig); err != nil {
		panic(err)
	}

	slog.Info("initializing logger")
	if err := logger.InitLogger(&AppConfig.LoggingConfig); err != nil {
		panic(err)
	}
	slog.Info("registering request logger")
	logger.InitRequestLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", AppConfig.AppGRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Experimental Rate Limiter
	// rateLim := interceptors.NewRateLimiter(rate.Every(time.Minute/10), 10)

	// Try using SSL / TLS
	var s *grpc.Server
	if AppConfig.SSLConfig.CertPath != "" && AppConfig.SSLConfig.KeyPath != "" {
		creds, err := credentials.NewServerTLSFromFile(AppConfig.SSLConfig.CertPath, AppConfig.SSLConfig.KeyPath)
		if err != nil {
			slog.Error("failed to load server TLS credentials", "error", err)
			panic(err)
		}
		slog.Info("using ssl/tls")
		s = grpc.NewServer(
			// grpc.UnaryInterceptor(intercept.UnaryServerInterceptor),
			grpc.Creds(creds),
		)
	} else {
		slog.Warn("not using ssl/tls")
		s = grpc.NewServer(
		// grpc.UnaryInterceptor(intercept.UnaryServerInterceptor),
		)
	}

	// Init gRPC Services
	bpjsService := rpc.InitRPCService()
	pbBPJS.RegisterParticipantServiceServer(s, bpjsService.ParticipantService)
	pbBPJS.RegisterReferenceServiceServer(s, bpjsService.ReferenceService)
	pbBPJS.RegisterSEPServiceServer(s, bpjsService.SEPService)

	go func() {
		slog.Info("BPJS gRPC server listening at", "port", lis.Addr().String())

		if err = s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Init HTTP Services
	routes.InitRoute()

	// Graceful shutdown
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)
	<-interupt

	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	slog.Info("Received Shutdown Signal. Shutting Down Server......")

	// Stopping gRPC Service
	s.Stop()

	// Closes connections
	slog.Info("Closing Database Connection")
	if err := storage.Close(); err != nil {
		slog.Error("Error closing DB connection")
		panic(err)
	}

	slog.Info("Closing Redis Connection")
	if err := storage.CloseRedis(); err != nil {
		slog.Error("Error closing redis connection")
		panic(err)
	}

	slog.Info("Closing Logger")
	if err := logger.CloseLogger(); err != nil {
		slog.Error("Error closing logger")
		panic(err)
	}
}
