package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/voxtmault/bpjs-rs-module/config"
	intercept "github.com/voxtmault/bpjs-rs-module/pkg/interceptors"
	"github.com/voxtmault/bpjs-rs-module/pkg/logger"
	rpc "github.com/voxtmault/bpjs-rs-module/pkg/rpc"
	"github.com/voxtmault/bpjs-rs-module/pkg/utils"

	pbBPJS "github.com/voxtmault/bpjs-service-proto/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	AppConfig := config.New(".env")
	timeLoc, _ := time.LoadLocation(AppConfig.AppTimezone)
	time.Local = timeLoc

	// Adjust to your needs
	// if err := storage.InitMariaDB(&AppConfig.DBConfig); err != nil {
	// 	panic(err)
	// }
	// if err := storage.InitRedis(&AppConfig.RedisConfig); err != nil {
	// 	panic(err)
	// }
	utils.InitValidator()
	if err := logger.InitLogger(&AppConfig.LoggingConfig); err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", AppConfig.AppPort))
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
			panic(fmt.Sprintf("SSL Config: %s", err))
		}
		log.Println("Using SSL / TLS")
		s = grpc.NewServer(
			grpc.UnaryInterceptor(intercept.UnaryServerInterceptor),
			grpc.Creds(creds),
		)
	} else {
		log.Println("Not Using SSL / TLS")
		s = grpc.NewServer(
			grpc.UnaryInterceptor(intercept.UnaryServerInterceptor),
		)
	}

	// Init gRPC Sercices
	bpjsService := rpc.InitRPCService()
	pbBPJS.RegisterParticipantServiceServer(s, bpjsService.ParticipantService)
	pbBPJS.RegisterReferenceServiceServer(s, bpjsService.ReferenceService)

	go func() {
		log.Printf("BPJS gRPC Server listening at %v", lis.Addr())

		if err = s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)
	<-interupt

	_, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	log.Println("Received Shutdown Signal. Shutting Down Server......")

	// Stopping gRPC Service
	s.Stop()

	// Closes connections
	// log.Println("Closing Database Connection")
	// if err := storage.Close(); err != nil {
	// 	log.Println("Error closing DB connection")
	// 	panic(err)
	// }
	// slog.Info("Closing Redis Connection")
	// if err := storage.CloseRedis(); err != nil {
	// 	log.Println("Error closing redis connection")
	// 	panic(err)
	// }
}
