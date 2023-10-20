package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

var DefaultHttpMiddleWare = []http.ServerOption{
	http.Middleware(
		recovery.Recovery(),
		validate.Validator(),
		selector.Server().
			Match(NewWhiteListMatcher()).
			Build(),
	),
	http.Filter(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}),
		handlers.AllowCredentials(),
	)),
}

var DefaultGrpcMiddleWare = []grpc.ServerOption{
	grpc.Middleware(
		recovery.Recovery(),
		validate.Validator(),
	),
}
