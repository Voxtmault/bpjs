package interceptors

import (
	"context"
	"log/slog"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RateLimiter struct {
	limiter *rate.Limiter
}

func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(r, b),
	}
}

// UnaryInterceptor is a gRPC unary server interceptor for rate limiting
func (rl *RateLimiter) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	slog.Info("Received new gRPC connection: " + info.FullMethod)
	if !rl.limiter.Allow() {
		return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
	}
	return handler(ctx, req)
}
