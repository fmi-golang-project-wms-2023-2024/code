package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/nikola-enter21/wms/backend/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	log = logging.MustNewLogger()

	// Quick and hacky fix to exclude some rpcs from authn/authz
	skipJWTCheckMethods = map[string]struct{}{
		"/user.v1.UserService/CreateUser": {},
		"/user.v1.UserService/LoginUser":  {},
	}
)

func UnaryJWTInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Infow("UnaryInterceptor", "Method", info.FullMethod)

		if _, skip := skipJWTCheckMethods[info.FullMethod]; skip {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("missing metadata from context")
		}

		authHeaders, exists := md["authorization"]
		if !exists || len(authHeaders) == 0 {
			return nil, errors.New("missing authorization header")
		}

		authHeader := authHeaders[0]
		tokenParts := strings.Fields(authHeader)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return nil, errors.New("invalid authorization header format")
		}

		jwtToken := tokenParts[1]
		claims, err := parseJWT(jwtToken)
		if err != nil {
			return nil, err
		}

		ctx = SetSubInContext(ctx, claims["sub"].(string))
		ctx = SetRoleInContext(ctx, claims["role"].(string))

		resp, err := handler(ctx, req)
		return resp, err
	}
}

func StreamJWTInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		err = handler(srv, stream)
		return err
	}
}
