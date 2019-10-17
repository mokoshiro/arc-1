package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const msgFmt = "invalid request: %v"

type requestValidator interface {
	Validate() error
}

// RequestValidationUnaryServerInterceptor validates the Unary gRPC's request payload if
// the request implements requestValidator interface.
// An InvalidArgument with the detail message will be returned to client if
// the validation was not passed.
func RequestValidationUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(requestValidator); ok {
			if err := v.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, fmt.Sprintf(msgFmt, err))
			}
		}
		return handler(ctx, req)
	}
}

// RequestValidationStreamServerInterceptor validates the streaming gRPC's request payload if
// the request implements requestValidator interface.
// An InvalidArgument with the detail message will be returned to client if
// the validation was not passed.
func RequestValidationStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{stream}
		return handler(srv, wrapper)
	}
}

type recvWrapper struct {
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	if v, ok := m.(requestValidator); ok {
		if err := v.Validate(); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(msgFmt, err))
		}
	}
	return nil
}
