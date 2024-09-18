package interfaces

import (
	"context"
	"net/http"
)

type RequestHandler interface {
	SendRequest(ctx context.Context, req *http.Request) (string, error)
}
