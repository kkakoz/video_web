package local

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func NewCtx(req *http.Request) context.Context {
	md := metadata.New(nil)
	for k, v := range req.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(req.Context(), md)
	return newCtx
}
