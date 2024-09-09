package interceptors

import (
	"context"
	"log/slog"

	"github.com/voxtmault/bpjs-rs-module/pkg/logger"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	slog.Info("New Request: " + info.FullMethod)

	serverLogger := logger.GetServerLogger()
	serverLogger.Write([]byte("New Request: " + info.FullMethod))

	return handler(ctx, req)
}
