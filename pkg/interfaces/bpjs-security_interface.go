package interfaces

import "context"

type BPJSSecurity interface {
	CreateSignature(ctx context.Context, timestamp int64) (string, error)
	DecryptResponse(ctx context.Context, timestamp int64, message string) (string, error)
}
